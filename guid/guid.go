package guid

import (
	"fmt"
)

type GUID uint64

const (
	Item          GUID = 0x4000 // blizz 4000
	Container     GUID = 0x4000 // blizz 4000
	Player        GUID = 0x0000 // blizz 0000
	GameObject    GUID = 0xF110 // blizz F110
	Transport     GUID = 0xF120 // blizz F120 (for GAMEOBJECT_TYPE_TRANSPORT)
	Unit          GUID = 0xF130 // blizz F130
	Pet           GUID = 0xF140 // blizz F140
	Vehicle       GUID = 0xF150 // blizz F550
	DynamicObject GUID = 0xF100 // blizz F100
	Corpse        GUID = 0xF101 // blizz F100
	Mo_Transport  GUID = 0x1FC0 // blizz 1FC0 (for GAMEOBJECT_TYPE_MO_TRANSPORT)
	Instance      GUID = 0x1F40 // blizz 1F40
	Group         GUID = 0x1F50
)

func (g GUID) Low() GUID {
	return g & 0xFFFFFFFFFFFF
}

func (g GUID) High() GUID {
	return (g >> 48) & 0x0000FFFF
}

func (g GUID) SetLow(l GUID) GUID {
	return (g.High() << 48) | l
}

func (g GUID) SetHigh(h GUID) GUID {
	return (h << 48) | g.Low()
}

func (g GUID) EncodePacked() []byte {
	guid := g
	packGUID := make([]byte, 9)
	packGUID[0] = 0
	size := 1

	var i uint8
	for i = 0; guid != 0; i++ {
		if (guid & 0xFF) > 0 {
			packGUID[0] |= uint8(1 << i)
			packGUID[size] = uint8(guid & 0xFF)
			size++
		}

		guid >>= 8
	}

	if guid == 0 {
		size++
		packGUID[size] = 0
	}

	return packGUID[:size]
}

func (g GUID) String() string {
	t := ""
	switch g.High() {
	case Item:
		t = "Item"
	case Player:
		t = "Player"
	case GameObject:
		t = "GameObject"
	case Transport:
		t = "Transport"
	case Unit:
		t = "Unit"
	case Pet:
		t = "Pet"
	case Vehicle:
		t = "Vehicle"
	case DynamicObject:
		t = "DynamicObject"
	case Corpse:
		t = "Corpse"
	case Mo_Transport:
		t = "Mo_Transport"
	case Instance:
		t = "Instance"
	case Group:
		t = "Group"
	default:
		t = fmt.Sprintf("0x%X", g.High().U64())
	}

	return fmt.Sprintf("(%s) 0x%016X", t, g.U64())
}

func (g GUID) U64() uint64 {
	return uint64(g)
}
