package worldserver

import (
	"encoding/binary"
	"net"
	"sync"

	log "github.com/superp00t/gophercraft/gcore/glogger"
	"github.com/superp00t/gophercraft/guid"
	"github.com/superp00t/gophercraft/warden"

	"github.com/superp00t/gophercraft/packet"
)

type Session struct {
	Account    string
	SessionKey []byte
	AddonData  []byte

	Warden *warden.Warden
	WS     *WorldServer

	C             net.Conn
	Crypter       *packet.Crypter
	ReadL, WriteL *sync.Mutex
	Char          *packet.Character

	CharList []packet.Character
	Handlers map[packet.WorldType]*SessionHandler
}

type SessionHandler struct {
	Type packet.WorldType
	Fn   func([]byte)
}

func (s *Session) ReadCrypt() ([][]byte, error) {
	return s.Crypter.ReadFrames()
}

func (s *Session) WriteCrypt(buf []byte) error {
	return s.Crypter.SendFrame(buf[2:])
}

func (s *Session) Send(p *packet.WorldPacket) {
	s.WriteCrypt(p.Finish())
}

func (s *Session) HandleFunc(t packet.WorldType, fn func([]byte)) {
	s.Handlers[t] = &SessionHandler{
		Type: t,
		Fn:   fn,
	}
}

func (s *Session) HandlePong(buf []byte) {
	ping := binary.LittleEndian.Uint32(buf[6:10])
	latency := binary.LittleEndian.Uint32(buf[10:14])
	log.Println("Ping: ", ping, "Latency", latency)
	pkt := packet.NewWorldPacket(packet.SMSG_PONG)
	pkt.Write(buf[6:10])
	s.WriteCrypt(pkt.Finish())
}

func (s *Session) HandleJoin(buf []byte) {
	log.Println("Player join requested")
	guid := guid.GUID(binary.LittleEndian.Uint64(buf[6:14]))
	for _, v := range s.CharList {
		if v.GUID == guid {
			s.Char = &v
			log.Println("GUID found for character", v.Name, v.GUID)
			s.SendInitialPackets()
			return
		}
	}

	// Todo handle unknown GUID
}

func (s *Session) AddMenuHandlers() {
	s.HandleFunc(packet.CMSG_CHAR_ENUM, s.CharacterList)
	s.HandleFunc(packet.CMSG_CHAR_DELETE, s.DeleteCharacter)
	s.HandleFunc(packet.CMSG_CHAR_CREATE, s.CreateCharacter)
	s.HandleFunc(packet.CMSG_PLAYER_LOGIN, s.HandleJoin)
}

func (s *Session) Handle() {
	s.HandleFunc(packet.CMSG_WARDEN_DATA, s.WardenResponse)

	s.HandleFunc(packet.CMSG_PING, s.HandlePong)

	for {
		bufs, err := s.ReadCrypt()
		if err != nil {
			log.Println(err)
			return
		}

		for _, buf := range bufs {
			t := packet.WorldType(binary.LittleEndian.Uint32(buf[2:6]))
			log.Println(t, "requested", len(buf))

			if h := s.Handlers[t]; h != nil {
				h.Fn(buf)
			}
		}
	}
}
