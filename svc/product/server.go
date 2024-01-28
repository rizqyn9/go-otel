package product

import (
	"learn/apps/pkg/server"
)

type ProductHandler struct{}

func InitProductSvc() {
	srv := server.CreateServer()
	handler := &ProductHandler{}

	handler.routes(srv.Router)
	srv.InitGlobalProvider("product-svc", "localhost:4318")

	srv.Run(3002)
}
