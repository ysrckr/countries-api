package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ysrckr/countries-api/internals/db"
)

type FiberServer struct {
	*fiber.App
	db db.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "test",
			AppName:      "test",
		}),

		db: db.New(),
	}

	return server
}
