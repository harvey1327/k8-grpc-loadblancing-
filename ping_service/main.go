package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/pong", handleRequest)
	http.HandleFunc("/health", handleHealth)
	log.Println("Starting up")
	var err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Panicln("Server failed starting. Error: %w", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling /pong")

	resp, err := http.Get("ping/ping")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.StatusCode)

	w.WriteHeader(200)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
