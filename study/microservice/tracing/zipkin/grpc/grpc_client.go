package main

import (
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
)

var (
	zkClientReporter reporter.Reporter
	zkClientTracer   opentracing.Tracer
)

func main() {
	if err := initClientZipkin(); err != nil {
		panic(err)
	}
	defer zkClientReporter.Close()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(zkClientTracer),
		),
	}
	conn, err := grpc.Dial("127.0.0.1:9801", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// register service
	// ...
}

func initClientZipkin() error {
	zkClientReporter = zkHttp.NewReporter("http://127.0.0.1:9411/api/v2/spans")
	endpoint, err := zipkin.NewEndpoint("grpc-client", "127.0.0.1:9802")
	if err != nil {
		panic(err)
	}

	nativeTracer, err := zipkin.NewTracer(zkClientReporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return err
	}

	_ = zipkinot.Wrap(nativeTracer)
	//opentracing.SetGlobalTracer(tracer)
	return nil
}