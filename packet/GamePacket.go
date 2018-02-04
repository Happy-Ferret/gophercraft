package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

const (
	OPCODE_SIZE_OUTGOING = 6
	OPCODE_SIZE_INCOMING = 4
)

type SMSGAuthPacket struct {
	Type  WorldType
	Size  uint16
	Salt  []byte
	Seed1 []byte
	Seed2 []byte
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
	gp.Seed1 = input[12:28]
	gp.Seed2 = input[28:44]

	return gp, nil
}

func (s *SMSGAuthPacket) Encode() []byte {
	smsg := NewWorldPacket(SMSG_AUTH_CHALLENGE)
	smsg.WriteUint32(0x01)
	smsg.Write(s.Salt)
	smsg.Write(s.Seed1)
	smsg.Write(s.Seed2)
	return smsg.Finish()
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

type CMSGAuthSession struct {
	Build           uint32
	LoginServerID   uint32
	Account         string // 0-terminated string
	LoginServerType uint32
	Seed            []byte
	RegionID        uint32
	BattlegroupID   uint32
	RealmID         uint32
	DosResponse     uint64
	Digest          []byte
	AddonData       []byte
}

func UnmarshalCMSGAuthSession(input []byte) (*CMSGAuthSession, error) {
	// opcode = input[0:4]
	// len    = input[4:6]
	if len(input) < 36 {
		return nil, fmt.Errorf("packet too small")
	}
	c := &CMSGAuthSession{}
	c.Build = binary.LittleEndian.Uint32(input[6:10])
	c.LoginServerID = binary.LittleEndian.Uint32(input[10:14])
	o := 14
	acc := new(bytes.Buffer)
	for {
		if o > len(input)-1 {
			return nil, fmt.Errorf("Invalid packet")
		}
		if input[o] == 0 {
			break
		}

		acc.WriteByte(input[o])
		o++
	}
	o++ // zero
	c.Account = acc.String()
	c.LoginServerType = binary.LittleEndian.Uint32(input[o : o+4])
	o += 4
	c.Seed = input[o : o+4]
	o += 4
	c.RegionID = binary.LittleEndian.Uint32(input[o : o+4])
	o += 4
	c.BattlegroupID = binary.LittleEndian.Uint32(input[o : o+4])
	o += 4
	c.RealmID = binary.LittleEndian.Uint32(input[o : o+4])
	o += 4
	c.DosResponse = binary.LittleEndian.Uint64(input[o : o+8])
	o += 8
	c.Digest = input[o : o+20]
	c.AddonData = input[o+20:]
	return c, nil
}

// type WorldPacket struct {
// 	Size uint16
// 	Type WorldType
// }

// func DecodeWorldPacket(input []byte) (*WorldPacket, error) {
// 	size := binary.BigEndian.Uint16(input[0:2])
// 	log.Println(len(input), size)
// 	wt := binary.LittleEndian.Uint16(input[2:4])
// 	wp := &WorldPacket{}
// 	wp.Size = size
// 	wp.Type = WorldType(wt)
// 	return wp, nil
// }

type SMSGAuthResponse struct {
	Cmd       uint8
	WaitQueue uint32
}

func UnmarshalSMSGAuthResponse(input []byte) (*SMSGAuthResponse, error) {
	p := NewEtcBuffer(input)
	p.RPos(4)
	s := &SMSGAuthResponse{}
	s.Cmd = p.ReadByte()

	if s.Cmd != AUTH_WAIT_QUEUE {
		return s, nil
	}

	s.WaitQueue = p.ReadUint32()
	p.ReadByte()
	return s, nil
}
