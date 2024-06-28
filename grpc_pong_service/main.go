package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"grpc_pong/proto/generated"
	"log"
	"net"
	"time"
)

func main() {

	config := Load()
	id := uuid.New().String()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.HOST, config.PORT))
	if err != nil {
		log.Fatalf("failed to listen: %s:%d", config.HOST, config.PORT)
	} else {
		log.Printf("listening on %s", lis.Addr().String())
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USERNAME, config.DB_PASSWORD, config.DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Failed to connect to database. Error: %w", err)
	}

	err = db.AutoMigrate(&model{})
	if err != nil {
		log.Panicln("Failed to auto migrate. Error: %w", err)
	}

	service := ServiceImpl{
		id: id,
		db: db,
	}

	// Forces connection timeout to allow client to reconnect with new connection
	// Does not give a perfect balance
	so := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     1 * time.Second,
			MaxConnectionAge:      1 * time.Second,
			MaxConnectionAgeGrace: 1 * time.Second,
			Time:                  1 * time.Second,
			Timeout:               1 * time.Second,
		}),
	}
	grpcServer := grpc.NewServer(so...)

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	generated.RegisterPongServer(grpcServer, &service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}

type ServiceImpl struct {
	generated.UnimplementedPongServer
	id string
	db *gorm.DB
}

func (s *ServiceImpl) Pong(ctx context.Context, req *generated.Request) (*generated.Response, error) {
	m := model{
		UUID: s.id,
		Type: "gRPC",
	}
	tx := s.db.Table("models").Clauses(clause.Locking{
		Strength: clause.LockingStrengthUpdate,
	}).Where("uuid = ?", s.id)
	if tx.Error != nil {
		log.Println("Error with transaction clause: ", tx.Error.Error())
		return nil, status.Error(codes.Aborted, tx.Error.Error())
	}
	tx.FirstOrCreate(&m)
	if tx.Error != nil {
		log.Println("Error with transaction FirstOrCreate: ", tx.Error.Error())
		return nil, status.Error(codes.Aborted, tx.Error.Error())
	}
	tx.Update("counter", m.Counter+1)
	if tx.Error != nil {
		log.Println("Error with transaction Update: ", tx.Error.Error())
		return nil, status.Error(codes.Aborted, tx.Error.Error())
	}
	log.Printf("UUID: %s, Counter: %d", m.UUID, m.Counter+1)
	return &generated.Response{}, nil
}

type model struct {
	gorm.Model
	UUID    string
	Type    string
	Counter int
}
