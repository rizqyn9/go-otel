package game

import (
	"learn/apps/pkg/server"
)

type GameHandler struct{}

func InitProductSvc() {
	srv := server.CreateServer()
	handler := &GameHandler{}

	handler.routes(srv.Router)
	srv.InitGlobalProvider("game-svc", "localhost:4318")

	srv.Run(3003)
}
