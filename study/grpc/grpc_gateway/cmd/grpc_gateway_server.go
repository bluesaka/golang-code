package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go-code/study/grpc/grpc_gateway/proto/gateway"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type echoServer struct{}

func (e echoServer) SayHello(ctx context.Context, request *gateway.EchoRequest) (*gateway.EchoReply, error) {
	return &gateway.EchoReply{Message: request.Name + " world"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	gateway.RegisterEchoServer(s, &echoServer{})

	log.Println("gRPC listening on 0.0.0.0:8080")
	go func() {
		s.Serve(lis)
	}()

	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	mux := runtime.NewServeMux()
	err = gateway.RegisterEchoHandler(context.Background(), mux, conn)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr: ":8081",
		Handler: mux,
	}
	log.Println("gRPC-Gateway listening on 0.0.0.0:8081")
	server.ListenAndServe()
}