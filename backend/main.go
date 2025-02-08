package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	api "open-tutor/internal/services/api"
	db "open-tutor/internal/services/db"
)

func main() {
	port := flag.String("port", "8080", "test port")
	flag.Parse()

	openTutor := api.Init()

	r := http.NewServeMux()
	h := api.HandlerFromMux(openTutor, r)
	s := &http.Server{
		Handler: h,
		Addr:    net.JoinHostPort("127.0.0.1", *port),
	}

	log.Fatal(s.ListenAndServe())
}
