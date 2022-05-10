package tracing

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// GetTraceID retrieves a trace id from the context span.
func GetTraceID(ctx context.Context) (string, bool) {
	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID()

	return traceID.String(), traceID.IsValid()
}
