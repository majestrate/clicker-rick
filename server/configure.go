package server

import (
	"github.com/majestrate/clicker-rick/config"
	"github.com/majestrate/clicker-rick/database"
)

func (s *Server) Configure(conf *config.Config) (err error) {

	s.AssetsRoot = conf.Assets.Root
	s.Name = conf.Instance.Domain

	s.db, err = database.New(&conf.DB, s.Name)
	if s.l == nil {
		// TODO: reconfigure listener on sighup if changed
		s.l, err = conf.HTTP.CreateListener()
	}
	return
}
