package worldserver

import (
	"bytes"
	"encoding/binary"

	log "github.com/superp00t/gophercraft/gcore/glogger"
	"github.com/superp00t/gophercraft/guid"

	"github.com/superp00t/gophercraft/gcore"
	"github.com/superp00t/gophercraft/packet"
)

func (s *Session) GetNewPlayerGUID() guid.GUID {
	var i []gcore.Character
	s.WS.DB.OrderBy("g_u_i_d").Asc().Limit(1).Find(&i)
	var g guid.GUID = 0
	if len(i) != 0 {
		g = i[0].GUID + 1
	}

	return g
}

func (s *Session) CharacterList(b []byte) {
	log.Println("Character list requested")

	var chars []gcore.Character
	var pktChars []packet.Character
	s.WS.DB.Where("account = ?", s.Account).Where("realm = ?", s.WS.Opts.RealmName).Find(&chars)
	for _, v := range chars {
		log.Println("Fetching", v.Name, "from database ", v.GUID)

		pktChars = append(pktChars, packet.Character{
			GUID:       v.GUID,
			Name:       v.Name,
			Race:       v.Race,
			Class:      v.Class,
			Gender:     v.Gender,
			Skin:       v.Skin,
			Face:       v.Face,
			HairStyle:  v.HairStyle,
			HairColor:  v.HairColor,
			FacialHair: v.FacialHair,
			Level:      v.Level,
			Zone:       87,
			Map:        0,
			X:          -9460.25,
			Y:          63.0612,
			Z:          56.8335,
			Equipment:  make([]packet.Item, 24),
		})
	}

	s.CharList = pktChars
	// list := &packet.CharacterList{
	// 	Characters: []packet.Character{
	// 	// packet.Character{
	// 	// 	GUID:   randomSalt(8),
	// 	// 	Name:   "Kleinem",
	// 	// 	Race:   packet.RACE_HUMAN,
	// 	// 	Class:  packet.CLASS_WARRIOR,
	// 	// 	Gender: 0,
	// 	// 	Bytes:  0,
	// 	// 	Facial: 3,
	// 	// 	Level:  10,
	// 	// 	Zone:   3,
	// 	// 	Map:    2,

	// 	// 	Equipment: make([]packet.Item, 24),
	// 	// },
	// 	},
	// }
	list := &packet.CharacterList{
		Characters: pktChars,
	}

	s.WriteCrypt(list.Encode())
}

// type Character struct {
// 	GUID       guid.GUID
// 	Name       string
// 	Race       uint8
// 	Class      uint8
// 	Gender     uint8
// 	Skin       uint8
// 	Face       uint8
// 	HairStyle  uint8
// 	HairColor  uint8
// 	FacialHair uint8
// }

func (s *Session) SendCharCreate(opcode uint8) {
	pkt := packet.NewWorldPacket(packet.SMSG_CHAR_CREATE)
	pkt.Write([]byte{opcode})
	s.WriteCrypt(pkt.Finish())
}

func (s *Session) DeleteCharacter(b []byte) {
	guid := binary.LittleEndian.Uint64(b[6:14])

	_, err := s.WS.DB.Where("g_u_i_d = ?", guid).Delete(new(gcore.Character))
	if err != nil {
		log.Fatal(err)
	}
	pkt := packet.NewWorldPacket(packet.SMSG_CHAR_DELETE)
	pkt.Write([]byte{CHAR_DELETE_SUCCESS})
	s.WriteCrypt(pkt.Finish())
}

func (s *Session) CreateCharacter(b []byte) {
	o := 6
	nameb := new(bytes.Buffer)
	for {
		if b[o] == 0 {
			break
		}

		nameb.WriteByte(b[o])
		o++
	}
	o++

	name := nameb.String()
	if name == "" {
		s.SendCharCreate(CHAR_NAME_NO_NAME)
		return
	}

	// Check if character name is in use
	var chars []gcore.Character
	s.WS.DB.Where("name = ?", name).Find(&chars)
	if len(chars) != 0 {
		s.SendCharCreate(CHAR_CREATE_NAME_IN_USE)
		return
	}

	log.Println("Registered name: ", name)
	pch := &gcore.Character{}
	pch.Account = s.Account
	pch.Realm = s.WS.Opts.RealmName
	pch.GUID = s.GetNewPlayerGUID()
	pch.Name = name
	pch.Race = b[o]
	o++
	pch.Class = b[o]
	o++
	pch.Gender = b[o]
	o++
	pch.Skin = b[o]
	o++
	pch.Face = b[o]
	o++
	pch.HairStyle = b[o]
	o++
	pch.HairColor = b[o]
	o++
	pch.FacialHair = b[o]
	pch.Level = 1
	_, err := s.WS.DB.Insert(pch)
	if err != nil {
		log.Fatal(err)
	}
	s.SendCharCreate(CHAR_CREATE_SUCCESS)
}

