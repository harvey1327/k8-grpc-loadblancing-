package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"ping/proto/generated"
	"time"
)

func main() {
	config := Load()
	client := http.Client{
		Transport: &http.Transport{
			IdleConnTimeout:       1 * time.Second,
			TLSHandshakeTimeout:   1 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   1,
			MaxIdleConns:          1,
			DisableKeepAlives:     true,
		},
		Timeout: 1 * time.Second,
	}

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.PONG_HOST, config.PONG_PORT), options...)
	if err != nil {
		log.Fatal(err)
	}
	grpcClient := generated.NewPongClient(conn)

	requestHandler := &requestHandler{
		config: config,
		client: &client,
	}

	http.HandleFunc("/start", requestHandler.start)
	http.HandleFunc("/stop", requestHandler.stop)
	http.HandleFunc("/health", handleHealth)
	address := fmt.Sprintf("%s:%d", config.HOST, config.PORT)
	log.Printf("Starting up on: '%s'\n", address)
	err = http.ListenAndServe(address, nil)

	if err != nil {
		log.Panicln("Server failed starting. Error: %w", err)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

type requestHandler struct {
	config *Config
	client *http.Client
	cancel context.CancelFunc
}

func (h *requestHandler) start(w http.ResponseWriter, req *http.Request) {
	if h.cancel == nil {
		ctx, cancel := context.WithCancel(context.Background())
		h.cancel = cancel

		go h.process(ctx)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
}

func (h *requestHandler) process(ctx context.Context) {
	t := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case <-t.C:
			resp, err := h.client.Get(fmt.Sprintf("http://%s/pong", fmt.Sprintf("%s:%d", h.config.PONG_HOST, h.config.PONG_PORT)))
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println(resp.StatusCode)
		case <-ctx.Done():
			return
		}
	}
}

func (h *requestHandler) stop(w http.ResponseWriter, req *http.Request) {
	if h.cancel != nil {
		h.cancel()
		w.WriteHeader(200)
		h.cancel = nil
	} else {
		w.WriteHeader(500)
	}
}
