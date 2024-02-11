package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	auditsystem "github.com/zeroalphat/image-audit/gen/proto/auditsystem/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const address = "localhost:8080"

type auditServer struct {
	auditsystem.UnimplementedImageAuditServiceServer
}

func (s *auditServer) PutImage(ctx context.Context, req *auditsystem.PutImageRequest) (*auditsystem.PutImageResponse, error) {
	return &auditsystem.PutImageResponse{
		Judgement: true,
	}, nil
}

func NewAuditServer() *auditServer {
	return &auditServer{}
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	auditsystem.RegisterImageAuditServiceServer(s, NewAuditServer())
	reflection.Register(s)
	go func() {
		log.Printf("start gRPC server: %v", address)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
