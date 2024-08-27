package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3/log"
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
	err = config.Server.StartServer(port)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %w", err))
	}
}
