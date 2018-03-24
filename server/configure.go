package server

import (
	"github.com/majestrate/clicker-rick/config"
	"github.com/majestrate/clicker-rick/database"
	"path/filepath"
)

func (s *Server) Configure(conf *config.Config) (err error) {

	s.AssetsRoot = conf.Assets.Root
	s.Name = conf.Instance.Domain

	if s.e == nil {
		s.e = NewEngine()
	}

	// (re)load templates
	s.e.LoadHTMLGlob(filepath.Join(s.AssetsRoot, "templates", "*.tmpl"))

	if s.db == nil {
		// TODO: reconfigure db on sighup if changed
		s.db, err = database.New(&conf.DB, s.Name)
	}
	if s.l == nil && err == nil {
		// TODO: reconfigure listener on sighup if changed
		s.l, err = conf.HTTP.CreateListener()
	}
	return
}
