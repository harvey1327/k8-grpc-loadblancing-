package main

import (
	"log"
	"net/http"
)

func main() {
	times := 10
	for i := 0; i < times; i++ {
		request()
	}
}

func request() {
	resp, err := http.Get("http://127.0.0.1:8081/ping")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", resp)
}
