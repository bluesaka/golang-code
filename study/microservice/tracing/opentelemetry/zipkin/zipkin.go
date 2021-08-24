package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"time"
)

func initTracer(url string) func() {
	exp, err := zipkin.New(
		url,
		zipkin.WithSDKOptions(tracesdk.WithSampler(tracesdk.AlwaysSample())),
	)
	if err != nil {
		panic(err)
	}

	sp := tracesdk.NewBatchSpanProcessor(exp)

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSpanProcessor(sp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("zipkin-test"),
		)),
	)
	otel.SetTracerProvider(tp)

	return func() {
		_ = tp.Shutdown(context.Background())
	}
}

func main() {
	shutdown := initTracer("http://localhost:9411/api/v2/spans")
	defer shutdown()

	ctx := context.Background()
	tr := otel.GetTracerProvider().Tracer("component-main")
	ctx, span := tr.Start(ctx, "foo", trace.WithSpanKind(trace.SpanKindServer))
	<-time.After(time.Millisecond * 6)
	bar(ctx)
	<-time.After(time.Millisecond * 6)
	span.End()
}

func bar(ctx context.Context) {
	tr := otel.GetTracerProvider().Tracer("component-bar")
	_, span := tr.Start(ctx, "bar")
	<-time.After(time.Millisecond * 6)
	span.End()
}
