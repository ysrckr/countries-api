package server

import "github.com/ysrckr/countries-api/internals/server/handlers"

func (s *FiberServer) CountriesRoutes() {
	countries := s.V1Routes().Group("/countries")

	countries.Get("/", handlers.HelloWorldHandler)
}
