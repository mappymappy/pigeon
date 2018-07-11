package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "marnie_playground/pigeon/pigeon/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	waitc := make(chan struct{})
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewChatServiceClient(conn)
	ctx := context.TODO()
	stream, err := client.Chattering(ctx)
	defer stream.CloseSend()
	//Entering PersonName
	fmt.Print("Enter YourName:")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	userName := stdin.Text()
	//Observe Stream
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				fmt.Printf("error::%#v \n", err)
			}
			fmt.Printf("%s:%s \n", in.Person, in.Body)
		}
	}()
	// Observe stdIn & SendMessage
	for stdin.Scan() {
		text := stdin.Text()
		message := &pb.ChatRequest{Body: text, Person: userName}
		if err := stream.Send(message); err != nil {
			log.Fatalf("failed to send: %v", err)
		}
	}
	<-waitc
}
