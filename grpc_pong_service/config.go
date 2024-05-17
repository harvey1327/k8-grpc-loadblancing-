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
	HOST        string
	PORT        int
	DB_HOST     string
	DB_PORT     int
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
}

func Load() *Config {
	once.Do(func() {
		instance = &Config{
			HOST:        getEnv("HOST"),
			PORT:        getEnvAsInt("PORT"),
			DB_HOST:     getEnv("DB_HOST"),
			DB_PORT:     getEnvAsInt("DB_PORT"),
			DB_USERNAME: getEnv("DB_USERNAME"),
			DB_PASSWORD: getEnv("DB_PASSWORD"),
			DB_NAME:     getEnv("DB_NAME"),
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
