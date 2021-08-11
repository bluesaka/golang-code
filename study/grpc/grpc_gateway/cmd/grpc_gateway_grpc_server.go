package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go-code/study/grpc/grpc_gateway/proto/gateway"
	"google.golang.org/grpc"
	"log"
	"net"
)

type echoGRPCServer struct{}

func (e echoGRPCServer) SayHello(ctx context.Context, request *gateway.EchoRequest) (*gateway.EchoReply, error) {
	return &gateway.EchoReply{Message: request.Name + " world"}, nil
}

func main() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gateway.RegisterEchoHandlerFromEndpoint(context.Background(), mux, "localhost:8088", opts)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	gateway.RegisterEchoServer(s, &echoGRPCServer{})

	log.Println("gRPC server listening on 0.0.0.0:8088")

	s.Serve(lis)
}
