package main

import (
	"context"
	"fmt"

	"github.com/jajajajaja15/gogogrpc/gen/pb/gogogrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := run(); err != nil {
		panic(fmt.Errorf("run: %w", err))
	}
}

func run() error {
	serverAddr := "127.0.0.1:8080"

	grpcOpt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(serverAddr, grpcOpt)

	if err != nil {
		return fmt.Errorf("grpc dial %s: %w", serverAddr, err)
	}
	defer conn.Close()

	client := gogogrpc.NewGreeterClient(conn)

	ctx := context.Background()
	req := &gogogrpc.HelloRequest{
		Name: "Jackson",
	}
	resp, err := client.SayHello(ctx, req)
	if err != nil {
		return fmt.Errorf("client sayhello")
	}

	fmt.Println(resp.Message)

	return nil
}
