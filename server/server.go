package server

import (
	"github.com/majestrate/apub"
	"github.com/majestrate/clicker-rick/database"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

type Server struct {
	handler    apub.APubHandler
	db         database.Database
	l          net.Listener
	e          *Engine
	Name       string
	AssetsRoot string
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
	s.handler.Database = s.db

	// setup apub routes
	s.handler.SetupRoutes(func(path string, handler http.Handler) {
		s.e.Any(path, WrapH(handler))
	}, func(subpath string, handler http.Handler) {
		s.e.Group(subpath).Any("/:extra", WrapH(handler))
	})

	// setup app routes
	s.SetupRoutes()

	// serve
	for s.l != nil {
		logrus.Infof("serving at %s as %s", s.l.Addr(), s.Name)
		err := http.Serve(s.l, s.e)
		if err != nil {
			logrus.Warnf("http.Serve(): %s", err.Error())
			time.Sleep(time.Millisecond * 10)
		}
	}
}
