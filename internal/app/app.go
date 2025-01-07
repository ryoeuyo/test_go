package app

import (
	"ryoeuyo/testgo/internal/app/server"
	"ryoeuyo/testgo/internal/domain/models"
	"time"
)

type App struct {
	Srv *server.Server
}

func NewApp(endpoints ...models.EndPoint) *App {
	srv := server.New(&server.ServerOptions{
		Timeout: 5 * time.Second,
		Port:    8081,
	}, endpoints...)

	return &App{
		Srv: srv,
	}
}
