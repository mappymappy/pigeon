package pigeon

import (
	"context"
	"log"
	"net"

	pb "marnie_playground/pigeon/pigeon/pb"
	chat "marnie_playground/pigeon/pigeon/pigeon/application/service/chat"

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
	service := chat.NewService(s.chatStream)
	pb.RegisterChatServiceServer(gsrv, service)
	//TODO: move cli args
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		return err
	}
	go func() {
		gsrv.Serve(listenPort)
		cancel()
	}()
	<-ctx.Done()
	/*
		//shutdown...
		s.Broadcast <- chat.StreamResponse{
			Timestamp: ptypes.TimestampNow(),
			Event: &chat.StreamResponse_ServerShutdown{
				&chat.StreamResponse_Shutdown{}}}
		close(s.Broadcast)
		ServerLogf(time.Now(), "shutting down")
		gsrv.GracefulStop
	*/

	return nil
}
