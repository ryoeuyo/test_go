package app

import (
	"ryoeuyo/testgo/internal/app/server"
	"time"
)

type App struct {
	Srv *server.Server
}

func NewApp() *App {
	srv := server.New(&server.ServerOptions{
		Timeout: 5 * time.Second,
		Host:    "localhost",
		Port:    8081,
	})

	return &App{
		Srv: srv,
	}
}
