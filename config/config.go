package config

import (
	"github.com/majestrate/configparser"
)

type Section = configparser.Section

type Configurable interface {
	Section() string
	Load(*Section) error
	Save(*Section) error
}

type Config struct {
	DB DBConfig
}

var DefaultConfig = Config{}

func (c Config) Save(fname string) (err error) {
	confs := []Configurable{
		&c.DB,
	}
	return
}

func (c Config) Load(fname string) error {
	confs := []Configurable{
		&c.DB,
	}
	conf, err = confparser.Load(fname)
	return
}

func Ensure(fname string) (err error) {
	err = DefaultConfig.Save(fname)
	return
}
