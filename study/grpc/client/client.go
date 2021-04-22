package main

import (
	"context"
	"go-code/study/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9801", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewCarServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	CarList(client, ctx)
	CarQuery(client, ctx)
	CarUpdate(client, ctx)
}

func CarList(client proto.CarServiceClient, ctx context.Context) {
	resp, err := client.List(ctx, &proto.CarReq{})
	if err != nil {
		log.Println("carList error:", err)
	}
	log.Printf("%+v\n\n", resp)
}

func CarQuery(client proto.CarServiceClient, ctx context.Context) {
	resp, err := client.Query(ctx, &proto.CarReq{Name: "benz"})
	if err != nil {
		log.Println("carQuery error:", err)
	}
	log.Printf("%+v\n\n", resp)
}

func CarUpdate(client proto.CarServiceClient, ctx context.Context) {
	resp, err := client.Update(ctx, &proto.CarReq{Name: "bmw"})
	if err != nil {
		log.Println("carUpdate error:", err)
	}
	log.Printf("%+v\n\n", resp)
}
