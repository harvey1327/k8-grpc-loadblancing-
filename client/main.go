package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Transport: &http.Transport{
			IdleConnTimeout:       1 * time.Second,
			TLSHandshakeTimeout:   1 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: 1 * time.Second,
	}
	times := 1000
	for i := 0; i < times; i++ {
		time.Sleep(5 * time.Millisecond)
		request(&client)
	}
}

func request(client *http.Client) {
	resp, err := client.Get("http://127.0.0.1:8081/ping")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("StatusCode: %d, Response: %s\n", resp.StatusCode, string(b))
}
