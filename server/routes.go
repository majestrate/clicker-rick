package server

func (s *Server) SetupAppRoutes() {
	s.e.GET("/", s.ServeIndex)
	s.e.Static("/static/", s.AssetsRoot)
}
