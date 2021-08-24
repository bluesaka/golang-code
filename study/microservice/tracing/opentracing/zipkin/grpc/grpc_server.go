package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	//zkGrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"github.com/openzipkin/zipkin-go/reporter"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"net"
)

var (
	zkServerReporter reporter.Reporter
	zkServerTracer   opentracing.Tracer
)

func main() {
	// init Zipkin
	if err := initServerZipkin(); err != nil {
		panic(err)
	}
	defer zkServerReporter.Close()

	// server options
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			otgrpc.OpenTracingServerInterceptor(zkServerTracer),
		),
	}

	// new grpc server
	grpcServer := grpc.NewServer(opts...)

	// register service
	// ...

	lis, err := net.Listen("tcp", ":9801")
	if err != nil {
		panic(err)
	}

	// serve
	grpcServer.Serve(lis)
}

func initServerZipkin() error {
	zkServerReporter = zkHttp.NewReporter("http://127.0.0.1:9411/api/v2/spans")
	endpoint, err := zipkin.NewEndpoint("grpc-server", "127.0.0.1:9802")
	if err != nil {
		return err
	}

	nativeTracer, err := zipkin.NewTracer(zkServerReporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return err
	}

	zkServerTracer = zkOt.Wrap(nativeTracer)
	return nil
}
