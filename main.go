package main

import (
	"Social-Media-Backend/internal/database"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("Server listening at localhost:8080\n")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/err", errHandler)
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
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, errors.New("internal server error"))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "GET, POST, OPTIONS, PUT, DELETE")
	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			w.WriteHeader(500)
			response, _ := json.Marshal(errorBody{
				Error: "Error marshalling JSON",
			})
			w.Write(response)
			return
		}
		w.WriteHeader(code)
		w.Write(response)
	}
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err != nil {
		log.Println(err)
		respondWithJSON(w, code, errorBody{
			Error: err.Error(),
		})
	}
	log.Println("don't call respondWithError with no error")
}

type errorBody struct {
	Error string `json:"error"`
}
