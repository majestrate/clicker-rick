package server

import (
	"net/http"
)

func (s *Server) ServeIndex(c *Context) {
	c.HTML(http.StatusOK, "index.html", H{
		"SiteName": s.Name,
	})
}
