/*
	package dbc implements a decoder for Blizzard's database format DBC
*/
package dbc

import (
	"encoding/binary"
	"fmt"
	"io"
)

const DBC_MAGIC uint32 = 0x43424457

type DBC struct {
	Records         uint32
	Fields          uint32
	Size            uint32
	StringBlockSize uint32

	io.ReadSeeker
}

type Record struct {
	ID        uint32
	Reference uint32
}

func Decode(i io.ReadSeeker) (*DBC, error) {
	d := new(DBC)
	magic := lu32(i)
	if magic != DBC_MAGIC {
		return nil, fmt.Errorf("Not a DBC file")
	}

	d.Records = lu32(i)
	d.Fields = lu32(i)
	d.Size = lu32(i)
	d.StringBlockSize = lu32(i)
	d.ReadSeeker = i

	return d, nil
}

func (d *DBC) Query(id uint32) (string, error) {
	d.Seek()
}

func lu16(i io.Reader) uint16 {
	buf := lb(i, 2)
	return binary.LittleEndian.Uint16(buf)
}

func lu32(i io.Reader) uint32 {
	buf := lb(i, 4)
	return binary.LittleEndian.Uint32(buf)
}

func lu64(i io.Reader) uint64 {
	buf := lb(i, 8)
	return binary.LittleEndian.Uint64(buf)
}

func lb(i io.Reader, l int) []byte {
	buf := make([]byte, l)
	_, err := io.ReadAtLeast(i, buf, l)
	if err != nil {
		panic(err)
	}
	return buf
}
