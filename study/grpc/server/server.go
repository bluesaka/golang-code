package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go-code/study/grpc/internal/server"
	"go-code/study/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// todo TLS
func main() {
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(),
	}

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterCarServiceServer(grpcServer, server.NewCarServer())
	proto.RegisterPhoneStreamServer(grpcServer, server.NewPhoneStreamServer())

	lis, err := net.Listen("tcp", ":9801")
	if err != nil {
		panic(err)
	}

	// use reflection.Register for grpcurl & grpcuri
	reflection.Register(grpcServer)

	log.Println("rpc server starting at :9801")
	grpcServer.Serve(lis)
}
