package database

import (
	"fmt"
	"github.com/majestrate/apub"
	"github.com/majestrate/clicker-rick/config"
	"io"
)

type Database interface {
	io.Closer
	apub.Database
	Init() error
}

func New(conf *config.DBConfig, hostname string) (db Database, err error) {
	switch conf.Type {
	case "postgres":
		db, err = newPostgresDB(conf.URL, hostname)
	default:
		err = fmt.Errorf("no such database type '%s'", conf.Type)
	}
	return
}
