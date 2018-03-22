package config

import (
	"github.com/majestrate/configparser"
)

type Section = configparser.Section

type Configurable interface {
	Section() string
	Load(*Section) error
	Save(*Section)
}

type Config struct {
	DB DBConfig
}

var DefaultConfig = Config{
	DB: DefaultDBConfig,
}

func (c Config) Save(fname string) (err error) {
	confs := []Configurable{
		&c.DB,
	}
	conf := configparser.NewConfiguration()
	for idx := range confs {
		s := conf.NewSection(confs[idx].Section())
		confs[idx].Save(s)
	}
	err = configparser.Save(conf, fname)
	return
}

func (c Config) Load(fname string) error {
	confs := []Configurable{
		&c.DB,
	}
	conf, err := configparser.Read(fname)
	for idx := range confs {
		s, _ := conf.Section(confs[idx].Section())
		err = confs[idx].Load(s)
		if err != nil {
			return err
		}
	}
	return err
}

func Ensure(fname string) (err error) {
	err = DefaultConfig.Save(fname)
	return
}
