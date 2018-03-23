package server

import (
	"path/filepath"
)

func (s *Server) SetupAppRoutes() {
	s.e.GET("/", s.ServeIndex)
	s.e.Static("/static/", filepath.Join(s.AssetsRoot, "static"))
}
