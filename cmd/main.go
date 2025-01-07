package main

import (
	"log"
	"os"
	"os/signal"
	"ryoeuyo/testgo/internal/app"
	"ryoeuyo/testgo/internal/domain/models"
	"ryoeuyo/testgo/internal/http/rest"
	"syscall"
)

// This is a test application:)
func main() {
	endpoints := []models.EndPoint{
		models.NewEndPoint("/ping", rest.Ping),
		models.NewEndPoint("/pingv2", rest.PingV2),
	}

	app := app.NewApp(endpoints...)
	go app.Srv.MustStartServer()

	log.Printf("Server started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	if err := app.Srv.GracefulStop(); err != nil {
		log.Fatalf("Failed graceful stop: %s", err.Error())
	}
}
