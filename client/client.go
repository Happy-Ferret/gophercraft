package client

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"fmt"

	"net"

	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/superp00t/gophercraft/packet"
	"github.com/superp00t/gophercraft/srp"
	"github.com/superp00t/gophercraft/warden"
)

type Config struct {
	Username, Password string
	Playername         string
	AuthServer         string
}

type Client struct {
	Player     string
	Cfg        *Config
	RealmList  *packet.RealmList_S
	Warden     *warden.Warden
	World      net.Conn
	Crypter    *packet.Crypter
	SessionKey []byte
	OpCode     packet.WorldType
	Handlers   map[packet.WorldType]*ClientHandler
}

func AuthConnect(cfg *Config) (*Client, error) {
	c := &Client{}
	c.Player = cfg.Playername
	cfg.Username = strings.ToUpper(cfg.Username)
	cfg.Password = strings.ToUpper(cfg.Password)
	c.Cfg = cfg
	conn, err := net.Dial("tcp", cfg.AuthServer)
	if err != nil {
		return nil, err
	}

	body1 := packet.LogonChallengePacket_C(cfg.Username)
	conn.Write(body1)

	buf := make([]byte, 512)
	_, err = conn.Read(buf)
	if err != nil {
		return nil, err
	}

	auth, err := packet.UnmarshalAuthLogonChallenge_S(buf)
	if err != nil {
		return nil, err
	}

	if auth.Error != packet.WOW_SUCCESS {
		return nil, fmt.Errorf("Server returned %s", auth.Error)
	}

	_, K, A, M1 := srp.SRPCalculate(cfg.Username, cfg.Password, auth.B, auth.N, auth.S)
	c.SessionKey = K
	proof := &packet.AuthLogonProof_C{
		Cmd:          packet.AUTH_LOGON_PROOF,
		A:            A,
		M1:           M1,
		CRC:          make([]byte, 20),
		NumberOfKeys: 0,
		SecFlags:     0,
	}

	conn.Write(proof.Encode())

	buf = make([]byte, 512)
	_, err = conn.Read(buf)
	if err != nil {
		return nil, err
	}

	alps, err := packet.UnmarshalAuthLogonProof_S(buf)
	if err != nil {
		return nil, err
	}

	if alps.Error != packet.WOW_SUCCESS {
		return nil, fmt.Errorf("Server returned %s", alps.Error)
	}

	conn.Write(packet.RealmList_C)

	buf = make([]byte, 1024)
	conn.Read(buf)
	rls, _ := packet.UnmarshalRealmList_S(buf)
	c.RealmList = rls
	return c, nil
}

func (cl *Client) WorldConnect(ip string) error {
	wc, err := net.Dial("tcp", ip)
	if err != nil {
		return err
	}

	buf := make([]byte, 512)
	wc.Read(buf)
	gp, _ := packet.UnmarshalSMSGAuthPacket(buf)
	fmt.Println(spew.Sdump(gp))

	seed := randomBuffer(4)
	h := hash(
		[]byte(cl.Cfg.Username),
		[]byte{0, 0, 0, 0},
		seed,
		gp.Salt,
		cl.SessionKey,
	)

	app := packet.NewGamePacket(packet.CMSG_AUTH_SESSION)
	app.PutUint(12340)
	app.PutUint(0)
	app.PutCString(cl.Cfg.Username)
	app.PutUint(0)
	app.Buf.Write(seed)
	app.PutUint(0)
	app.PutUint(0)
	app.PutUint(0)
	app.PutUint(0)
	app.PutUint(0)
	app.Buf.Write(h)
	app.Buf.Write(packet.ClientAddonData)
	// Addon data
	pkt := app.Finish()
	wc.Write(pkt)

	cl.Handlers = make(map[packet.WorldType]*ClientHandler)
	cl.World = wc
	cl.Crypter = packet.NewCrypter(wc, cl.SessionKey, false)
	go cl.Handle()
	// for {
	// 	wp, _ := packet.DecodeWorldPacket(cl.ReadCrypt())
	// 	log.Println("got", wp.Type)
	// 	if wp.Type == packet.SMSG_WARDEN_DATA {
	// 		log.Println("warden detected")
	// 	}
	// }
	// buff := make([]byte, 512)
	// wre, _ := wc.Read(buff)
	// cbuf := cl.Crypt.Decrypt(buff[:wre])l;
	// typ := packet.WorldType(binary.LittleEndian.Uint16(cbuf[2:4]))
	// log.Println("Got ", typ)
	return nil
}

