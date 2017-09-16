package client

import (
	"fmt"
	"net"
	"strings"

	"github.com/superp00t/gophercraft/packet"
	"github.com/superp00t/gophercraft/srp"
)

type Config struct {
	Username, Password string
	AuthServer         string
}

type Client struct {
	RealmList *packet.RealmList_S
}

func Connect(cfg *Config) (*Client, error) {
	c := &Client{}

	cfg.Username = strings.ToUpper(cfg.Username)
	cfg.Password = strings.ToUpper(cfg.Password)

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

	A, M1 := srp.SRPCalculate(cfg.Username, cfg.Password, auth.B, auth.N, auth.S)

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
