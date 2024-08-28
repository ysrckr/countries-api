package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ysrckr/countries-api/internals/server/handlers"
)

func (s *FiberServer) CountriesRoutes() fiber.Router {
	countries := s.RootRoutes().Group("/countries")

	countries.Get("/", handlers.GetCountriesHandler)

	return countries
}
