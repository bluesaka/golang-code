package main

import (
	"context"
	"go-code/study/etcd/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	"time"

	pb "go-code/study/etcd/proto"
)

type service struct{}

func (s service) SendMail(ctx context.Context, request *pb.MailRequest) (*pb.MailResponse, error) {
	log.Printf("mail: %s, content: %s\n", request.Mail, request.Text)
	return &pb.MailResponse{
		Ok: true,
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":8999")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterMailServiceServer(s, &service{})

	reflection.Register(s)

	endpoints := []string{"127.0.0.1:2379"}
	//endpoints := []string{"127.0.0.1:12379","127.0.0.1:22379","127.0.0.1:32379"}
	//endpoints := []string{"10.28.150.23:2379","10.80.60.87:2379","10.31.122.175:2379"}
	rand.Seed(time.Now().UnixNano())
	reg, err := registry.NewService(registry.ServiceInfo{
		Name: "rpc.mail/1",// + cast.ToString(rand.Intn(1000)),
		IP:   "127.0.0.1:8999",
	}, endpoints)

	if err != nil {
		panic(err)
	}

	reg2, err := registry.NewService(registry.ServiceInfo{
		Name: "rpc.test/2",// + cast.ToString(rand.Intn(1000)),
		IP:   "127.0.0.1:7999",
	}, endpoints)

	if err != nil {
		panic(err)
	}

	go reg.Start()
	go reg2.Start()

	log.Println("grpc server start")

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
