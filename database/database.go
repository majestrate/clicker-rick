package database

import (
	"github.com/majestrate/apub"
	"io"
)

type Database interface {
	io.Closer
	apub.Database
	Init() error
}
