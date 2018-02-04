package mpq

/* package mpq supports extraction of compressed files from MPQ (Mo'PaQ) archive files.

   Based on http://www.zezula.net/en/mpq/mpqformat.html
   Cryptographic functions taken from https://github.com/aphistic/go.Zamara

   TODO: - support table encryption
		 - implement test files for MPQ versions 1-4
		 - be memory efficient
*/

import (
	"fmt"
	"io"
	"math"
	"os"
	"sync"
)

const (
	MPQ_HEADER_DATA uint32 = 0x1A51504D
	MPQ_USER_DATA   uint32 = 0x1B51504D

	SectorSize = 512

	MD5_ListSize = 6

	MD5_BlockTable int = iota
	MD5_HashTable
	MD5_HiBlockTable
	MD5_BETTable
	MD5_HETTable
	MD5_MPQHeader
)

type MPQ struct {
	Header     *Header
	UserData   *UserData
	File       io.ReadSeeker
	HashTable  []*HashEntry
	BlockTable []*BlockEntry

	// Prevent access of multiple files at the same time
	L *sync.Mutex
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func (m *MPQ) BlockSize() int {
	return SectorSize * pow(2, int(m.Header.BlockSize))
}

func (m *MPQ) Version() int {
	return int(m.Header.FormatVersion) + 1
}

func Open(filename string) (*MPQ, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return Decode(f)
}

func Decode(i io.ReadSeeker) (*MPQ, error) {
	i.Seek(0, 0)

	h := new(Header)
	t := lu32(i)

	m := new(MPQ)
	m.Header = h
	m.File = i
	m.L = new(sync.Mutex)
	switch t {
	case MPQ_HEADER_DATA:
		if e := m.ReadHeaderData(); e != nil {
			return nil, e
		}
	case MPQ_USER_DATA:
		m.ReadUserData()
		m.File.Seek(m.Header.ArchiveOffset, 0)
		t = lu32(m.File)
		if t != MPQ_HEADER_DATA {
			return nil, fmt.Errorf("Could not find MPQ header")
		}
		if e := m.ReadHeaderData(); e != nil {
			return nil, e
		}
	default:
		return nil, fmt.Errorf("Invalid MPQ header")
	}

	m.ReadHashTable()
	m.ReadBlockTable()

	return m, nil
}