const (
	CHAR_CREATE_IN_PROGRESS                                uint8 = 0x2E
	CHAR_CREATE_SUCCESS                                    uint8 = 0x2F
	CHAR_CREATE_ERROR                                      uint8 = 0x30
	CHAR_CREATE_FAILED                                     uint8 = 0x31
	CHAR_CREATE_NAME_IN_USE                                uint8 = 0x32
	CHAR_CREATE_DISABLED                                   uint8 = 0x33
	CHAR_CREATE_PVP_TEAMS_VIOLATION                        uint8 = 0x34
	CHAR_CREATE_SERVER_LIMIT                               uint8 = 0x35
	CHAR_CREATE_ACCOUNT_LIMIT                              uint8 = 0x36
	CHAR_CREATE_SERVER_QUEUE                               uint8 = 0x37
	CHAR_CREATE_ONLY_EXISTING                              uint8 = 0x38
	CHAR_CREATE_EXPANSION                                  uint8 = 0x39
	CHAR_CREATE_EXPANSION_CLASS                            uint8 = 0x3A
	CHAR_CREATE_LEVEL_REQUIREMENT                          uint8 = 0x3B
	CHAR_CREATE_UNIQUE_CLASS_LIMIT                         uint8 = 0x3C
	CHAR_CREATE_CHARACTER_IN_GUILD                         uint8 = 0x3D
	CHAR_CREATE_RESTRICTED_RACECLASS                       uint8 = 0x3E
	CHAR_CREATE_CHARACTER_CHOOSE_RACE                      uint8 = 0x3F
	CHAR_CREATE_CHARACTER_ARENA_LEADER                     uint8 = 0x40
	CHAR_CREATE_CHARACTER_DELETE_MAIL                      uint8 = 0x41
	CHAR_CREATE_CHARACTER_SWAP_FACTION                     uint8 = 0x42
	CHAR_CREATE_CHARACTER_RACE_ONLY                        uint8 = 0x43
	CHAR_CREATE_CHARACTER_GOLD_LIMIT                       uint8 = 0x44
	CHAR_CREATE_FORCE_LOGIN                                uint8 = 0x45
	CHAR_NAME_SUCCESS                                      uint8 = 0x57
	CHAR_NAME_FAILURE                                      uint8 = 0x58
	CHAR_NAME_NO_NAME                                      uint8 = 0x59
	CHAR_NAME_TOO_SHORT                                    uint8 = 0x5A
	CHAR_NAME_TOO_LONG                                     uint8 = 0x5B
	CHAR_NAME_INVALID_CHARACTER                            uint8 = 0x5C
	CHAR_NAME_MIXED_LANGUAGES                              uint8 = 0x5D
	CHAR_NAME_PROFANE                                      uint8 = 0x5E
	CHAR_NAME_RESERVED                                     uint8 = 0x5F
	CHAR_NAME_INVALID_APOSTROPHE                           uint8 = 0x60
	CHAR_NAME_MULTIPLE_APOSTROPHES                         uint8 = 0x61
	CHAR_NAME_THREE_CONSECUTIVE                            uint8 = 0x62
	CHAR_NAME_INVALID_SPACE                                uint8 = 0x63
	CHAR_NAME_CONSECUTIVE_SPACES                           uint8 = 0x64
	CHAR_NAME_RUSSIAN_CONSECUTIVE_SILENT_CHARACTERS        uint8 = 0x65
	CHAR_NAME_RUSSIAN_SILENT_CHARACTER_AT_BEGINNING_OR_END uint8 = 0x66
	CHAR_NAME_DECLENSION_DOESNT_MATCH_BASE_NAME            uint8 = 0x67

	CHAR_DELETE_IN_PROGRESS                uint8 = 70
	CHAR_DELETE_SUCCESS                    uint8 = 71
	CHAR_DELETE_FAILED                     uint8 = 72
	CHAR_DELETE_FAILED_LOCKED_FOR_TRANSFER uint8 = 73
	CHAR_DELETE_FAILED_GUILD_LEADER        uint8 = 74
	CHAR_DELETE_FAILED_ARENA_CAPTAIN       uint8 = 75
)
