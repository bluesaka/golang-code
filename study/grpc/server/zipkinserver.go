package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"go-code/study/grpc/internal/server"
	"go-code/study/grpc/proto"
	"google.golang.org/grpc"
	"net"
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
	tracer := opentracing.Wrap(nativeTracer)

	_ = grpc.StatsHandler(zipkingrpc.NewClientHandler(nativeTracer))

	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			otgrpc.OpenTracingServerInterceptor(tracer),
		),
		//grpc.StatsHandler(zipkingrpc.NewClientHandler(nativeTracer)),
	}

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterCarServiceServer(grpcServer, server.NewCarServer())

	lis, err := net.Listen("tcp", ":9801")
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(lis)
}
