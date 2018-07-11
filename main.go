package main

import (
	"context"
	"marnie_playground/pigeon/pigeon/pigeon"
)

func main() {
	srv := pigeon.NewServer()
	ctx := context.TODO()
	//TODO:client mode & cli args..
	srv.Start(ctx)
}
