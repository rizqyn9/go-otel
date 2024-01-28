package exporter

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

func NewOtlp(endpoint string) *otlptrace.Exporter {
	ctx := context.Background()
	traceClient := otlptracehttp.NewClient(
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint(endpoint),
	)

	traceExp, err := otlptrace.New(ctx, traceClient)
	if err != nil {
		log.Fatal().Err(err)
	}

	return traceExp
}
