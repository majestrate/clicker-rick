package main

import (
	"github.com/majestrate/clicker-rick/config"
	"github.com/majestrate/clicker-rick/server"
	"github.com/sirupsen/logrus"
	"os"
)

var serv server.Server

func main() {
	fname := "config.ini"
	if len(os.Args) > 1 {
		fname = os.Args[1]
	}
	err := config.Ensure(fname)
	if err != nil {
		logrus.Fatalf("failed to generate configs: %s", err.Error())
	}
	var conf config.Config
	err = conf.Load(fname)
	if err != nil {
		logrus.Fatalf("failed to load config: %s", err.Error())
	}
	err = serv.Configure(&conf)
	if err != nil {
		logrus.Fatalf("failed to configure: %s", err.Error())
	}
	go serv.Run()
	serv.HandleSignals(fname)
}
