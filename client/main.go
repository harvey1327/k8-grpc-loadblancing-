package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: 1 * time.Second}
	times := 10
	for i := 0; i < times; i++ {
		request(&client)
	}
}

func request(client *http.Client) {
	resp, err := client.Get("http://127.0.0.1:8081/ping")
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
}
