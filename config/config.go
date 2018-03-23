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
	DB       DBConfig
	HTTP     HTTPConfig
	Instance InstanceConfig
	Assets   AssetsConfig
}

var DefaultConfig = Config{
	DB: DefaultDBConfig,
}

func (c *Config) Configurables() []Configurable {
	return []Configurable{
		&c.DB, &c.HTTP, &c.Instance, &c.Assets,
	}
}

func (c *Config) Save(fname string) (err error) {
	confs := c.Configurables()
	conf := configparser.NewConfiguration()
	for idx := range confs {
		s := conf.NewSection(confs[idx].Section())
		confs[idx].Save(s)
	}
	err = configparser.Save(conf, fname)
	return
}

func (c *Config) Load(fname string) error {
	confs := c.Configurables()
	conf, err := configparser.Read(fname)
	for idx := range confs {
		s, err := conf.Section(confs[idx].Section())
		if err != nil {
			return err
		}
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
