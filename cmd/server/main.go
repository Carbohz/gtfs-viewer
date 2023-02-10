package main

import (
	"context"
	"github.com/Carbohz/gtfs-viewer/api/rest/downloader"
	_ "github.com/Carbohz/gtfs-viewer/service/server/v1"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, ctxCancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer ctxCancel()

	// service := v1.NewService()

	apiServer, err := downloader.NewAPIServer()
	if err != nil {
		log.Fatalf("Failed to create a server: %v", err)
	}

	err = apiServer.Run(ctx)
	if err != nil {
		return
	}
}
