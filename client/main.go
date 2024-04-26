package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	client := http.Client{}
	times := 1000
	for i := 0; i < times; i++ {
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
