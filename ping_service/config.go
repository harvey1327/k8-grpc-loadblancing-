package main

import (
	"log"
	"os"
	"strconv"
	"sync"
)

var once sync.Once
var instance *Config

type Config struct {
	HOST           string
	PORT           int
	PONG_HOST      string
	PONG_PORT      int
	GRPC_PONG_HOST string
	GRPC_PONG_PORT int
}

func Load() *Config {
	once.Do(func() {
		instance = &Config{
			HOST:           getEnv("HOST"),
			PORT:           getEnvAsInt("PORT"),
			PONG_HOST:      getEnv("PONG_HOST"),
			PONG_PORT:      getEnvAsInt("PONG_PORT"),
			GRPC_PONG_HOST: getEnv("GRPC_PONG_HOST"),
			GRPC_PONG_PORT: getEnvAsInt("GRPC_PONG_PORT"),
		}
	})
	return instance
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("No value for %s", key)
	return ""
}

func getEnvAsInt(key string) int {
	valueS := getEnv(key)
	if valueI, err := strconv.Atoi(valueS); err == nil {
		return valueI
	}
	log.Fatalf("%s is not an int", key)
	return 0
}
