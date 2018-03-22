package server

import (
	"github.com/gin-gonic/gin"
	"github.com/majestrate/apub"
	"github.com/majestrate/clicker-rick/database"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
)

type Server struct {
	apub.APubHandler
	db database.Database
	l  net.Listener
}

func (s *Server) Close() error {
	s.db.Close()
	return s.l.Close()
}

func (s *Server) Run() {
	err := s.db.Init()
	if err != nil {
		logrus.Fatalf("failed init db: %s", err.Error())
	}
	s.Database = s.db

	r := gin.Default()

	// setup apub routes
	s.SetupRoutes(func(path string, handler http.Handler) {
		r.Any(path, gin.WrapH(handler))
	}, func(subpath string, handler http.Handler) {
		r.Group(subpath).Any("/:extra", gin.WrapH(handler))
	})
	// setup app routes
	s.SetupAppRoutes(r)
}
