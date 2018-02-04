package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/ogier/pflag"
	"github.com/superp00t/gophercraft/authserver"
	"github.com/superp00t/gophercraft/client"
	log "github.com/superp00t/gophercraft/gcore/glogger"
	"github.com/superp00t/gophercraft/packet"
	"github.com/superp00t/gophercraft/srp"
)

var (
	cl *client.Client
	// This is what we use to communicate with the WoW client.
	clientKey []byte
	remoteIP  string
	// This is what we use to communicate with the remote WoW server.
	serverKey []byte

	realmName = pflag.StringP("realm_name", "n", "Icecrown", "the name of the realm to connect to")

	user = pflag.StringP("user", "u", "user", "WoW account username")

	pass       = pflag.StringP("pass", "p", "password", "WoW account password")
	authServer = pflag.StringP("authserver_addr", "r", "logon.warmane.com:3724", "IP address:port of auth server")
)

func GetAccount(usern string) *authserver.Account {
	pUser := strings.ToUpper(*user)
	pPass := strings.ToUpper(*pass)

	if usern != pUser {
		return nil
	}

	return &authserver.Account{
		Username:     pUser,
		IdentityHash: srp.HashCredentials(pUser, pPass),
	}
}

func StoreKey(user string, K []byte) {
	clientKey = K
}

func ListRealms(usern string) []packet.RealmListing {
	var rlst []packet.RealmListing
	if cl == nil {
		var err error
		cl, err = client.AuthConnect(&client.Config{
			Username:   *user,
			Password:   *pass,
			AuthServer: *authServer,
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	for _, v := range cl.RealmList.Realms {
		if v.Name == *realmName {
			remoteIP = v.Address
			v.Address = "localhost:8085"
		}

		rlst = append(rlst, v)
	}

	return rlst
}
func randomSalt(le int) []byte {
	salt := make([]byte, le)
	rand.Read(salt)
	return salt
}

func handleWorld(c net.Conn) {
	log.Println("New worldserver connection from", c.RemoteAddr().String())

	// Authentication
	salt := randomSalt(4)

	smsg := &packet.SMSGAuthPacket{
		Salt:  salt,
		Seed1: randomSalt(16),
		Seed2: randomSalt(16),
	}

	dat := smsg.Encode()
	c.Write(dat)

	// dp, _ := packet.UnmarshalSMSGAuthPacket(dat)
	// log.Fatal(bytes.Equal(dp.Salt, salt))

	buf := make([]byte, 512)
	wr, err := c.Read(buf)
	if err != nil {
		c.Close()
		return
	}

	cmsg, err := packet.UnmarshalCMSGAuthSession(buf[:wr])
	if err != nil {
		log.Println("Invalid protocol: ", err)
		c.Close()
		return
	}

	acc := GetAccount(cmsg.Account)
	if acc == nil {
		loginFail(c)
		return
	}

	digest := hash(
		[]byte(cmsg.Account),
		[]byte{0, 0, 0, 0},
		cmsg.Seed,
		salt,
		clientKey,
	)

	if !bytes.Equal(digest, cmsg.Digest) {
		log.Println("Cryptographic error")
		loginFail(c)
		return
	}

	for _, v := range cmsg.AddonData {
		fmt.Printf("0x%X, ", v)
	}
	fmt.Println()

	log.Println("Huge success!")
	err = cl.WorldConnect(remoteIP)
	if err != nil {
		log.Fatal(err)
	}

	crypt := packet.NewCrypt(clientKey, true)

	go func() {
		for {
			// Read packets from server
			data := <-cl.Steal
			op := binary.LittleEndian.Uint16(data[2:4])
			opc := packet.WorldType(op)
			log.Println("S => ", opc)
			if strings.HasPrefix(opc.String(), "WorldType") {
				log.Println("Invalid packet: ", data)
			}

			if opc == packet.SMSG_WARDEN_DATA {
				continue
			}

			// Send to client
			rHeader := data[:4]
			// if len(data) > 899 {
			// 	headerLen := len(data[4:]) + 2
			// 	headerOp := packet.SMSG_CHAR_ENUM
			// 	rHeader = make([]byte, 4)
			// 	binary.BigEndian.PutUint16(rHeader[:2], uint16(headerLen))
			// 	binary.LittleEndian.PutUint16(rHeader[2:4], uint16(headerOp))
			// 	log.Println("Raw header", rHeader)
			// }

			header := crypt.Encrypt(rHeader)
			_, err := c.Write(append(header, data[4:]...))
			if err != nil {

				log.Fatal("Cant write to client", err)
			}
		}
	}()

	for {
		// Read packets from client
		buff := make([]byte, 1024)
		rd, err := c.Read(buff)
		if err != nil {
			log.Println(err)
			break
		}

		buf := buff[:rd]

		header := crypt.Decrypt(buf[:6])
		op := binary.LittleEndian.Uint16(header[2:6])
		log.Println("C => ", packet.WorldType(op))
		// Send to server
		cl.SendCrypt(append(header, buf[6:]...))
	}
}

func loginFail(c net.Conn) {
	log.Println("Login failure")
	wp := packet.NewWorldPacket(packet.SMSG_AUTH_RESPONSE)
	wp.Write([]byte{packet.AUTH_REJECT})
	c.Write(wp.Encode())
	time.Sleep(1 * time.Second)
	c.Close()
	return

}

func hash(input ...[]byte) []byte {
	bt := sha1.Sum(bytes.Join(input, nil))
	return bt[:]
}

func main() {
	pflag.Parse()

	go func() {
		l, err := net.Listen("tcp", "localhost:8085")
		if err != nil {
			log.Fatal(err)
		}

		for {
			c, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}

			go handleWorld(c)
		}
	}()

	log.Fatal(authserver.Start(&authserver.Config{
		Listen:     "localhost:3724",
		GetAccount: GetAccount,
		ListRealms: ListRealms,
		StoreKey:   StoreKey,
	}))

}
