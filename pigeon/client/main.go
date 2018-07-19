package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/mappymappy/pigeon/pb"
	"github.com/mappymappy/pigeon/pigeon/pigeon_const"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s%s", os.Getenv(pigeon_const.ServerDSNEnvName), os.Getenv(pigeon_const.ServerPortEnvName)), grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewChatServiceClient(conn)
	ctx := context.Background()
	//Entering PersonName & signIn
	fmt.Print("Enter YourName:")
	stdin := bufio.NewScanner(os.Stdin)
	stdout := os.Stdout
	stdin.Scan()
	resp, err := client.SignIn(ctx, &pb.SignInRequest{Person: stdin.Text()})
	defer client.SignOut(ctx, &pb.SignOutRequest{})
	if err != nil {
		stdout.Write([]byte(fmt.Sprintf("signIn error::%#v \n", err)))
		return
	}
	md := metadata.New(map[string]string{pigeon_const.AuthHeaderName: resp.Token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	stream, err := client.Chattering(ctx)
	go chat(ctx, stream)
	// Observe stdIn & SendMessage
	for stdin.Scan() {
		text := stdin.Text()
		message := &pb.ChatRequest{Body: text}
		if err := stream.Send(message); err != nil {
			log.Fatalf("failed to send: %v", err)
			return
		}
	}
}
func chat(ctx context.Context, stream pb.ChatService_ChatteringClient) {
	ctx, cancel := context.WithCancel(ctx)
	//Observe ByDirectricalStreaming.
	defer stream.CloseSend()
	defer cancel()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Printf("Chat Stream error::%#v \n", err)
			return
		}
		if in != nil {
			log.Printf("%s:%s \n", in.Person, in.Body)
		}
	}
}
