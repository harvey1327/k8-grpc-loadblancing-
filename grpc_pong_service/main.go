package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"grpc_pong/proto/generated"
	"log"
	"net"
)

func main() {

	conf := Load()
	id := "grpc-" + uuid.New().String()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.HOST, conf.PORT))
	if err != nil {
		log.Fatalf("failed to listen: %s:%d", conf.HOST, conf.PORT)
	} else {
		log.Printf("listening on %s", lis.Addr().String())
	}

	service := ServiceImpl{
		id: id,
	}

	grpcServer := grpc.NewServer()
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

func (s *ServiceImpl) Health(ctx context.Context, req *generated.Request) (*generated.Response, error) {
	return &generated.Response{}, nil
}

type model struct {
	gorm.Model
	UUID    string
	Counter int
}
