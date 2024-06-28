package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
)

var globalCounter *int32 = new(int32)

func main() {

	id := uuid.New().String()
	config := Load()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USERNAME, config.DB_PASSWORD, config.DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Failed to connect to database. Error: %w", err)
	}

	err = db.AutoMigrate(&model{})
	if err != nil {
		log.Panicln("Failed to auto migrate. Error: %w", err)
	}

	handler := requestHandler{id: id, db: db}

	http.HandleFunc("/pong", handler.handle)
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
	id string
	db *gorm.DB
}

func (h *requestHandler) handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	m := model{
		UUID: h.id,
		Type: "HTTP",
	}
	tx := h.db.Table("models").Clauses(clause.Locking{
		Strength: clause.LockingStrengthUpdate,
	}).Where("uuid = ?", h.id)
	if tx.Error != nil {
		w.WriteHeader(500)
		log.Println("Error with transaction clause: ", tx.Error.Error())
		return
	}
	tx.FirstOrCreate(&m)
	if tx.Error != nil {
		w.WriteHeader(500)
		log.Println("Error with transaction FirstOrCreate: ", tx.Error.Error())
		return
	}
	tx.Update("counter", m.Counter+1)
	if tx.Error != nil {
		w.WriteHeader(500)
		log.Println("Error with transaction Update: ", tx.Error.Error())
		return
	}
	log.Printf("UUID: %s, Counter: %d", m.UUID, m.Counter+1)
}

type model struct {
	gorm.Model
	UUID    string
	Type    string
	Counter int
}
