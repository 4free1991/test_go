package trace

import (
	"app/lesson4/pkg/logger"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"time"
)

func SetTracerProvider(options *Options) func() {
	ctx := context.Background()

	// 通过grpc协议上送数据

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint("172.28.34.35:4318"), otlptracegrpc.WithInsecure())
	if err != nil {
		logger.Logger.Fatalf("%s: %v", "Failed to create the collector trace exporter", err)
	}

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(options.AppName),
			attribute.Int("service.id", options.ServerId),
			attribute.Int("company.id", options.CompanyId),
			attribute.String("env", ""),
		),
	)
	if err != nil {
		logger.Logger.Fatalf("%s: %v", "failed to create resource", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	// 设置全局propagator为tracecontext（默认不设置）。
	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tracerProvider)

	return func() {
		cxt, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		if err := traceExporter.Shutdown(cxt); err != nil {
			otel.Handle(err)
		}
	}
}
