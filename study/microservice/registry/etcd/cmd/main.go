package main

import (
	"go-code/study/microservice/registry/etcd/registry"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	endpoints := []string{"127.0.0.1:2379"}
	reg, err := registry.NewService(registry.ServiceInfo{
		Name: "rpc.mail/1",
		IP:   "127.0.0.1:8999",
	}, endpoints)

	if err != nil {
		panic(err)
	}

	reg2, err := registry.NewService(registry.ServiceInfo{
		Name: "rpc.test/2",
		IP:   "127.0.0.1:7999",
	}, endpoints)

	if err != nil {
		panic(err)
	}

	go reg.Start()
	go reg2.Start()


	lis, err := net.Listen("tcp", ":8999")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	log.Println("grpc server start")

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
