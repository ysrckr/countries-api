package server

import "github.com/gofiber/fiber/v3"

func (s *FiberServer) ApiRoutes() fiber.Router {
	api := s.App.Group("/api")

	return api
}
