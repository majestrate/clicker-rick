package config

type DBConfig struct {
	Type string
	URL  string
}

var DefaultDBConfig = DBConfig{
	Type: "postgres",
	URL:  "host=/var/run/postgres",
}

func (c *DBConfig) Section() string {
	return "database"
}

func (c *DBConfig) Load(s *Section) error {
	return nil
}

func (c *DBConfig) Save(s *Section) {
	return
}
