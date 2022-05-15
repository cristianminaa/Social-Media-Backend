package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("Server listening at localhost:8080\n")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := http.Server{
		Handler:      mux,
		Addr:         "localhost:8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	err := server.ListenAndServe()
	log.Fatal(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
