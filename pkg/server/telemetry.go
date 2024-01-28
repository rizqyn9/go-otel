package server

import (
	ttrace "learn/apps/pkg/telemetry/trace"
	traceExporter "learn/apps/pkg/telemetry/trace/exporter"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func (s *Server) InitGlobalProvider(name, endpoint string) {
	spanExporter := traceExporter.NewOtlp(endpoint)
	tracerProvider, tracerProviderCloseFn, err := ttrace.NewTraceProviderBuilder(name).
		SetExporter(spanExporter).
		Build()
	if err != nil {
		log.Fatal().Err(err).Msgf("failed initializing the tracer provider")
	}

	s.traceProviderCloseFn = append(s.traceProviderCloseFn, tracerProviderCloseFn)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tracerProvider)

	log.Info().Msg("Success initialized otel")
}
