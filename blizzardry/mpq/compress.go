package mpq

import (
	"bytes"
	"compress/bzip2"
	"compress/zlib"
	"fmt"
	"io"

	"github.com/superp00t/gophercraft/blizzardry/pkzip"
)

type CompressionType uint8

const (
	MPQ_COMPRESSION_HUFFMANN     CompressionType = 0x01 // Huffmann compression (used on WAVE files only)
	MPQ_COMPRESSION_ZLIB         CompressionType = 0x02 // ZLIB compression
	MPQ_COMPRESSION_PKWARE       CompressionType = 0x08 // PKWARE DCL compression
	MPQ_COMPRESSION_BZIP2        CompressionType = 0x10 // BZIP2 compression (added in Warcraft III)
	MPQ_COMPRESSION_SPARSE       CompressionType = 0x20 // Sparse compression (added in Starcraft 2)
	MPQ_COMPRESSION_ADPCM_MONO   CompressionType = 0x40 // IMA ADPCM compression (mono)
	MPQ_COMPRESSION_ADPCM_STEREO CompressionType = 0x80 // IMA ADPCM compression (stereo)
	MPQ_COMPRESSION_LZMA         CompressionType = 0x12 // LZMA compression. Added in Starcraft 2. This value is NOT a combination of flags.
	MPQ_COMPRESSION_NEXT_SAME    CompressionType = 0xFF // Same compression

)

func (c CompressionType) String() string {
	for _, v := range CompressionTable {
		if v.Key == c {
			return v.DC.Name
		}
	}
	return fmt.Sprintf("%d", c)
}

type Decompressor struct {
	Name string
	DC   func(in []byte) []byte
}

type DCTable struct {
	Key CompressionType
	DC  *Decompressor
}

var CompressionTable = []DCTable{
	{MPQ_COMPRESSION_HUFFMANN, &Decompressor{"Huffman trees", nil}},
	{MPQ_COMPRESSION_ZLIB, &Decompressor{"zlib", dcZlib}},
	{MPQ_COMPRESSION_PKWARE, &Decompressor{"PKWARE dcl", dcZip}},
	{MPQ_COMPRESSION_BZIP2, &Decompressor{"bzip2", dcBzip2}},
	{MPQ_COMPRESSION_ADPCM_MONO, &Decompressor{"wave (mono)", nil}},
	{MPQ_COMPRESSION_ADPCM_STEREO, &Decompressor{"wave (stereo)", nil}},
	{MPQ_COMPRESSION_SPARSE | MPQ_COMPRESSION_ZLIB, &Decompressor{"Sparse ZLIB", nil}},
	{MPQ_COMPRESSION_SPARSE | MPQ_COMPRESSION_BZIP2, &Decompressor{"Sparse Bzip!", nil}},
}

// func GetDecompressors(flags uint8) ([]*Decompressor, error) {
// 	if flags == COMPRESS_BZIP2 {
// 		log.Println("Showing bias")
// 		return []*Decompressor{
// 			{"bzip", dcBzip2},
// 		}, nil
// 	}

// 	var d []*Decompressor

// 	for _, tb := range CompressionTable {
// 		mask := tb.Key
// 		fn := tb.DC
// 		if (flags & uint8(mask)) != 0 {
// 			if fn.DC == nil {
// 				return nil, fmt.Errorf("Unsupported algorithm type: %s", mask)
// 			}

// 			d = append(d, fn)
// 		}
// 	}

// 	return d, nil
// }

func DecompressBlock(n []byte) ([]byte, error) {
	flags := n[0]
	fl := CompressionType(flags)
	for _, v := range CompressionTable {
		if v.Key == fl {
			if v.DC.DC == nil {
				return nil, fmt.Errorf("No handler for %s", v.Key)
			}

			return v.DC.DC(n[1:]), nil
		}
	}

	return nil, fmt.Errorf("Unsupported algorithm type: %s", fl)
}

func dcZip(in []byte) []byte {
	b, err := pkzip.Decompress(in)
	if err != nil {
		panic(err)
	}

	return b
}

func dcBzip2(in []byte) []byte {
	bf := bytes.NewBuffer(in)
	dr := bzip2.NewReader(bf)
	out := new(bytes.Buffer)
	_, err := io.Copy(out, dr)
	if err != nil {
		panic(err)
	}

	return out.Bytes()
}

func dcZlib(in []byte) []byte {
	bf := bytes.NewBuffer(in)
	dr, err := zlib.NewReader(bf)
	if err != nil {
		panic(err)
	}

	out := new(bytes.Buffer)
	_, err = io.Copy(out, dr)
	if err != nil {
		panic(err)
	}

	return out.Bytes()
}
