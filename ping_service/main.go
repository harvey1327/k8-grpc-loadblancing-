package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", handleRequest)
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
	config := Load()
	address := fmt.Sprintf("%s:%d", config.PONG_HOST, config.PONG_PORT)
	log.Printf("Calling '%s/pong'\n", address)

	resp, err := http.Get(fmt.Sprintf("http://%s/pong", address))
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode == 200 {
		w.WriteHeader(200)
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(500)
		}
		_, err = w.Write(b)
		if err != nil {
			w.WriteHeader(500)
		}
	} else {
		w.WriteHeader(500)
	}

}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
