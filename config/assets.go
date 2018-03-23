package config

type AssetsConfig struct {
	Root string
}

func (c *AssetsConfig) Section() string {
	return "assets"
}

func (c *AssetsConfig) Load(s *Section) error {
	if s != nil {
		c.Root = s.ValueOf("dir")
	}
	return nil
}

func (c *AssetsConfig) Save(s *Section) {
	s.Add("dir", c.Root)
}
