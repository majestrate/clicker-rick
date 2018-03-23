package config

type InstanceConfig struct {
	Domain string
	Email  string
}

func (c *InstanceConfig) Section() string {
	return "instance"
}

func (c *InstanceConfig) Load(s *Section) error {
	if s != nil {
		c.Domain = s.ValueOf("domain")
		c.Email = s.ValueOf("admin_email")
	}
	return nil
}

func (c *InstanceConfig) Save(s *Section) {
	s.Add("domain", c.Domain)
	s.Add("admin_email", c.Email)
}
