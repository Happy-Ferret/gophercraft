package mpq

import (
	"fmt"
	"os"
	"sync"
)

// Pool pre-loads multiple MPQ archive headers, allowing fast, concurrent access of archive files
type Pool struct {
	// Archive data
	fmap map[string]*archiveEntry

	names []string
}

type archiveEntry struct {
	nameIndex  uint16
	header     *Header
	hashTable  []*HashEntry
	blockTable []*BlockEntry
}

// OpenPool opens a Pool using a slice of MPQ file paths
func OpenPool(names []string) (*Pool, error) {
	p := &Pool{}
	p.fmap = make(map[string]*archiveEntry)
	p.names = names

	for i, v := range names {
		m, err := Open(v)
		if err != nil {
			return nil, err
		}

		ae := new(archiveEntry)
		ae.nameIndex = uint16(i)
		ae.header = m.Header
		ae.hashTable = m.HashTable
		ae.blockTable = m.BlockTable

		lf := m.ListFiles()

		for _, fv := range lf {
			p.fmap[fv] = ae // map filepath string to MPQ data pointer
		}
	}

	return p, nil
}

func (p *Pool) OpenFile(name string) (*File, error) {
	ae := p.fmap[name]
	if ae == nil {
		return nil, fmt.Errorf("File not found")
	}

	m := new(MPQ)
	var err error
	m.File, err = os.Open(p.names[ae.nameIndex])
	if err != nil {
		return nil, err
	}

	m.L = new(sync.Mutex)
	m.Header = ae.header
	m.BlockTable = ae.blockTable
	m.HashTable = ae.hashTable

	return m.OpenFile(name)
}
