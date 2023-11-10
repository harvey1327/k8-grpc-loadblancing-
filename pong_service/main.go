package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var globalCounter *int32 = new(int32)

func main() {

	http.HandleFunc("/pong", handleRequest)
	http.HandleFunc("/health", handleHealth)

	config := Load()
	address := fmt.Sprintf("%s:%d", config.HOST, config.PORT)
	log.Printf("Starting up on: '%s'\n", address)
	var err = http.ListenAndServe(address, nil)

	if err != nil {
		log.Panicln("Server failed starting. Error: %w", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving Request number: ", atomic.AddInt32(globalCounter, 1))
	_, err := w.Write([]byte("pong"))
	if err != nil {
		w.WriteHeader(500)
	}
	w.WriteHeader(200)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
