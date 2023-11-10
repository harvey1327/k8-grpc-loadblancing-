package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", handleRequest)
	log.Println("Starting up")
	var err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Panicln("Server failed starting. Error: %w", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving Request")
	w.Write([]byte("pong"))
	w.WriteHeader(200)
}
