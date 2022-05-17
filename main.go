package main

import (
	"Social-Media-Backend/internal/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("Server listening at localhost:8080\n")

	// initiating the DB
	const dbPath = "db.json"
	dbClient := database.NewClient(dbPath)
	err := dbClient.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	apiCfg := apiConfig{
		dbClient: dbClient,
	}

	// creating the handlers
	mux := http.NewServeMux()

	// handling requests at the following paths:
	mux.HandleFunc("/users", apiCfg.endpointUsersHandler)
	mux.HandleFunc("/users/", apiCfg.endpointUsersHandler)
	mux.HandleFunc("/posts", apiCfg.endpointPostsHandler)
	mux.HandleFunc("/posts/", apiCfg.endpointPostsHandler)

	// starting the server
	server := http.Server{
		Handler:      mux,
		Addr:         "localhost:8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	err = server.ListenAndServe()

	// server blocks forever until an error is encountered
	log.Fatal(err)
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

type apiConfig struct {
	dbClient database.Client
}

// Use SQLite or PostgresQL instead of a JSON file for your database layer
// Let users save other data with their posts, maybe an image URL?
// Add proper authentication to each request, you could even use my password validator
// Add more unit tests!
// Deploy your api on AWS, GCP, or Digital Ocean
// Dockerize it
// Add documentation using markdown files
// Write a client that interacts with the API, maybe a webpage or a mobile app
