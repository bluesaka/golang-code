package main

import (
	"context"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"go-code/study/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	reporter := zipkinhttp.NewReporter("http://127.0.0.1:9411/api/v2/spans")
	defer reporter.Close()

	endpoint, err := zipkin.NewEndpoint("myService", "127.0.0.1:9802")
	if err != nil {
		panic(err)
	}

	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		panic(err)
	}

	_= zipkinot.Wrap(nativeTracer)
	//opentracing.SetGlobalTracer(tracer)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		//grpc.WithUnaryInterceptor(
		//	otgrpc.OpenTracingClientInterceptor(tracer),
		//),
		grpc.WithStatsHandler(
			zipkingrpc.NewClientHandler(nativeTracer),
		),
	}
	conn, err := grpc.Dial("127.0.0.1:9801", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewCarServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	CarList2(client, ctx)
	CarQuery2(client, ctx)
	CarUpdate2(client, ctx)
}

func CarList2(client proto.CarServiceClient, ctx context.Context) {
	resp, err := client.List(ctx, &proto.CarReq{})
	if err != nil {
		log.Println("carList error:", err)
	}
	log.Printf("%+v\n\n", resp)
}

func CarQuery2(client proto.CarServiceClient, ctx context.Context) {
	resp, err := client.Query(ctx, &proto.CarReq{Name: "benz"})
	if err != nil {
		log.Println("carQuery error:", err)
	}
	log.Printf("%+v\n\n", resp)
}

func CarUpdate2(client proto.CarServiceClient, ctx context.Context) {
	resp, err := client.Update(ctx, &proto.CarReq{Name: "bmw"})
	if err != nil {
		log.Println("carUpdate error:", err)
	}
	log.Printf("%+v\n\n", resp)
}
