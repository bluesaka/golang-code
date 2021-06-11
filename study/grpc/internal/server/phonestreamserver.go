package server

import (
	"github.com/spf13/cast"
	"go-code/study/grpc/proto"
	"io"
	"log"
	"sync"
	"time"
)

type PhoneStreamServer struct{}

func NewPhoneStreamServer() *PhoneStreamServer {
	return &PhoneStreamServer{}
}

func (p *PhoneStreamServer) ServerStream(req *proto.StreamReq, res proto.PhoneStream_ServerStreamServer) error {
	i := 1
	for {
		res.Send(&proto.StreamResp{
			Value: cast.ToString(i),
			Os: proto.PhoneOS_ANDROID,
		})
		i++
		if i > 3 {
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (p *PhoneStreamServer) ClientStream(req proto.PhoneStream_ClientStreamServer) error {
	for {
		r, err := req.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("recv err:", err)
			break
		}
		log.Println("rece data:", r)
	}
	return nil
}

func (p *PhoneStreamServer) DoubleStream(req proto.PhoneStream_DoubleStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			data, err := req.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("DoubleStream Recv error:", err)
				break
			}
			log.Println(data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			err := req.Send(&proto.StreamResp{
				Value: "server msg",
				Os: proto.PhoneOS_IOS,
			})
			if err != nil {
				log.Println("DoubleStream Send error:", err)
				break
			}
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}