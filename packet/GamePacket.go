package packet

import (
	"bytes"
	"encoding/binary"
	"log"
)

const (
	OPCODE_SIZE_OUTGOING = 6
	OPCODE_SIZE_INCOMING = 4
)

type SMSGAuthPacket struct {
	Size uint16
	Type WorldType
	Salt []byte
}

type GamePacket struct {
	Type WorldType
	Buf  *bytes.Buffer
}

func (gp *GamePacket) PutCString(c string) {
	bf := append([]byte(c), 0)
	gp.Buf.Write(bf)
}

func (gp *GamePacket) PutUint(i uint32) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, i)
	gp.Buf.Write(buf)
}

func (gp *GamePacket) Finish() []byte {
	head := new(bytes.Buffer)
	size := uint16(gp.Buf.Len() + 4)
	head.Write(u16(size))
	head.Write(opEncode(gp.Type))
	head.Write(gp.Buf.Bytes())
	log.Println(size, len(head.Bytes()))
	return head.Bytes()
}

func NewGamePacket(typ WorldType) *GamePacket {
	gp := &GamePacket{
		Type: typ,
		Buf:  new(bytes.Buffer),
	}

	return gp
}

func UnmarshalSMSGAuthPacket(input []byte) (*SMSGAuthPacket, error) {
	gp := &SMSGAuthPacket{}
	gp.Size = binary.BigEndian.Uint16(input[0:2])
	gp.Type = WorldType(binary.LittleEndian.Uint16(input[2:4]))
	unint := binary.LittleEndian.Uint32(input[4:8])
	log.Printf("0x%x\n", unint)
	gp.Salt = input[8:12]

	return gp, nil
}

func u16(i uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, i)
	return buf
}

func opEncode(w WorldType) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(w))
	return buf
}

type WorldPacket struct {
	Size uint16
	Type WorldType
}

func DecodeWorldPacket(input []byte) (*WorldPacket, error) {
	size := binary.BigEndian.Uint16(input[0:2])
	log.Println(len(input), size)
	wt := binary.LittleEndian.Uint16(input[2:4])
	wp := &WorldPacket{}
	wp.Size = size
	wp.Type = WorldType(wt)
	return wp, nil
}
