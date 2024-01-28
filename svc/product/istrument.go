package product

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("learn/app/svc/product")
