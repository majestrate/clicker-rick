package config

import "net"

type HTTPConfig struct {
	Type string
	Addr string
}

func (c *HTTPConfig) Section() string {
	return "listener"
}

func (c *HTTPConfig) Load(s *Section) error {
	if s != nil {
		c.Type = s.Get("type", "tcp")
		c.Addr = s.Get("addr", ":3000")
	}
	return nil
}

func (c *HTTPConfig) Save(s *Section) {
	if s != nil {
		s.Add("type", c.Type)
		s.Add("addr", c.Addr)
	}
}

func (c *HTTPConfig) CreateListener() (net.Listener, error) {
	return net.Listen(c.Type, c.Addr)
}
