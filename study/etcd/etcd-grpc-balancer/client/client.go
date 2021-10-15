package main

import (
	"context"
	"go-code/study/etcd/etcd-grpc-balancer/etcdv3"
	"go-code/study/etcd/etcd-grpc-balancer/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"strconv"
)

func main() {
	builder := etcdv3.NewServiceDiscovery([]string{"localhost:2379"})
	resolver.Register(builder)

	conn, err := grpc.Dial(
		builder.Scheme()+":///"+"grpclb_test2",
		grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "weight"}`),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	grpcClient := proto.NewHelloServiceClient(conn)
	ctx := context.Background()
	for i := 1; i <= 20; i++ {
		resp, err := grpcClient.Say(ctx, &proto.HelloRequest{
			Name: "grpc" + strconv.Itoa(i),
		})
		if err != nil {
			panic(err)
		}
		log.Println(resp)
	}
}