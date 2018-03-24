package server

import (
	"path/filepath"
)

func (s *Server) SetupRoutes() {
	s.e.NoRoute(s.NotFound)
	s.e.GET("/", s.ServeIndex)
	s.e.Static("/static/", filepath.Join(s.AssetsRoot, "static"))

	// user profile subrouter
	profile := s.e.Group(ProfilePath)
	profile.Use(s.CORSMiddleware())
	{
		profile.GET("/:username", s.ServeUserProfile)
	}
	// api subrouter
	api := s.e.Group(UIApiPath)
	api.Use(s.CORSMiddleware())
	{
	}
}
