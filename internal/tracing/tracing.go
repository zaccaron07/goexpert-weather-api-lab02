package tracing

import (
	context "context"

	gootel "go.opentelemetry.io/otel"
	gootelexportotlp "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	gootelresource "go.opentelemetry.io/otel/sdk/resource"
	gootelsdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func InitTracer(serviceName, collectorURL string) (func(context.Context) error, error) {
	exporter, err := gootelexportotlp.New(
		context.Background(),
		gootelexportotlp.WithEndpoint(collectorURL),
		gootelexportotlp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	resource := gootelresource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
	)
	provider := gootelsdktrace.NewTracerProvider(
		gootelsdktrace.WithBatcher(exporter),
		gootelsdktrace.WithResource(resource),
	)
	gootel.SetTracerProvider(provider)
	gootel.SetTextMapPropagator(propagation.TraceContext{})
	return provider.Shutdown, nil
}
