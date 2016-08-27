package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/uvekilledkenny/WebChip8/core"
)

var (
	c    = core.New()
	port int
)

func main() {
	flag.IntVar(&port, "port", 8000, "port")
	flag.Parse()

	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("127.0.0.1:%v", port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	r.HandleFunc("/ws", chipHandler)

	log.Println("Listening on " + srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
