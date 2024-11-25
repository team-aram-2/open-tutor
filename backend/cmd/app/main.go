package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "pong"}

	json.NewEncoder(w).Encode(response)
	fmt.Println(w, "pong")
}

func main() {
	http.HandleFunc("/ping", ping)

	fmt.Println("Starting server on 8080")

	http.ListenAndServe(":8080", nil)
}
