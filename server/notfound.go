package server

func (s *Server) NotFound(c *Context) {
	c.HTML(NotFound, "error_404.html", s.DefaultParams())
}
