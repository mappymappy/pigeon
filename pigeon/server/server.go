package server

import (
	"context"
	"log"
	"marnie_playground/pigeon/pigeon/pigeon/pigeon_const"
	"net"
	"os"

	pb "github.com/mappymappy/pigeon/pb"
	chat "github.com/mappymappy/pigeon/pigeon/server/application/service/chat"

	"google.golang.org/grpc"
)

type Server struct {
	chatStream chan pb.ChatResponse
}

func NewServer() *Server {
	return &Server{
		chatStream: make(chan pb.ChatResponse, 1000),
	}
}

func (s *Server) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	log.Print("start pigeon server...\n")
	gsrv := grpc.NewServer()
	service := chat.NewService(ctx, s.chatStream, os.Getenv(pigeon_const.ServerSecureSalt))
	pb.RegisterChatServiceServer(gsrv, service)
	listenPort, err := net.Listen("tcp", os.Getenv(pigeon_const.ServerPortEnvName))
	if err != nil {
		return err
	}
	go func() {
		gsrv.Serve(listenPort)
		cancel()
	}()
	<-ctx.Done()
	gsrv.GracefulStop()

	return nil
}
