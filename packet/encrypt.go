package packet

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/superp00t/gophercraft/arc4"
)

type Crypt struct {
	ServerDecryptionKey []byte
	ServerEncryptionKey []byte

	DecClient, EncServer *arc4.Cipher
}

func NewCrypt(sessionKey []byte, server bool) *Crypt {
	c := &Crypt{}
	c.ServerEncryptionKey = []byte{
		0xC2, 0xB3, 0x72, 0x3C, 0xC6, 0xAE, 0xD9, 0xB5,
		0x34, 0x3C, 0x53, 0xEE, 0x2F, 0x43, 0x67, 0xCE,
	}

	c.ServerDecryptionKey = []byte{
		0xCC, 0x98, 0xAE, 0x04, 0xE8, 0x97, 0xEA, 0xCA,
		0x12, 0xDD, 0xC0, 0x93, 0x42, 0x91, 0x53, 0x57,
	}

	decryptClient := hmac.New(sha1.New, c.ServerDecryptionKey)
	encryptServer := hmac.New(sha1.New, c.ServerEncryptionKey)
	encryptServer.Write(sessionKey)
	decryptClient.Write(sessionKey)
	encryptHash := encryptServer.Sum(nil)
	decryptHash := decryptClient.Sum(nil)

	if server {
		c.EncServer = arc4.ARC4(decryptHash)
		c.DecClient = arc4.ARC4(encryptHash)
	} else {
		c.DecClient = arc4.ARC4(decryptHash)
		c.EncServer = arc4.ARC4(encryptHash)
	}

	// Drop-1024 ARC4
	for i := 0; i < 1024; i++ {
		c.EncServer.Next()
		c.DecClient.Next()
	}

	return c
}

func (c *Crypt) Decrypt(input []byte) {
	c.DecClient.Decrypt(input)
}

func (c *Crypt) Encrypt(input []byte) {
	c.EncServer.Encrypt(input)
}

type Crypter struct {
	C          net.Conn
	SessionKey []byte
	Server     bool
	K          *Crypt

	ByteStream *bytes.Buffer

	TmpFrames [][]byte
	NewFrame  chan []byte

	errc chan error

	closed bool
}

func NewCrypter(c net.Conn, sessionKey []byte, server bool) *Crypter {
	cr := new(Crypter)
	cr.C = c
	cr.SessionKey = sessionKey
	cr.Server = server
	cr.K = NewCrypt(sessionKey, server)
	cr.NewFrame = make(chan []byte)
	cr.errc = make(chan error)

	go cr.SendHandler()

	return cr
}

func (c *Crypter) SendFrame(b []byte) error {
	ln := 4
	if c.Server {
		ln = 2
	}

	lnBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(lnBuf, uint16(len(b)))
	dat := append(lnBuf, b...)

	c.K.Encrypt(dat[:ln+2])

	// _, err := c.C.Write(dat)
	// if err != nil {
	// 	return err
	// }

	c.NewFrame <- dat
	e := <-c.errc
	log.Println("snd error", e)
	return e
}

const MaxFrameSize = 2048

// SendHandler awaits new data packets, and concatenates them into a single TCP packet
func (c *Crypter) SendHandler() {
	for {
		select {
		case msg := <-c.NewFrame:
			log.Println("Got new frame")
			c.TmpFrames = append(c.TmpFrames, msg)
			if len(bytes.Join(c.TmpFrames, nil)) > MaxFrameSize {
				c.sendOff()
			}
		case <-time.After(9 * time.Millisecond):
			if len(c.TmpFrames) != 0 {
				c.sendOff()
			}
		}
	}
}

//
func (c *Crypter) sendOff() {
	buf := bytes.Join(c.TmpFrames, nil)
	rem := len(buf)
	o := 0
	fmt.Println("Total remaining", rem)
	var err error
	for {
		if rem == 0 {
			break
		}

		higho := 0
		if (o + MaxFrameSize) > len(buf) {
			higho = len(buf)
		} else {
			higho = o + MaxFrameSize
		}

		fmt.Println("Encrypting ", o, "-", higho, "( ", len(buf), ")")

		z := buf[o:higho]
		zl := len(z)

		_, e := c.C.Write(z)
		if e != nil {
			err = e
			break
		}

		o = higho

		rem -= zl
	}

	c.errc <- err

	c.TmpFrames = nil
}

// ReadFrames reads one or more WoW protocol frames, properly decrypting headers and concatenating data
func (cl *Crypter) ReadFrames() ([][]byte, error) {
	var out [][]byte

	stillReading := false
	var remainingBytes int64
	tmp := new(bytes.Buffer)
	init := true
	/* Read bytes to temporary buffer in loop.
	 */

	ln := 2

	// if we are the server, we expect opcode lengths of 4
	if cl.Server {
		ln = 4
	}

	for {
		if stillReading == false && remainingBytes == 0 && init == false {
			break
		}

		init = false
		log.Println("yeee")
		buf := make([]byte, 4096)
		i, err := cl.C.Read(buf)
		if err != nil {
			return nil, err
		}
		log.Println("eeep")

		o := 0
		bu := buf[:i]

		if !stillReading {
			cl.K.Decrypt(bu[o : o+2+ln]) // Decrypt header
			remainingBytes = int64(binary.BigEndian.Uint16(bu[o : o+2]))
			tmp.Write(bu[o : o+2])
			o += 2
		}

		for {
			log.Println("eeeeee")

			// We finished reading from the current TCP packet, but there is still data reading to be done.
			if o == len(bu) && remainingBytes != 0 {
				stillReading = true
				break
			}

			// This packet contains additional data.
			if o < len(bu) && remainingBytes == 0 {
				// Include opcode in byte array.
				cl.K.Decrypt(bu[o : o+2+ln]) // Decrypt header
				remainingBytes = int64(binary.BigEndian.Uint16(bu[o : o+2]))
				tmp.Write(bu[o : o+2])
				o += 2
			}

			// We are done reading the TCP packet, and also we don't have to read anymore data for now.
			if remainingBytes == 0 {
				out = append(out, tmp.Bytes())
				tmp = new(bytes.Buffer)
				stillReading = false
				break
			}

			tmp.WriteByte(bu[o])
			o++
			remainingBytes--
		}
	}

	fmt.Println("EEE", WorldType(binary.LittleEndian.Uint16(out[0][2:4])))

	return out, nil
}
