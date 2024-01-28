package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

func (h *ProductHandler) routes(r *mux.Router) {
	r.Use(otelmux.Middleware("product-svc"))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, parentSpan := tracer.Start(r.Context(), "product")

		defer parentSpan.End()
		w.Write([]byte("Product service"))
	})
}
