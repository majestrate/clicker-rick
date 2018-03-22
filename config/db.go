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
	if s != nil {
		c.Type = s.ValueOf("type")
		c.URL = s.ValueOf("url")
	}
	return nil
}

func (c *DBConfig) Save(s *Section) {
	s.Add("type", c.Type)
	s.Add("url", c.URL)
	return
}
