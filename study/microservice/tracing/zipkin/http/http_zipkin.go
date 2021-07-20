package main

import (
	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"net/http"
)

var (
	zkReporter reporter.Reporter
	zkTracer   opentracing.Tracer
)

func main() {
	if err := initZipkin(); err != nil {
		panic(err)
	}
	defer zkReporter.Close()

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		span := zkTracer.StartSpan("/test")
		defer span.Finish()
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":5005", nil)
}

func initZipkin() error {
	zkReporter = zkHttp.NewReporter("http://localhost:9411/api/v2/spans")
	endpoint, err := zipkin.NewEndpoint("test-service", "127.0.0.1:6666")
	if err != nil {
		return err
	}

	nativeTracer, err := zipkin.NewTracer(zkReporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return err
	}
	zkTracer = zkOt.Wrap(nativeTracer)
	opentracing.SetGlobalTracer(zkTracer)

	return nil
}
