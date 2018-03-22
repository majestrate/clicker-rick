// +build windows

package server

import (
	"os"
	"os/signal"
)

func (s *Server) HandleSignals(fname string) {
	chnl := make(chan os.Signal)
	signal.Notify(chnl, os.Interrupt)
	for {
		sig <- chnl
		if sig == os.Interrupt {
			s.Close()
			return
		}
	}
}
