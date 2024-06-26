package observability

import (
	"context"

	jaegerPropagator "go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
)

// NewTracer ...
func NewTracer(ctx context.Context, otlpEndpoint string, svcName string, samplingRate float64) (trace.Tracer, error) {
	traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpointURL(otlpEndpoint))
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(samplingRate)),
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(svcName),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(jaegerPropagator.Jaeger{})

	return otel.Tracer(svcName), nil
}

// NewMockTracer ...
func NewMockTracer(ctx context.Context, svcName string) (trace.Tracer, error) {
	mockTraceExporter := tracetest.NewInMemoryExporter()

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(0.01)),
		sdktrace.WithBatcher(mockTraceExporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(svcName),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(jaegerPropagator.Jaeger{})

	return otel.Tracer(svcName), nil
}
