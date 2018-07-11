package chat

import (
	"context"
	"io"
	"log"
	pb "marnie_playground/pigeon/pigeon/pb"
)

type Service struct {
	chatStream chan pb.ChatResponse
}

func NewService(chatStream chan pb.ChatResponse) *Service {
	s := &Service{chatStream}
	return s
}

func (s *Service) bootChatDeliverProcess(server pb.ChatService_ChatteringServer) {
	for {
		res := <-s.chatStream
		log.Printf("response:%#v", res)
		if err := server.Send(&res); err != nil {
			log.Printf("error:%#v", err)
		}
	}
}

func (s *Service) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	return &pb.SignInResponse{}, nil
}

func (s *Service) SignOut(ctx context.Context, in *pb.SignOutRequest) (*pb.SignOutResponse, error) {
	return &pb.SignOutResponse{}, nil
}

func (s *Service) Chattering(server pb.ChatService_ChatteringServer) error {

	go s.bootChatDeliverProcess(server)
	for {
		in, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		resp := pb.ChatResponse{Body: in.Body, Person: in.Person}
		s.chatStream <- resp
	}
}
