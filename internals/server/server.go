package server

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type Server interface {
	RegisterRoutes()
	StartServer(int) error
	ShutDown() error
}

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "countries_api",
			AppName:      "Countries Api",
		}),
	}

	return server
}

func (s *FiberServer) StartServer(port int) error {
	if err := s.App.Listen(fmt.Sprintf(":%d", port), fiber.ListenConfig{
		EnablePrefork:         true,
		DisableStartupMessage: false,
	}); err != nil {
		return err
	}

	return nil
}

func (s *FiberServer) ShutDown() error {
	err := s.App.Shutdown()
	if err != nil {
		return err
	}

	return nil
}
