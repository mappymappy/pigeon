package chat

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "github.com/mappymappy/pigeon/pb"
	"github.com/mappymappy/pigeon/pigeon/pigeon_const"

	"google.golang.org/grpc/metadata"
)

type Service struct {
	postOffice        chan pb.ChatResponse
	posts             map[string]chan pb.ChatResponse
	loginSessions     map[string]string
	chatResponseMutex sync.RWMutex
	loginMutex        sync.RWMutex
	salt              string
}

func NewService(chatStream chan pb.ChatResponse, salt string) *Service {
	s := &Service{
		postOffice:    chatStream,
		posts:         make(map[string]chan pb.ChatResponse),
		loginSessions: map[string]string{},
		salt:          salt,
	}
	go s.bootPostOfficeProcess()
	return s
}

func (s *Service) RegisterPost(id string) {
	s.chatResponseMutex.Lock()
	s.posts[id] = make(chan pb.ChatResponse, 1000)
	s.chatResponseMutex.Unlock()
}

func (s *Service) bootPostmanProcess(id string, server pb.ChatService_ChatteringServer) {
	for {
		res := <-s.posts[id]
		log.Printf("send message user:%s body:%s", id, res.Body)
		if err := server.Send(&res); err != nil {
			log.Printf("error:%#v \n", err)
		}
	}
}

func (s *Service) bootPostOfficeProcess() {
	for {
		res := <-s.postOffice
		s.chatResponseMutex.RLock()
		for _, post := range s.posts {
			post <- res
		}
		s.chatResponseMutex.RUnlock()
	}
}

func (s *Service) generateToken(userName string) string {
	token := make([]byte, 64)
	rand.Seed(time.Now().UnixNano())
	rand.Read(token)
	seed := s.salt + string(token)
	hash := sha256.Sum256([]byte(seed))
	return fmt.Sprintf("%x", hash)
}

func (s *Service) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	token := s.generateToken(in.Person)
	s.loginMutex.Lock()
	s.loginSessions[token] = in.Person
	s.loginMutex.Unlock()
	log.Printf("signIn user:%s token:%s", in.Person, token)
	resp := pb.ChatResponse{Body: fmt.Sprintf("signIn user:%s", in.Person), Person: pigeon_const.SystemManagerName}
	s.postOffice <- resp

	return &pb.SignInResponse{Token: token}, nil
}

func (s *Service) SignOut(ctx context.Context, in *pb.SignOutRequest) (*pb.SignOutResponse, error) {
	token, err := s.getAuthTokenFromCtx(ctx)
	if err != nil {
		return &pb.SignOutResponse{}, err
	}
	userName, _ := s.getLoginSession(token)
	delete(s.loginSessions, token)
	log.Printf("signOut user:%s token:%s", userName, token)
	resp := pb.ChatResponse{Body: fmt.Sprintf("signOut user:%s", userName), Person: pigeon_const.SystemManagerName}
	s.postOffice <- resp

	return &pb.SignOutResponse{}, nil
}

func (s *Service) Chattering(server pb.ChatService_ChatteringServer) error {
	token, err := s.getAuthTokenFromCtx(server.Context())
	if err != nil {
		return err
	}
	userName, err := s.authenticateByToken(token)
	if err != nil {
		return err
	}
	s.RegisterPost(token)
	go s.bootPostmanProcess(token, server)
	for {
		in, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		resp := pb.ChatResponse{Body: in.Body, Person: userName}
		s.postOffice <- resp
	}
}

func (s *Service) getAuthTokenFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md[pigeon_const.AuthHeaderName]) == 0 {
		return "", errors.New("Not Found AuthToken.")
	}
	return md[pigeon_const.AuthHeaderName][0], nil
}

func (s *Service) authenticateByToken(token string) (string, error) {
	userName, exist := s.getLoginSession(token)
	if !exist {
		return "", errors.New("UnAuthorizedToken")
	}
	return userName, nil
}

func (s *Service) getLoginSession(key string) (string, bool) {
	s.loginMutex.Lock()
	userName, exist := s.loginSessions[key]
	s.loginMutex.Unlock()
	return userName, exist
}
