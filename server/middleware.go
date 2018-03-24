package server

func (s *Server) CORSMiddleware() Middleware {
	return func(c *Context) {
		// TODO: implement cors here
		c.Next()
	}
}
