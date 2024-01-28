package game

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

func (h *GameHandler) routes(r *mux.Router) {
	r.Use(otelmux.Middleware("product-svc"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, parentSpan := tracer.Start(r.Context(), "game")

		defer parentSpan.End()
		w.Write([]byte("Game service"))
	})

	r.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		ctx, parentSpan := tracer.Start(r.Context(), "request to product svc")
		requestToProductSvc(ctx)

		defer parentSpan.End()
		w.Write([]byte("Game service call product svc"))
	})
}

func requestToProductSvc(ctx context.Context) {
	ctx, parentSpan := tracer.Start(ctx, "Request to product svc")

	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	bag, _ := baggage.Parse("username=donuts")
	ctx = baggage.ContextWithBaggage(ctx, bag)
	url := "http://localhost:3002"
	// otelhttp.Get(ctx, url)

	tr := otel.Tracer("example/client")
	err := func(ctx context.Context) error {
		ctx, span := tr.Start(ctx, "say hello", trace.WithAttributes(semconv.PeerService("ExampleService")))
		defer span.End()
		req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

		fmt.Printf("Sending request...\n")
		_, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		return err
	}(ctx)
	if err != nil {
		log.Fatal().Err(err)
	}

	defer parentSpan.End()
}
