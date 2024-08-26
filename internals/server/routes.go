package server

func (s *FiberServer) RegisterRoutes() {
	s.RootRoutes()
	s.CountriesRoutes()
}
