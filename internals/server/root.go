package server

import "github.com/ysrckr/countries-api/internals/server/handlers"

func (s *FiberServer) RootRoutes() {
	root := s.V1Routes()

	root.Get("/", handlers.HelloWorldHandler)
}
