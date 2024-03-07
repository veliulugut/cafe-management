package server

func (s *Server) linkRoutes() {
	v1Grp := s.router.Group("/v1")
	v1Grp.POST("/login", s.hnd.login.Login)
	v1Grp.POST("/register", s.hnd.login.Register)
	v1Grp.POST("/user", s.hnd.user.CreateUser)

}
