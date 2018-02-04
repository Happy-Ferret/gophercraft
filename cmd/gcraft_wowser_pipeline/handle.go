package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

func ErrResp(rw http.ResponseWriter, r *http.Request, err error, status int) {
	fmt.Println("responding with error", err)
	Resp(rw, r, struct {
		Err string `json:"error"`
	}{err.Error()}, status)
}

func ListFiles(rw http.ResponseWriter, r *http.Request) {
	log.Println("Attempted to list mpq files")
	Resp(rw, r, fileMap, 200)
}

func ListMpqFiles(rw http.ResponseWriter, r *http.Request) {
	Resp(rw, r, mpqFiles, 200)
}

func Resource(rw http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, ".png") {
		ServeBLP(rw, r)
		return
	}

	res := mux.Vars(r)["resource"]
	fmt.Println(res)
	data, err := LoadData(res)
	if err != nil {
		ErrResp(rw, r, err, 500)
		return
	}

	rw.Write(data)
}

type Entry struct {
	Filename string `json:"filename"`
	Name     string `json:"name"`
	Size     int    `json:"size"`
	Link     string `json:"link"`
}

func Find(rw http.ResponseWriter, r *http.Request) {
	qr := mux.Vars(r)["query"]
	var out []Entry
	for k := range fileMap {
		if strings.Contains(k, qr) {

			y := strings.Split(k, "\\")
			proto := "http"
			if r.TLS != nil {
				proto = "https"
			}
			str := url.QueryEscape(k)
			if strings.HasSuffix(str, "blp") {
				str += ".png"
			}
			out = append(out, Entry{
				Filename: k,
				Name:     y[len(y)-1],
				Size:     500,
				Link:     proto + "://" + r.Host + "/pipeline/" + str,
			})
		}
	}
	Resp(rw, r, out, 200)
}

func Resp(rw http.ResponseWriter, r *http.Request, v interface{}, status int) {
	rw.WriteHeader(status)
	e := json.NewEncoder(rw)
	if r.URL.Query().Get("p") == "1" {
		e.SetIndent("", "  ")
	}
	e.Encode(v)
}
