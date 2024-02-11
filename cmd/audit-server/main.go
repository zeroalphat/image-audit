package main

import (
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"

	auditsystem "github.com/zeroalphat/image-audit/gen/proto/auditsystem/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const address = "localhost:8080"

type auditServer struct {
	logger *slog.Logger
	auditsystem.UnimplementedImageAuditServiceServer
}

func (s *auditServer) VerifyImage(ctx context.Context, req *auditsystem.VerifyImageRequest) (*auditsystem.VerifyImageResponse, error) {
	s.logger.Info("Receive request", "name", req.Name)

	return &auditsystem.VerifyImageResponse{
		Judgement: true,
	}, nil
}

func NewAuditServer(logger *slog.Logger) *auditServer {
	return &auditServer{
		logger: logger,
	}
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	s := grpc.NewServer()
	auditsystem.RegisterImageAuditServiceServer(s, NewAuditServer(logger))
	reflection.Register(s)
	go func() {
		logger.Info("start gRPC server", "address", address)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("stopping gRPC server")
	s.GracefulStop()
}
