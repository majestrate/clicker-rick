package server

func (s *Server) DefaultParams() H {
	return H{
		"SiteName": s.Name,
	}
}
