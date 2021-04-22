package server

import (
	"context"
	"go-code/study/grpc/proto"
)

type CarServer struct {
}

func NewCarServer() *CarServer {
	return &CarServer{}
}

func (c *CarServer) List(ctx context.Context, req *proto.CarReq) (*proto.CarListResp, error) {
	return &proto.CarListResp{
		List: []*proto.CarResp{
			{
				Name:  "benz",
				Price: 111,
			},
			{
				Name:  "bmw",
				Price: 222,
			},
			{
				Name:  "audi",
				Price: 333,
			},
		},
	}, nil
}

func (c *CarServer) Query(ctx context.Context, req *proto.CarReq) (*proto.CarResp, error) {
	return &proto.CarResp{
		Name:  req.Name,
		Price: 888,
	}, nil
}

func (c *CarServer) Update(ctx context.Context, req *proto.CarReq) (*proto.CarResp, error) {
	return &proto.CarResp{
		Name:  req.Name,
		Price: 999,
	}, nil
}
