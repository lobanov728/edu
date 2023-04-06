package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request", r.URL.Path)
		time.Sleep(time.Second * 5)
		w.Write([]byte("response"))
	})
	srv := http.Server{
		Addr:    ":8081",
		Handler: nil,
	}
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = srv.Serve(ln)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err != nil {
		panic("Unable to start HTTP server")
	}
}
