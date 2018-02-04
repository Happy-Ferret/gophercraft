package main

import (
	"fmt"
	"image/png"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/superp00t/gophercraft/blizzardry/blp"
	"github.com/superp00t/gophercraft/blizzardry/mpq"
)

func LoadData(path string) ([]byte, error) {
	ff, err := LoadFile(path)
	if err != nil {
		return nil, err
	}
	data, err := ff.ReadBlock()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func LoadFile(path string) (*mpq.File, error) {
	mp := FindMPQ(path)
	if mp == "" {
		return nil, fmt.Errorf("no MPQ found")
	}

	m, err := mpq.Open(mp)
	if err != nil {
		return nil, err
	}

	ff, err := m.OpenFile(path)
	if err != nil {
		return nil, err
	}

	return ff, nil
}

func ServeBLP(rw http.ResponseWriter, r *http.Request) {
	f := mux.Vars(r)["resource"]
	f = strings.Replace(f, ".png", "", -1)
	fmt.Println("BLP image requested", f)
	data, err := LoadData(f)
	if err != nil {
		ErrResp(rw, r, err, 500)
		return
	}
	bp, err := blp.DecodeBLP(data)
	if err != nil {
		ErrResp(rw, r, err, 500)
		return
	}
	img := bp.Mipmap(0)
	e := png.Encode(rw, img)
	if e != nil {
		panic(e)
	}

	// Resp(rw, r, []string{fmt.Sprintf("%d", len(data))}, 200)
}
