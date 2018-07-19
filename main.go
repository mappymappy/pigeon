package main

import (
	"context"

	"github.com/mappymappy/pigeon/pigeon/server"
)

func main() {
	// start chat server
	srv := server.NewServer()
	ctx := context.Background()
	srv.Start(ctx)
}
