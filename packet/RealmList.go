package packet

import (
	"bytes"
	"encoding/binary"
	"log"
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

func UnmarshalRealmList_S(input []byte) (*RealmList_S, error) {
	rls := &RealmList_S{}
	rls.Cmd = AuthType(input[0])
	size := binary.LittleEndian.Uint16(input[1:3])
	request := binary.LittleEndian.Uint32(input[3:7])
	log.Println("Request: ", request)
	realmsCount := binary.LittleEndian.Uint16(input[7:9])
	log.Println("Realmscount", realmsCount)
	o := 9
	// for i := 0; i < int(lists); i++ {
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
		log.Println(rlst.Address)
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
		if o > int(size) {
			break
		}
	}

	return rls, nil
}
