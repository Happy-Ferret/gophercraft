package packet

import (
	"crypto/hmac"
	"crypto/sha1"

	"github.com/superp00t/gophercraft/arc4"
)

type Crypt struct {
	ServerDecryptionKey []byte
	ServerEncryptionKey []byte

	DecClient, EncServer *arc4.Cipher
}

func NewCrypt(sessionKey []byte) *Crypt {
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
	// var err error
	// c.Cipher, err = rc4.NewCipher(key)
	// if err != nil {
	// 	panic(err)
	// }

	c.DecClient = arc4.ARC4(decryptHash)
	c.EncServer = arc4.ARC4(encryptHash)

	for i := 0; i < 1024; i++ {
		c.EncServer.Next()
		c.DecClient.Next()
	}
	// zeroBuffer := make([]byte, 1024)
	// syncBuffer := make([]byte, 1024)
	// c.EncServer.XORKeyStream(zeroBuffer, syncBuffer)
	// zeroBuffer = make([]byte, 1024)
	// syncBuffer = make([]byte, 1024)
	// c.DecClient.XORKeyStream(zeroBuffer, syncBuffer)

	return c
}

func (c *Crypt) Decrypt(input []byte) []byte {
	out := make([]byte, len(input))
	copy(out, input)
	c.DecClient.Decrypt(out)
	return out
}

func (c *Crypt) Encrypt(input []byte) []byte {
	out := make([]byte, len(input))
	copy(out, input)
	c.EncServer.Encrypt(out)
	return out
}
