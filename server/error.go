package server

func (s *Server) Error(c *Context, msg string) {
	c.String(InternalError, msg)
}
