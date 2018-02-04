package packet

import (
	"bytes"
	"encoding/binary"
)

type WorldPacket struct {
	Type WorldType
	*EtcBuffer
}

func NewWorldPacket(t WorldType) *WorldPacket {
	return &WorldPacket{t, NewEtcBuffer(nil)}
}

func (wp *WorldPacket) Finish() []byte {
	head := new(bytes.Buffer)

	// Write length
	l := make([]byte, 2)
	sz := uint16(wp.Len()) + 2
	binary.BigEndian.PutUint16(l, sz) // 42
	head.Write(l)

	t := make([]byte, 2)
	binary.LittleEndian.PutUint16(t, uint16(wp.Type))
	head.Write(t)

	head.Write(wp.Encode())
	return head.Bytes()
}

func (wp *WorldPacket) Write(buf []byte) {
	wp.WriteBytes(buf)
}
