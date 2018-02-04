package packet

import (
	"bytes"
	"compress/zlib"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"io"
	"math"
	"time"

	"github.com/superp00t/gophercraft/gcore/glogger"
	"github.com/superp00t/gophercraft/guid"
)

func GetMSTime() uint32 {
	return uint32(time.Now().UnixNano() / int64(time.Millisecond))
}

type EtcBuffer struct {
	b    *bytes.Buffer
	rpos int
}

func (e *EtcBuffer) WriteBytes(data []byte) {
	e.b.Write(data)
}

func (e *EtcBuffer) WriteByte(data uint8) {
	e.b.WriteByte(data)
}

func (e *EtcBuffer) WriteCString(data string) {
	e.b.Write(append([]byte(data), 0)) // NULL terminator
}

func (e *EtcBuffer) WriteUint16(v uint16) {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, v)
	e.WriteBytes(buf)
}

func (e *EtcBuffer) WriteUint32(v uint32) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, v)
	e.WriteBytes(buf)
}

func (e *EtcBuffer) WriteUint64(v uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, v)
	e.WriteBytes(buf)
}

func (e *EtcBuffer) WriteFloat32(v float32) {
	f32 := math.Float32bits(v)
	e.WriteUint32(f32)
}

func (e *EtcBuffer) GetRPos() int {
	return e.rpos
}
func (e *EtcBuffer) ReadBytes(leng int) []byte {
	buf := make([]byte, leng)
	l := e.rpos + leng
	if l > e.b.Len() {
		// remove in production, will cause DoS vuln
		panic("FATAL OVERREAD!")
	}

	if e.b.Len() < e.rpos+leng {
		buf = e.b.Bytes()[e.rpos:]
	} else {
		buf = e.b.Bytes()[e.rpos:l]
	}
	e.rpos += leng
	return buf
}

func (e *EtcBuffer) Available() int {
	return len(e.b.Bytes()[e.rpos:])
}

func (e *EtcBuffer) ReadRemainingBytes() []byte {
	bff := e.b.Bytes()[e.rpos:]
	e.rpos += len(bff)
	return bff
}

func (e *EtcBuffer) ReadBigUint16() uint16 {
	return binary.BigEndian.Uint16(e.ReadBytes(2))
}

func (e *EtcBuffer) ReadUint16() uint16 {
	return binary.LittleEndian.Uint16(e.ReadBytes(2))
}

func (e *EtcBuffer) ReadUint32() uint32 {
	return binary.LittleEndian.Uint32(e.ReadBytes(4))
}

func (e *EtcBuffer) ReadUint64() uint64 {
	return binary.LittleEndian.Uint64(e.ReadBytes(8))
}

func (e *EtcBuffer) ReadByte() uint8 {
	return e.ReadBytes(1)[0]
}

func (e *EtcBuffer) ReadFloat32() float32 {
	v := e.ReadUint32()
	return math.Float32frombits(v)
}

func (e *EtcBuffer) ReadCString() string {
	buf := new(bytes.Buffer)
	for {
		if e.b.Bytes()[e.rpos] == 0 {
			break
		}

		buf.WriteByte(e.b.Bytes()[e.rpos])
		e.rpos++
	}

	e.rpos++
	return buf.String()
}

func (e *EtcBuffer) WritePackedGUID(g guid.GUID) {
	b := g.EncodePacked()
	e.WriteBytes(b)
}

func (e *EtcBuffer) ReadPackedGUID() guid.GUID {
	mask := e.ReadByte()
	if mask == 0 {
		return 0
	}

	glogger.Warnln("Byte mask", mask)
	var value guid.GUID
	var i uint32 = 0
	var flag uint8 = 1
	for {
		glogger.Println("Reading packed GUID", i, 8)
		if i == 8 {
			break
		}
		e := e.ReadByte()
		if e == 0 {
			break
		}
		if mask&(flag<<i) != 0 {
			value |= (guid.GUID(e) << (i * 8))
		}
		i++
	}
	// fmt.Printf("0x%X, 0x%X\n", res.High().U64(), res.U64())
	// glogger.Warnln("What?", res.High(), res)
	return value
}

type PackedXYZ struct {
	X, Y, Z, O float32
}

func (e *EtcBuffer) ReadPackedXYZ() *PackedXYZ {
	packed := e.ReadUint32()
	x := float32(((packed & 0x7FF) << 21 >> 21)) * 0.25
	z := float32((((packed>>11)&0x7FF)<<21)>>21) * 0.25
	y := float32((packed>>22<<22)>>22) * 0.25

	return &PackedXYZ{
		x, y, z, 0,
	}
}

func (e *EtcBuffer) Encode() []byte {
	return e.b.Bytes()
}

func (e *EtcBuffer) Len() int {
	return e.b.Len()
}

func (e *EtcBuffer) RPos(i int) {
	e.rpos = i
}

func NewEtcBuffer(b []byte) *EtcBuffer {
	eb := &EtcBuffer{}
	eb.b = new(bytes.Buffer)
	eb.b.Write(b)
	return eb
}

func reverseBuffer(input []byte) []byte {
	buf := make([]byte, len(input))
	inc := 0
	for x := len(input) - 1; x > -1; x-- {
		buf[inc] = input[x]
		inc++
	}
	return buf
}

func ReverseBuffer(input []byte) []byte {
	return reverseBuffer(input)
}

func packetString(input string) []byte {
	data := []byte(input)
	data = bytes.Replace(data, []byte("."), []byte{0}, -1)
	return data
}

func randomBuffer(l int) []byte {
	buf := make([]byte, l)
	rand.Read(buf)
	return buf
}

func PutU32(u uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, u)
	return buf
}

func PutF32(u float32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, math.Float32bits(u))
	return buf
}

func Hash(input ...[]byte) []byte {
	bt := sha1.Sum(bytes.Join(input, nil))
	return bt[:]
}

func BuildChecksum(data []byte) uint32 {
	h := Hash(data)
	var c uint32
	for i := 0; i < 5; i++ {
		o := i * 4
		nt := binary.LittleEndian.Uint32(h[o : o+4])
		c = c ^ nt
	}
	return c
}

func uncompress(input []byte) []byte {
	buf := new(bytes.Buffer)
	buf.Write(input)
	obuf := new(bytes.Buffer)
	rdr, err := zlib.NewReader(buf)
	if err != nil {
		panic(err)
		return nil
	}
	io.Copy(obuf, rdr)
	return obuf.Bytes()
}
