package main

import (
	"log"
	"net"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/gorilla/mux"
	"github.com/ogier/pflag"
	"github.com/superp00t/gophercraft/gcore/glogger"
)

var (
	addr  = pflag.StringP("addr", "a", "localhost:9090", "address to listen on")
	raddr = pflag.StringP("raddr", "r", "ds", "address to connect to")
	id    = 0
)

func proxy(rw http.ResponseWriter, r *http.Request) {
	c, err := net.Dial("tcp", *raddr)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	srv := websocket.Server{Handler: websocket.Handler(func(ws *websocket.Conn) {
		id++
		me := id
		glogger.Println("Opening", me)
		fail := false
		go func() {
			for {
				if fail {
					break
				}
				dat := make([]byte, 4096)
				i, err := ws.Read(dat)
				if err != nil {
					glogger.Warnln(me, "Error reading from client", err)
					break
				}
				glogger.Println(me, "read", i, "from client")
				_, err = c.Write(dat[:i])
				if err != nil {
					glogger.Warnln(me, "Error relaying to server", err)
					break
				}
				glogger.Println("Successful C->S transfer!")
			}
		}()

		for {
			if fail {
				break
			}
			dat := make([]byte, 12000)
			r, err := c.Read(dat)
			if err != nil {
				glogger.Warnln(me, "error receiving from server", err)
				break
			}
			err = websocket.Message.Send(ws, dat[:r])
			if err != nil {
				glogger.Warnln(me, "Error relaying to client", err)
				break
			}
			glogger.Println("Successful S->C transfer!")
		}
		glogger.Println("Closing", me)
	})}

	srv.ServeHTTP(rw, r)
}

func main() {
	pflag.Parse()
	r := mux.NewRouter()

	r.HandleFunc("/", proxy)
	log.Fatal(http.ListenAndServe(*addr, r))
}
