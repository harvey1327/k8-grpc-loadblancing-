package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var globalCounter *int32 = new(int32)

func main() {
	config := Load()
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        1,
			MaxIdleConnsPerHost: 1,
			MaxConnsPerHost:     1,
		},
		Timeout: 1 * time.Second,
	}
	requestHandler := requestHandler{
		config: config,
		cache:  NewCache(),
		client: &client,
	}

	http.HandleFunc("/ping", requestHandler.handle)
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
	config *Config
	cache  *Cache
	client *http.Client
}

func (h *requestHandler) handle(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(fmt.Sprintf("http://%s/pong", fmt.Sprintf("%s:%d", h.config.PONG_HOST, h.config.PONG_PORT)))

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		resp := response{}
		err = json.Unmarshal(b, &resp)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		h.cache.Increment(resp.ID)
		output := h.cache.Print()
		_, err = w.Write([]byte(output))
		w.WriteHeader(200)
		log.Println(output)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

	} else {
		w.WriteHeader(resp.StatusCode)
	}
}

type response struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}
