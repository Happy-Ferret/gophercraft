package packet

import (
	"bytes"
	"encoding/binary"
	"math"
)

var RealmList_C = []byte{
	uint8(REALM_LIST),
	0x00,
	0x00,
	0x00,
	0x00,
}

type RealmList_S struct {
	Cmd AuthType

	Realms []RealmListing
}

type RealmListing struct {
	Type       uint8 //
	Color      uint8
	Locked     bool
	Flags      uint8
	Name       string
	Address    string
	Population float32
	Characters uint8
	Timezone   uint8
	ID         uint8
}

func (rlst *RealmList_S) Encode() []byte {

	listBuffer := new(bytes.Buffer)

	for _, v := range rlst.Realms {
		listBuffer.WriteByte(v.Type)
		listBuffer.WriteByte(v.Color)
		var lck uint8
		if v.Locked {
			lck++
		}
		listBuffer.WriteByte(lck)
		listBuffer.Write(append([]byte(v.Name), 0))
		listBuffer.Write(append([]byte(v.Address), 0))
		pop := make([]byte, 4)
		u := math.Float32bits(v.Population)
		binary.LittleEndian.PutUint32(pop, u)
		listBuffer.Write(pop)
		listBuffer.WriteByte(v.Characters)
		listBuffer.WriteByte(v.Timezone)
		listBuffer.WriteByte(v.ID)
	}

	listBuffer.WriteByte(0x10)
	listBuffer.WriteByte(0x00)

	head := new(bytes.Buffer)
	head.WriteByte(uint8(rlst.Cmd))

	sb := new(bytes.Buffer)

	request := make([]byte, 4)
	binary.LittleEndian.PutUint32(request, 0)
	sb.Write(request)

	realmCount := make([]byte, 2)
	binary.LittleEndian.PutUint16(realmCount, uint16(len(rlst.Realms)))
	sb.Write(realmCount)

	sizeBuf := make([]byte, 2)
	binary.LittleEndian.PutUint16(sizeBuf, uint16(listBuffer.Len()+sb.Len()))
	head.Write(sizeBuf)
	head.Write(sb.Bytes())
	head.Write(listBuffer.Bytes())

	// request := make([]byte, 4)
	// realmsCount := make([]byte, 2)
	// szbuf := make([]byte, 2)
	// binary.LittleEndian.PutUint32(request, uint32(0))
	// binary.LittleEndian.PutUint16(realmsCount, uint16(len(rlst.Realms)))
	// size := uint16(listBuffer.Len() + 10)
	// binary.LittleEndian.PutUint16(szbuf, size)
	// sizeBuffer.Write(szbuf)
	// sizeBuffer.Write(request)
	// sizeBuffer.Write(realmsCount)
	// head.Write(sizeBuffer.Bytes())
	// head.Write(listBuffer.Bytes())
	return head.Bytes()
}

func UnmarshalRealmList_S(input []byte) (*RealmList_S, error) {
	rls := &RealmList_S{}
	rls.Cmd = AuthType(input[0])
	size := binary.LittleEndian.Uint16(input[1:3])
	// request := binary.LittleEndian.Uint32(input[3:7])
	// realmsCount := binary.LittleEndian.Uint16(input[7:9])
	o := 9
	for {
		rlst := RealmListing{}
		rlst.Type = input[o]
		o++
		rlst.Color = input[o]
		o++
		rlst.Locked = input[o] == 1
		o++
		name := new(bytes.Buffer)
		for {
			if input[o] == 0 {
				break
			}
			name.WriteByte(input[o])
			o++
		}
		o++
		rlst.Name = name.String()
		ip := new(bytes.Buffer)
		for {
			if input[o] == 0 {
				break
			}
			ip.WriteByte(input[o])
			o++
		}
		o++
		rlst.Address = ip.String()
		pop := input[o : o+4]
		u := binary.LittleEndian.Uint32(pop)
		rlst.Population = math.Float32frombits(u)
		o += 4
		rlst.Characters = input[o]
		o++
		rlst.Timezone = input[o]
		o++
		rlst.ID = input[o]
		o++
		rls.Realms = append(rls.Realms, rlst)
		if o > int(size)+9 {
			break
		}
	}

	return rls, nil
}
