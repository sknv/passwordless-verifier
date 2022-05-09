package tracing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestInit(t *testing.T) {
	type args struct {
		cfg Config
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "when creating a new tracer provider returns an error it returns an error",
			args: args{
				cfg: Config{
					Host: "invalid-host",
					Port: ":invalid-port",
				},
			},
			wantErr: true,
		},
		{
			name: "when init steps are ok it returns no error",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := Init(tt.args.cfg)
			assert.Equalf(t, tt.wantErr, err != nil, "Init(%v)", tt.args.cfg)
		})
	}
}

func TestInitTracer(t *testing.T) {
	type args struct {
		tracerProvider trace.TracerProvider
	}

	dummyProvider, _ := NewJaegerTracerProvider(Config{})

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "when no provider is given it returns an error",
			wantErr: true,
		},
		{
			name: "when a provider is given it returns no error",
			args: args{
				tracerProvider: dummyProvider,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := InitTracer(tt.args.tracerProvider)
			assert.Equalf(t, tt.wantErr, err != nil, "InitTracer(%v)", tt.args.tracerProvider)
		})
	}
}

func TestNewJaegerTracerProvider(t *testing.T) {
	type args struct {
		config Config
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "when an empty config is provided it uses the default values and returns no error",
		},
		{
			name: "when an invalid config is provided it returns an error",
			args: args{
				config: Config{
					Host: "invalid-host",
					Port: ":invalid-port",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewJaegerTracerProvider(tt.args.config)
			assert.Equalf(t, tt.wantErr, err != nil, "NewJaegerTracerProvider(%v)", tt.args.config)
		})
	}
}
