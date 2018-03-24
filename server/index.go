package server

func (s *Server) ServeIndex(c *Context) {
	c.HTML(OK, "index.html", s.DefaultParams())
}
