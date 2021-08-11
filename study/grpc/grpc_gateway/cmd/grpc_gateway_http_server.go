package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go-code/study/grpc/grpc_gateway/proto/gateway"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gateway.RegisterEchoHandlerFromEndpoint(context.Background(), mux, "localhost:8088", opts)
	if err != nil {
		panic(err)
	}

	log.Println("gRPC http server listening on 0.0.0.0:8081")
	http.ListenAndServe(":8081", mux)
}