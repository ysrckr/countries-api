package server

import (
	"github.com/gofiber/fiber/v3"
)

type Server interface {
	RegisterRoutes()
	GetApp() *fiber.App
}

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "test",
			AppName:      "test",
		}),
	}

	return server
}

func (s *FiberServer) GetApp() *fiber.App {
	return s.App
}
