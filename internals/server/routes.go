package server

import (
	"github.com/ysrckr/countries-api/internals/server/handlers"
)

func (s *FiberServer) RegisterRoutes() {
	s.App.Get("/", handlers.HelloWorldHandler)
}
