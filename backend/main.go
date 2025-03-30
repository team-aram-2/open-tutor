package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	api "open-tutor/internal/services/api"
	middleware "open-tutor/middleware"
	zoom "open-tutor/zoom"

	"github.com/joho/godotenv"
)

type Middleware func(http.Handler) http.HandlerFunc

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	zoom.LoadAccessToken()

	port := flag.String("port", "8080", "OpenTutor API port")
	flag.Parse()

	openTutor := api.Init()

	r := http.NewServeMux()
	h := api.HandlerFromMux(openTutor, r)

	// Middleware chain in order of precedence //
	middlewareChain := []Middleware{
		middleware.EnableCORS,
		middleware.JSONContentType,
		middleware.Authenticate,
		middleware.ModerationCheck,
	}
	/////////////////////////////////////////////

	for i := 0; i < len(middlewareChain); i++ {
		h = middlewareChain[i](h)
	}

	s := &http.Server{
		Handler: h,
		Addr:    net.JoinHostPort("127.0.0.1", *port),
	}

	log.Fatal(s.ListenAndServe())
}