func (cl *Client) SendCrypt(data []byte) {
	cl.Crypter.SendFrame(data)
}

func (cl *Client) ReadCrypt() [][]byte {
	bufs, err := cl.Crypter.ReadFrames()
	if err != nil {
		panic(err)
	}

	return bufs
}

// func (cl *Client) ReadCrypt() [][]byte {
// 	bff := make([]byte, 8192)
// 	i, err := cl.World.Read(bff)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	glogger.Println("Read ", i, "bytes")
// 	buf := bff[:i]

// 	p := packet.NewEtcBuffer(buf)
// 	var pkts [][]byte
// 	for {
// 		if p.GetRPos() >= len(buf) {
// 			break
// 		} else {
// 			glogger.Println("Continuing", p.GetRPos(), len(buf))
// 		}

// 		if cl.Remaining > 0 {
// 			start := cl.DataBuffer.Len()
// 			glogger.Println("Reading remaining byte range", cl.Remaining)
// 			rmndr := len(buf) - p.GetRPos()
// 			if cl.Remaining > rmndr {
// 				cl.DataBuffer.Write(p.ReadBytes(rmndr))
// 			} else {
// 				cl.DataBuffer.Write(p.ReadBytes(cl.Remaining - rmndr))
// 			}

// 			cl.Remaining -= cl.DataBuffer.Len() - start
// 			if cl.Remaining == 0 {
// 				pkts = append(pkts, cl.DataBuffer.Bytes())
// 				cl.DataBuffer = new(bytes.Buffer)
// 			}
// 			continue
// 		}

// 		hBuf := p.ReadBytes(4)
// 		hBff := cl.Crypt.Decrypt(hBuf)
// 		lng := binary.BigEndian.Uint16(hBff[0:2])
// 		op := packet.WorldType(binary.LittleEndian.Uint16(hBff[2:4]))
// 		glogger.Println("Decoded op", op)
// 		// if (p.GetRPos() + int(lng)) > len(buf) {
// 		// 	lng = uint16(p.Len() - p.GetRPos())
// 		// }

// 		glogger.Println("Reading", lng)
// 		id := int(lng) - 2
// 		if p.GetRPos()+id > p.Len() {
// 			tBuf := p.ReadRemainingBytes()
// 			cl.DataBuffer.Write(tBuf)
// 			cl.Remaining = int(lng) - len(tBuf) - 2
// 			break
// 		}
// 		data := p.ReadBytes(id)
// 		pkts = append(pkts, append(hBff, data...))
// 	}

// 	return pkts
// }

func hash(input ...[]byte) []byte {
	bt := sha1.Sum(bytes.Join(input, nil))
	return bt[:]
}

func randomBuffer(l int) []byte {
	b := make([]byte, l)
	rand.Read(b)
	return b
}

// func (cl *Client) ReadCrypt() [][]byte {
// 	bff := make([]byte, 16000)
// 	i, err := cl.World.Read(bff)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	glogger.Println("Read ", i, "bytes")
// 	buf := bff[:i]
// 	p := packet.NewEtcBuffer(buf)

// 	var pkts [][]byte

// 	for {
// 		if cl.Remaining == 0 && cl.DataBuffer.Len() > 0 {
// 			pkts = append(pkts, cl.DataBuffer.Bytes())
// 			cl.DataBuffer = new(bytes.Buffer)
// 		}
// 		if p.Available() == 0 {
// 			break
// 		}
// 		if cl.Remaining == 0 {
// 			// Header buffer
// 			hBuf := p.ReadBytes(4)
// 			hBff := packet.NewEtcBuffer(cl.Crypt.Decrypt(hBuf))
// 			cl.Remaining = hBff.ReadBigUint16() - 2 // Minus 2 because we read the opcode
// 			op := packet.WorldType(hBff.ReadUint16())
// 			glogger.Warnln("Recv opcode", op)
// 			cl.DataBuffer.Write(hBff.Encode())
// 		}

// 		if cl.Remaining > 0 {
// 			cl.DataBuffer.WriteByte(p.ReadByte())
// 			cl.Remaining--
// 		}

// 	}
// 	glogger.Warnln("Return", cl.Remaining, p.Available(), cl.DataBuffer.Len())
// 	return pkts
// }
