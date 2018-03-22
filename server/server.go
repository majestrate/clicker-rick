package server

import (
	"github.com/gin-gonic/gin"
	"github.com/majestrate/apub"
	"github.com/majestrate/clicker-rick/database"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

type Server struct {
	apub.APubHandler
	db database.Database
	l  net.Listener
	e  *gin.Engine
}

func (s *Server) Close() error {
	s.db.Close()
	if s.l != nil {
		s.l.Close()
		s.l = nil
	}
	return nil
}

func (s *Server) Run() {
	err := s.db.Init()
	if err != nil {
		logrus.Fatalf("failed init db: %s", err.Error())
	}
	s.Database = s.db

	s.e = gin.Default()

	// setup apub routes
	s.SetupRoutes(func(path string, handler http.Handler) {
		s.e.Any(path, gin.WrapH(handler))
	}, func(subpath string, handler http.Handler) {
		s.e.Group(subpath).Any("/:extra", gin.WrapH(handler))
	})
	// setup app routes
	s.SetupAppRoutes()

	// serve
	for s.l != nil {
		logrus.Infof("serving at %s", s.l.Addr())
		err := http.Serve(s.l, s.e)
		if err != nil {
			logrus.Warnf("http.Serve(): %s", err.Error())
			time.Sleep(time.Millisecond * 10)
		}
	}
}
