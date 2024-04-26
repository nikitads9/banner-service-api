package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

// mockTracer ...
type mockTracer struct {
	embedded.Tracer
}

// NewMockTracer ...
func NewMockTracer() mockTracer {
	return mockTracer{}
}

// Start ...
func (m mockTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return ctx, NewMockSpan()
}

// mockSpan ...
type mockSpan struct {
	embedded.Span
}

// NewMockSpan ...
func NewMockSpan() mockSpan {
	return mockSpan{}
}

func (ms mockSpan) End(options ...trace.SpanEndOption) {

}

// AddEvent adds an event with the provided name and options.
func (ms mockSpan) AddEvent(name string, options ...trace.EventOption) {

}

// AddLink adds a link.
// Adding links at span creation using WithLinks is preferred to calling AddLink
// later, for contexts that are available during span creation, because head
// sampling decisions can only consider information present during span creation.
func (ms mockSpan) AddLink(link trace.Link) {

}

// IsRecording returns the recording state of the Span. It will return
// true if the Span is active and events can be recorded.
func (ms mockSpan) IsRecording() bool {
	return true
}

// RecordError will record err as an exception span event for this span. An
// additional call to SetStatus is required if the Status of the Span should
// be set to Error, as this method does not change the Span status. If this
// span is not being recorded or err is nil then this method does nothing.
func (ms mockSpan) RecordError(err error, options ...trace.EventOption) {

}

// SpanContext returns the SpanContext of the Span. The returned SpanContext
// is usable even after the End method has been called for the Span.
func (ms mockSpan) SpanContext() trace.SpanContext {
	return trace.NewSpanContext(trace.SpanContextConfig{})
}

// SetStatus sets the status of the Span in the form of a code and a
// description, provided the status hasn't already been set to a higher
// value before (OK > Error > Unset). The description is only included in a
// status when the code is for an error.
func (ms mockSpan) SetStatus(code codes.Code, description string) {

}

// SetName sets the Span name.
func (ms mockSpan) SetName(name string) {

}

// SetAttributes sets kv as attributes of the Span. If a key from kv
// already exists for an attribute of the Span it will be overwritten with
// the value contained in kv.
func (ms mockSpan) SetAttributes(kv ...attribute.KeyValue) {

}

// TracerProvider returns a TracerProvider that can be used to generate
// additional Spans on the same telemetry pipeline as the current Span.
func (ms mockSpan) TracerProvider() trace.TracerProvider {
	mockTraceExporter := tracetest.NewInMemoryExporter()

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(0.01)),
		sdktrace.WithSyncer(mockTraceExporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(""),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp
}
