package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: 2 * time.Second}
	times := 1000
	for i := 0; i < times; i++ {
		time.Sleep(10 * time.Millisecond)
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
