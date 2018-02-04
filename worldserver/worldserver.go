package worldserver

import (
	"bytes"
	"crypto/rand"
	"net"
	"sync"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/superp00t/gophercraft/gcore"
	log "github.com/superp00t/gophercraft/gcore/glogger"
	"github.com/superp00t/gophercraft/packet"
)

type Config struct {
	Listen         string
	RealmName      string
	Source, Driver string
	PublicAddress  string
	WardenEnabled  bool
}

type WorldServer struct {
	Opts *Config
	DB   *xorm.Engine

	PhaseL *sync.Mutex
	Phases map[int64]*Phase
}

func Start(opt *Config) error {
	var err error
	ws := &WorldServer{}
	ws.Opts = opt
	ws.DB, err = xorm.NewEngine(opt.Driver, opt.Source)
	if err != nil {
		return err
	}

	// TODO: replace with pure DB solution
	go func() {
		c := gcore.Core{ws.DB}

		for {
			c.PublishRealmInfo(ws.Opts.RealmName, ws.Opts.PublicAddress, gcore.PVPServer)
			time.Sleep(9 * time.Second)
		}
	}()

	usrc, err := ws.DB.Where("realm = ?", opt.RealmName).Count(new(gcore.Character))
	if err != nil {
		return err
	}

	log.Println("Gophercraft Core World Server database opened without issue.")
	log.Println(usrc, " characters on this realm.")

	l, err := net.Listen("tcp", opt.Listen)
	if err != nil {
		return err
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		go ws.Handle(c)
	}
}

func (ws *WorldServer) Handle(c net.Conn) {
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

	buf := make([]byte, 512)
	wr, err := c.Read(buf)
	if err != nil {
		log.Println(err)
		c.Close()
		return
	}

	cmsg, err := packet.UnmarshalCMSGAuthSession(buf[:wr])
	if err != nil {
		log.Println("Invalid protocol: ", err)
		c.Close()
		return
	}

	var sessionKeys []gcore.SessionKey
	ws.DB.Where("account = ?", cmsg.Account).Find(&sessionKeys)
	if len(sessionKeys) == 0 {
		// Unauthenticated
		loginFail(c)
		return
	}

	digest := hash(
		[]byte(cmsg.Account),
		[]byte{0, 0, 0, 0},
		cmsg.Seed,
		salt,
		sessionKeys[0].K,
	)

	if !bytes.Equal(digest, cmsg.Digest) {
		log.Println("Cryptographic error")
		loginFail(c)
		return
	}

	crypt := packet.NewCrypter(c, sessionKeys[0].K, true)
	ssh := &Session{
		Account:    cmsg.Account,
		AddonData:  cmsg.AddonData,
		SessionKey: sessionKeys[0].K,
		WS:         ws,
		C:          c,
		Crypter:    crypt,
		Handlers:   make(map[packet.WorldType]*SessionHandler),

		ReadL: new(sync.Mutex), WriteL: new(sync.Mutex),
	}

	// for x := 1000; x > 1; x-- {
	// 	log.Println("Sent queue", x)
	// 	// size := 8
	// 	// szBuffer := make([]byte, 2)
	// 	// binary.BigEndian.PutUint16(szBuffer, uint16(size))
	// 	// tyBuffer := make([]byte, 2)
	// 	// binary.LittleEndian.PutUint16(tyBuffer, uint16(packet.SMSG_AUTH_RESPONSE))
	// 	// dtt := new(bytes.Buffer)
	// 	// dtt.Write(szBuffer)
	// 	// dtt.Write(tyBuffer)
	// 	// waitBuf := make([]byte, 4)
	// 	// dtt.WriteByte(packet.AUTH_WAIT_QUEUE)
	// 	// binary.LittleEndian.PutUint32(waitBuf, 120)
	// 	// dtt.Write(waitBuf)
	// 	// dtt.WriteByte(0)
	// 	// cut := ssh.Crypt.Encrypt(dtt.Bytes())
	// 	// _, err := ssh.C.Write(cut)
	// 	// log.Println("Queue len", len(cut))
	// 	// time.Sleep(5 * time.Second)

	// 	wait := packet.NewWorldPacket(packet.SMSG_AUTH_RESPONSE)
	// 	wait.Write([]byte{packet.AUTH_WAIT_QUEUE})
	// 	waitBuf := make([]byte, 4)
	// 	binary.LittleEndian.PutUint32(waitBuf, uint32(x)
	// 	wait.Write(waitBuf)
	// 	wait.Write([]byte{0})
	// 	_, err := ssh.Write(wait.Encode())
	// 	if err != nil {
	// 		ssh.C.Close()
	// 		return
	// 	}
	// 	time.Sleep(5 * time.Second)
	// }

	v := packet.NewWorldPacket(packet.SMSG_AUTH_RESPONSE)
	v.Write([]byte{packet.AUTH_OK})
	v.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0}) // random billing stuff
	v.Write([]byte{2})                         // Version: WotLK
	dt := v.Finish()
	ssh.WriteCrypt(dt)
	log.Println("Auth response sent,")

	if ssh.WS.Opts.WardenEnabled {
		ssh.InitWarden()
	}

	ssh.IntroductoryPackets()
	ssh.AddMenuHandlers()

	ssh.Handle()
}

func randomSalt(le int) []byte {
	salt := make([]byte, le)
	rand.Read(salt)
	return salt
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
	return packet.Hash(input...)
}
