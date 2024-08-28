package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/ysrckr/countries-api/internals/conf"
	"github.com/ysrckr/countries-api/internals/db"
	"github.com/ysrckr/countries-api/internals/server"
)

var seed bool

func init() {
	flag.BoolVar(&seed, "seed", false, "Initilize DB seeding before the server starts")
	flag.Parse()
}

func main() {
	config := conf.Config

	health := db.DB.Health()
	log.Println(health["message"])

	server.Srv.RegisterRoutes()
	port, err := strconv.Atoi(config.Envs["PORT"])
	if err != nil {
		log.Fatalf("error: %w", err)
	}

	shutDownChan := make(chan error, 1)

	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go server.Srv.StartServer(shutdownCtx, shutDownChan, port)

	select {
	case err := <-shutDownChan:
		log.Fatalf("error starting the server. error is %w", err)

	case <-shutdownCtx.Done():
		if !fiber.IsChild() {
			log.Println("Shutting down server...")
			log.Println("Shutting down DB")
		}

		if err = db.DB.Close(shutdownCtx); err != nil {
			log.Fatalln("DB Shutdown Failed:%+v", err)
		}

		if err := server.Srv.ShutDown(); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}

	}
}
