package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

var Srv = New()

type Server interface {
	RegisterRoutes()
	StartServer(context.Context, chan<- error, int)
	ShutDown() error
}

type FiberServer struct {
	*fiber.App
}

func New() Server {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "countries_api",
			AppName:      "Countries Api",
		}),
	}

	return server
}

func (s *FiberServer) StartServer(ctx context.Context, errChan chan<- error, port int) {
	if err := s.App.Listen(fmt.Sprintf(":%d", port), fiber.ListenConfig{
		EnablePrefork:         false,
		DisableStartupMessage: false,
	}); err != nil {
		errChan <- err
	}

	defer close(errChan)
}

func (s *FiberServer) ShutDown() error {
	err := s.App.Shutdown()
	if err != nil {
		return err
	}

	return nil
}
