package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/ogier/pflag"
	"github.com/superp00t/gophercraft/blizzardry/mpq"
)

var (
	wowPath = pflag.StringP("wowpath", "w", "C:\\Program Files (x86)\\World of Warcraft\\", "/path/to/WOW/")
	addr    = pflag.StringP("addr", "a", "localhost:8090", "pipeline address")

	mpqPatterns = []string{
		"common.MPQ",
		"common-2.MPQ",
		"expansion.MPQ",
		"lichking.MPQ",
		"*/locale-*.MPQ",
		"*/speech-*.MPQ",
		"*/expansion-locale-*.MPQ",
		"*/lichking-locale-*.MPQ",
		"*/expansion-speech-*.MPQ",
		"*/lichking-speech-*.MPQ",
		"*/patch-*.MPQ",
		"patch.MPQ",
		"patch-*.MPQ",
	}

	mpqFiles = []string{}
	fileMap  = map[string]*fileEntry{}
)

type fileEntry struct {
	MpqPtr int
}

func FindMPQ(file string) string {
	fi := fileMap[file]
	if fi == nil {
		return ""
	}

	return mpqFiles[fi.MpqPtr]
}

func GetFileIndex(name string) int {
	for i, v := range mpqFiles {
		if name == v {
			return i
		}
	}
	return 0
}

func main() {
	pflag.Parse()

	for _, v := range mpqPatterns {
		m, _ := filepath.Glob(filepath.Join(*wowPath, v))
		mpqFiles = append(mpqFiles, m...)
	}

	for ptr, v := range mpqFiles {
		arch, err := mpq.Open(v)
		if err != nil {
			log.Fatal(err)
		}

		for _, ad := range arch.ListFiles() {
			fileMap[ad] = &fileEntry{ptr}
		}

		fmt.Println(v, "opened")
	}

	r := mux.NewRouter()
	r.PathPrefix("/test/").Handler(http.StripPrefix("/test", http.FileServer(http.Dir("./test/"))))
	p := r.PathPrefix("/pipeline/").Subrouter()

	// p.HandleFunc("/{resource}/list", ListMpq)
	p.StrictSlash(false)
	r.HandleFunc("/list_mpq", ListMpqFiles)
	r.HandleFunc("/list_files", ListFiles)

	p.HandleFunc("/{resource}", Resource)
	p.HandleFunc("/find/{query}", Find)

	log.Fatal(http.ListenAndServe(*addr, r))
}
