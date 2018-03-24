// +build !windows

package server

import (
	"github.com/majestrate/clicker-rick/config"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func (s *Server) HandleSignals(fname string) {
	chnl := make(chan os.Signal)
	signal.Notify(chnl, os.Interrupt, syscall.SIGHUP)
	for {
		sig := <-chnl
		if sig == os.Interrupt {
			logrus.Info("Closing...")
			s.Close()
			return
		}
		if sig == syscall.SIGHUP {
			logrus.Info("Reloading config file...")
			var conf config.Config
			err := conf.Load(fname)
			if err == nil {
				log.Infof("Loaded %s", fname)
				err = s.Configure(&conf)
			}
			if err == nil {
				logrus.Info("Reconfigure success")
			} else {
				logrus.Errorf("Failed to reconfigure: %s", err.Error())
			}
		}
	}

}
