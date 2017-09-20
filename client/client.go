package client

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/superp00t/gophercraft/packet"
	"github.com/superp00t/gophercraft/srp"
)

type Config struct {
	Username, Password string
	AuthServer         string
}

type Client struct {
	Cfg       *Config
	RealmList *packet.RealmList_S

	World      net.Conn
	Crypt      *packet.Crypt
	SessionKey []byte
}

func AuthConnect(cfg *Config) (*Client, error) {
	c := &Client{}

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
	app.PutUint(0)
	// Addon data
	pkt := app.Finish()
	wc.Write(pkt)

	cl.World = wc
	cl.Crypt = packet.NewCrypt(cl.SessionKey)
	// for {
	// 	wp, _ := packet.DecodeWorldPacket(cl.ReadCrypt())
	// 	log.Println("got", wp.Type)
	// 	if wp.Type == packet.SMSG_WARDEN_DATA {
	// 		log.Println("warden detected")
	// 	}
	// }

	return nil
}

func (cl *Client) ReadCrypt() []byte {
	buf := make([]byte, 1024)
	i, err := cl.World.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(buf[:i])
	return cl.Crypt.Decrypt(buf[:i])
}

func hash(input ...[]byte) []byte {
	bt := sha1.Sum(bytes.Join(input, nil))
	return bt[:]
}

func randomBuffer(l int) []byte {
	b := make([]byte, l)
	rand.Read(b)
	return b
}
