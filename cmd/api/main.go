package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/ysrckr/countries-api/internals/conf"
)

var mode string

func init() {
	flag.StringVar(&mode, "mode", "development", "Set env-file path")
	flag.Parse()
}

func main() {
	config := conf.New(mode)

	config.Server.RegisterRoutes()
	port, err := strconv.Atoi(config.Envs["PORT"])
	if err != nil {
		log.Fatalf("error: %w", err)
	}

	shutDownChan := make(chan error, 1)
	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go config.Server.StartServer(shutdownCtx, shutDownChan, port)

	select {
	case err := <-shutDownChan:
		log.Fatalf("error starting the server. error is %w", err)

	case <-shutdownCtx.Done():
		log.Println("Shutting down server...")

		if err := config.Server.ShutDown(); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}

	}
}
