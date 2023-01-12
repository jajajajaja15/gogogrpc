package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jajajajaja15/gogogrpc/gen/pb/gogogrpc"

	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	listenOn := ":8080"
	listener, err := net.Listen("tcp", listenOn)

	if err != nil {
		return fmt.Errorf("listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	gogogrpc.RegisterGreeterServer(server, &GreeterServerHandler{})

	log.Println("Listening on", listenOn)

	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("server gRPC server: %w", err)
	}
	return nil
}

var _ gogogrpc.GreeterServer = &GreeterServerHandler{}

type GreeterServerHandler struct {
	gogogrpc.UnimplementedGreeterServer
}

func (h *GreeterServerHandler) SayHello(
	ctx context.Context, req *gogogrpc.HelloRequest,
) (*gogogrpc.HelloReply, error) {
	fmt.Println(req.Name, "say hello")
	return &gogogrpc.HelloReply{
		Message: fmt.Sprintln("hello", req.Name),
	}, nil
}
