package otel

import (
	"net/http"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const (
	instrumentationName = "github.com/google/go-github/v89/otel"
)

type Transport struct {
	Base   http.RoundTripper
	Tracer trace.Tracer
	Meter  metric.Meter
}

func NewTransport(base http.RoundTripper, opts ...Option) *Transport {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Option func(*Transport)

func WithTracerProvider(tp trace.TracerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMeterProvider(mp metric.MeterProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
