package authserver

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"math/big"
	"net"
	"time"

	"github.com/superp00t/gophercraft/packet"
	"github.com/superp00t/gophercraft/srp"
)

type Account struct {
	Username     string
	IdentityHash []byte
}

type Config struct {
	Listen string

	GetAccount func(user string) *Account
	ListRealms func(user string) []packet.RealmListing
}

func (cfg *Config) Handle(c net.Conn) {
	buf := make([]byte, 512)
	rd, err := c.Read(buf)
	if err != nil {
		c.Close()
		return
	}

	alc, err := packet.UnmarshalAuthLogonChallenge_C(buf[:rd])
	if err != nil {
		c.Close()
		return
	}

	if alc.Build != 12340 {
		c.Write([]byte{
			uint8(packet.AUTH_LOGON_CHALLENGE),
			0x00,
			uint8(packet.WOW_FAIL_VERSION_INVALID),
		})

		time.Sleep(1 * time.Second)
		c.Close()
	}

	acc := cfg.GetAccount(string(alc.I))
	if acc == nil {
		// User could not be found.
		c.Write([]byte{
			uint8(packet.AUTH_LOGON_CHALLENGE),
			0x00,
			uint8(packet.WOW_FAIL_UNKNOWN_ACCOUNT),
		})
		time.Sleep(1 * time.Second)
		c.Close()
		return
	}

	nh := "894B645E89E1535BBDAD5B8B290650530801B18EBFBF5E8FAB3C82872A3E9BB7"
	nb, _ := hex.DecodeString(nh)
	N := &srp.BigNum{X: new(big.Int).SetBytes(nb)}
	v, s, _ := srp.ServerCalcVSX(acc.IdentityHash, N)

	g := srp.BigNumFromInt(7)
	b := srp.BigNumFromRand(19)
	gmod := g.ModExp(b, N)
	B := ((v.Multiply(srp.BigNumFromInt(3))).Add(gmod)).Mod(N)
	pkt := &packet.AuthLogonChallenge_S{
		Cmd:   packet.AUTH_LOGON_CHALLENGE,
		Error: packet.WOW_SUCCESS,
		B:     B.ToArray(),
		G:     7,
		N:     N.ToArray(),
		S:     s.ToArray(),
		Unk3:  srp.BigNumFromRand(16).ToArray(),
	}

	c.Write(pkt.Encode())
	pb := make([]byte, 512)
	c.Read(pb)

	alpc, err := packet.UnmarshalAuthLogonProof_C(pb)
	if err != nil {
		c.Close()
		return
	}

	_, valid, M3 := srp.ServerLogonProof(acc.Username,
		srp.BigNumFromArray(alpc.A),
		srp.BigNumFromArray(alpc.M1),
		b,
		B,
		s,
		N,
		v)

	if !valid {
		log.Println(acc.Username, "Invalid login")
		c.Write([]byte{
			uint8(packet.AUTH_LOGON_PROOF),
			uint8(packet.WOW_FAIL_UNKNOWN_ACCOUNT),
			0x00,
			0,
		})
		time.Sleep(1 * time.Second)
		c.Close()
		return
	}

	log.Println(acc.Username, "successfully authenticated")

	// flags := 0x00800000
	var flags uint32 = 8388608

	proof := &packet.AuthLogonProof_S{
		Cmd:          packet.AUTH_LOGON_PROOF,
		Error:        packet.WOW_SUCCESS,
		M2:           M3,
		AccountFlags: flags,
		SurveyID:     0,
		Unk3:         0,
	}

	_, err = c.Write(proof.Encode())
	if err != nil {
		c.Close()
		return
	}

	for {
		buff := make([]byte, 4)
		_, err = c.Read(buff)
		if err != nil {
			c.Close()
			return
		}

		hdr := packet.AuthType(buff[0])
		if hdr == packet.REALM_LIST {
			rls := cfg.ListRealms(acc.Username)
			rlst := &packet.RealmList_S{
				Cmd:    packet.REALM_LIST,
				Realms: rls,
			}
			_, err := c.Write(rlst.Encode())
			if err != nil {
				break
			}
		}
	}
	c.Close()
}

func Start(cfg *Config) error {
	l, err := net.Listen("tcp", cfg.Listen)
	if err != nil {
		return err
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go cfg.Handle(c)
	}
}

func rnd(l int) []byte {
	b := make([]byte, l)
	rand.Read(b)
	return b
}
