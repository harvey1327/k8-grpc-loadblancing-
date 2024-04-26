package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

var globalCounter *int32 = new(int32)

func main() {

	id := uuid.New().String()
	config := Load()
	handler := requestHandler{id: id, config: config}

	http.HandleFunc("/pong", handler.handle)
	http.HandleFunc("/health", handleHealth)

	address := fmt.Sprintf("%s:%d", config.HOST, config.PORT)
	log.Printf("Starting up on: '%s'\n", address)
	var err = http.ListenAndServe(address, nil)

	if err != nil {
		log.Panicln("Server failed starting. Error: %w", err)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

type requestHandler struct {
	id     string
	config *Config
}

func (h *requestHandler) handle(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(h.config.TIMEOUT) * time.Millisecond)
	resp := response{ID: h.id, Message: "pong"}
	b, err := json.Marshal(&resp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	_, err = w.Write(b)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	log.Println("success")
}

type response struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}
