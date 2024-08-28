package server

import "github.com/gofiber/fiber/v3"

func (s *FiberServer) RootRoutes() fiber.Router {
	root := s.V1Routes().Group("/")

	return root
}
