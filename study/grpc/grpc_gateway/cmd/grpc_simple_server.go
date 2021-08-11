package main

import (
	"context"
	"go-code/study/grpc/grpc_gateway/proto/simple"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type simpleServer struct{}

func (s *simpleServer) SayHello(ctx context.Context, in *simple.HelloRequest) (*simple.HelloReply, error) {
	return &simple.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	simple.RegisterGreeterServer(s, &simpleServer{})
	reflection.Register(s)
	log.Println("gRPC server listening on 0.0.0.0:8080")
	s.Serve(lis)
}
