package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/nktknshn/go-ergo-handler-example/app"

	"github.com/nktknshn/go-ergo-handler-example/libs/http_server"
)

func main() {
	ctx := context.Background()
	app := app.NewApp()
	err := app.Init(ctx)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	err = initState(ctx, app)
	if err != nil {
		log.Fatalf("failed to initialize state: %v", err)
	}
	server := http_server.NewHttpServer()
	app.AddHttpHandlers(server)
	slog.Info("starting server")
	server.Start()
}
