package packet

import "github.com/superp00t/gophercraft/guid"

const (
	RACE_NONE               uint8 = 0
	RACE_HUMAN              uint8 = 1
	RACE_ORC                uint8 = 2
	RACE_DWARF              uint8 = 3
	RACE_NIGHTELF           uint8 = 4
	RACE_UNDEAD_PLAYER      uint8 = 5
	RACE_TAUREN             uint8 = 6
	RACE_GNOME              uint8 = 7
	RACE_TROLL              uint8 = 8
	RACE_GOBLIN             uint8 = 9
	RACE_BLOODELF           uint8 = 10
	RACE_DRAENEI            uint8 = 11
	RACE_FEL_ORC            uint8 = 12
	RACE_NAGA               uint8 = 13
	RACE_BROKEN             uint8 = 14
	RACE_SKELETON           uint8 = 15
	RACE_VRYKUL             uint8 = 16
	RACE_TUSKARR            uint8 = 17
	RACE_FOREST_TROLL       uint8 = 18
	RACE_TAUNKA             uint8 = 19
	RACE_NORTHREND_SKELETON uint8 = 20
	RACE_ICE_TROLL          uint8 = 21
	RACE_WORGEN             uint8 = 22
	RACE_GILNEAN            uint8 = 23
	RACE_PANDAREN_NEUTRAL   uint8 = 24
	RACE_PANDAREN_ALLIANCE  uint8 = 25
	RACE_PANDAREN_HORDE     uint8 = 26
)

const (
	CLASS_NONE         uint8 = 0
	CLASS_WARRIOR      uint8 = 1
	CLASS_PALADIN      uint8 = 2
	CLASS_HUNTER       uint8 = 3
	CLASS_ROGUE        uint8 = 4
	CLASS_PRIEST       uint8 = 5
	CLASS_DEATH_KNIGHT uint8 = 6
	CLASS_SHAMAN       uint8 = 7
	CLASS_MAGE         uint8 = 8
	CLASS_WARLOCK      uint8 = 9
	CLASS_MONK         uint8 = 10
	CLASS_DRUID        uint8 = 11
	CLASS_DEMON_HUNTER uint8 = 12
)

type Character struct {
	GUID       guid.GUID
	Name       string
	Race       uint8
	Class      uint8
	Gender     uint8
	Skin       uint8
	Face       uint8
	HairStyle  uint8
	HairColor  uint8
	FacialHair uint8
	Level      uint8
	Zone       uint32
	Map        uint32
	X, Y, Z    float32
	Guild      uint32
	Flags      uint32

	Customization uint32
	Unk           uint8

	PetModel, PetLevel, PetFamily uint32

	Equipment []Item
}

type Item struct {
	Model       uint32
	Type        uint8
	Enchantment uint32
}

type CharacterList struct {
	Characters []Character
}

func (c *CharacterList) Encode() []byte {
	p := NewWorldPacket(SMSG_CHAR_ENUM)
	p.WriteByte(uint8(len(c.Characters)))
	for _, v := range c.Characters {
		p.WriteUint64(v.GUID.U64())
		p.WriteCString(v.Name)
		p.WriteByte(v.Race)
		p.WriteBytes([]byte{v.Class})
		p.WriteBytes([]byte{v.Gender})
		p.WriteBytes([]byte{v.Skin, v.Face, v.HairStyle, v.HairColor})
		p.WriteBytes([]byte{v.FacialHair})
		p.WriteBytes([]byte{v.Level})
		p.WriteUint32(v.Zone)
		p.WriteUint32(v.Map)
		p.WriteFloat32(v.X)
		p.WriteFloat32(v.Y)
		p.WriteFloat32(v.Z)
		p.WriteUint32(v.Guild)
		p.WriteUint32(v.Flags)
		p.WriteUint32(v.Customization)
		p.WriteBytes([]byte{v.Unk})
		p.WriteUint32(v.PetModel)
		p.WriteUint32(v.PetLevel)
		p.WriteUint32(v.PetFamily)

		for i := 0; i < 23; i++ {
			item := v.Equipment[i]
			p.WriteUint32(item.Model)
			p.WriteBytes([]byte{item.Type})
			p.WriteUint32(item.Enchantment)
		}
	}
	return p.Finish()
}

func UnmarshalCharacterList(input []byte) (*CharacterList, error) {
	pkt := NewEtcBuffer(input)
	pkt.RPos(4)
	count := int(pkt.ReadByte())
	var chh CharacterList
	for x := 0; x < count; x++ {
		ch := Character{}
		ch.GUID = guid.GUID(pkt.ReadUint64())
		ch.Name = pkt.ReadCString()
		ch.Race = pkt.ReadByte()
		ch.Class = pkt.ReadByte()
		ch.Gender = pkt.ReadByte()
		ch.Skin = pkt.ReadByte()
		ch.Face = pkt.ReadByte()
		ch.HairStyle = pkt.ReadByte()
		ch.HairColor = pkt.ReadByte()
		ch.FacialHair = pkt.ReadByte()
		ch.Level = pkt.ReadByte()
		ch.Zone = pkt.ReadUint32()
		ch.Map = pkt.ReadUint32()
		ch.X = pkt.ReadFloat32()
		ch.Y = pkt.ReadFloat32()
		ch.Z = pkt.ReadFloat32()
		ch.Guild = pkt.ReadUint32()
		ch.Flags = pkt.ReadUint32()
		ch.Customization = pkt.ReadUint32()
		ch.Unk = pkt.ReadByte()
		ch.PetModel = pkt.ReadUint32()
		ch.PetLevel = pkt.ReadUint32()
		ch.PetFamily = pkt.ReadUint32()

		// Get equipment
		for j := 0; j < 23; j++ {
			model := pkt.ReadUint32()
			typ := pkt.ReadByte()
			enchant := pkt.ReadUint32()
			item := Item{
				Model:       model,
				Type:        typ,
				Enchantment: enchant,
			}
			ch.Equipment = append(ch.Equipment, item)
		}

		chh.Characters = append(chh.Characters, ch)
	}
	return &chh, nil
}
