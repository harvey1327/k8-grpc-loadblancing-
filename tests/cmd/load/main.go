package main

import (
	"log"
	"net/http"
)

func main() {
	run := false

	if run {
		start()
	} else {
		stop()
	}
}

func stop() {
	client := http.Client{}
	resp, err := client.Get("http://127.0.0.1:8081/stop")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.StatusCode)
}

func start() {
	client := http.Client{}
	resp, err := client.Get("http://127.0.0.1:8081/start")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.StatusCode)
}
