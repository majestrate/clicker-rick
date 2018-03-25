package server

import "github.com/majestrate/apub"

func (s *Server) ServeUserProfile(c *Context) {
	username := c.Param("username")
	u, err := s.db.LocalUser(username)
	if err != nil {
		s.Error(c, err.Error())
		return
	}
	if u == nil {
		s.NotFound(c)
		return
	}
	if WantsActivityPub(c) {
		c.Header("Content-Type", apub.ContentType)
		c.JSON(OK, u)
		return
	}
	p := s.DefaultParams()
	p["User"] = u
	c.HTML(OK, "user-profile.html", p)
}
