package main

import (
	"context"
	"go-code/study/grpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9801", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//carTest(conn)
	streamTest(conn)
}

func streamTest(conn *grpc.ClientConn) {
	client := proto.NewPhoneStreamClient(conn)

	// 服务端推送流
	serverStream, err := client.ServerStream(context.Background(), &proto.StreamReq{
		Name: "name1",
	})
	if err != nil {
		log.Println("ServerStream error:", err)
		return
	}
	for {
		r, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("ServerStream Recv error:", err)
			break
		}
		log.Println("phone os:", getPhoneOs(r.Os))
		log.Printf("ServerStream Recv: %+v", r)
	}

	// 客户端推送流
	clientStream, err := client.ClientStream(context.Background())
	if err != nil {
		log.Println("ClientStream error:", err)
		return
	}
	i := 1
	for {
		clientStream.Send(&proto.StreamReq{
			Name: "name2",
		})
		i++
		if i > 3 {
			break
		}
		time.Sleep(time.Second)
	}

	// 双向流
	doubleStream, err := client.DoubleStream(context.Background())
	if err != nil {
		log.Println("DoubleStream error:", err)
		return
	}

	go func() {
		for {
			data, err := doubleStream.Recv()
			if err != nil {
				log.Println("DoubleStream Recv error:", err)
				break
			}
			log.Println(data)
		}
	}()

	go func() {
		for {
			err := doubleStream.Send(&proto.StreamReq{
				Name: "name3",
			})
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("DoubleStream Send error:", err)
				break
			}
			time.Sleep(time.Second)
		}
	}()

	select{}
}

func getPhoneOs(os proto.PhoneOS) string {
	switch os {
	case proto.PhoneOS_IOS:
		return "ios -- 1"
	case proto.PhoneOS_ANDROID:
		return "android -- 2"
	default:
		return ""
	}
}

func carTest(conn *grpc.ClientConn) {
	client := proto.NewCarServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	carList(client, ctx)
	carQuery(client, ctx)
	carUpdate(client, ctx)
}

func carList(client proto.CarServiceClient, ctx context.Context) {
	resp, err := client.List(ctx, &proto.CarReq{})
	if err != nil {
		log.Println("carList error:", err)
	}
	log.Printf("%+v\n\n", resp)
}

func carQuery(client proto.CarServiceClient, ctx context.Context) {
	resp, err := client.Query(ctx, &proto.CarReq{Name: "benz"})
	if err != nil {
		log.Println("carQuery error:", err)
	}
	log.Printf("%+v\n\n", resp)
}

func carUpdate(client proto.CarServiceClient, ctx context.Context) {
	resp, err := client.Update(ctx, &proto.CarReq{Name: "bmw"})
	if err != nil {
		log.Println("carUpdate error:", err)
	}
	log.Printf("%+v\n\n", resp)
}
