package tracing

import (
	"errors"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var ErrEmptyProvider = errors.New("empty tracer provider")

// Config is a jaeger tracer config.
type Config struct {
	Host        string
	Port        string
	ServiceName string
	Ratio       float64
}

// Init the global tracer from the provided config.
func Init(cfg Config) error {
	provider, err := NewJaegerTracerProvider(cfg)
	if err != nil {
		return fmt.Errorf("init jaeger tracer provider: %w", err)
	}

	if err = InitTracer(provider); err != nil {
		return fmt.Errorf("init tracer: %w", err)
	}

	return nil
}

// NewJaegerTracerProvider return a jaeger exporter for the provided url.
func NewJaegerTracerProvider(config Config) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter from config
	exporter, err := jaeger.New(
		jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost(config.Host),
			jaeger.WithAgentPort(config.Port),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("new jaeger exporter: %w", err)
	}

	provider := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter), // always be sure to batch
		tracesdk.WithSampler(tracesdk.TraceIDRatioBased(config.Ratio)),
		// Record information about this application in a resource
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ServiceName),
		)),
	)
	return provider, nil
}

// InitTracer inits and sets the global tracer for the provided tracer provider and name.
func InitTracer(tracerProvider trace.TracerProvider) error {
	if tracerProvider == nil {
		return ErrEmptyProvider
	}

	// Register globals
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)
	return nil
}
