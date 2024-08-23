package server

import "github.com/gofiber/fiber/v3"

func (s *FiberServer) V1Routes() fiber.Router {
	v1 := s.ApiRoutes().Group("/v1")

	return v1
}
