package server

func (s *Server) linkRoutes() {
	v1Grp := s.router.Group("/v1")
	v1Grp.POST("/login", s.hnd.login.Login)
	v1Grp.POST("/register", s.hnd.login.Register)

	adminGrp := v1Grp.Group("")
	adminGrp.Use(s.mw.Check())
	adminGrp.POST("/user", s.hnd.user.CreateUser)
	adminGrp.POST("/user/:id", s.hnd.user.UpdateUser)
	adminGrp.DELETE("/user/:id", s.hnd.user.DeleteUser)
	adminGrp.GET("/user/get/:id", s.hnd.user.GetById)
	adminGrp.GET("/user/:name", s.hnd.user.GetByUserName)
	adminGrp.GET("/user", s.hnd.user.ListUser)
	adminGrp.POST("menu", s.hnd.menu.CreateMenu)
	adminGrp.DELETE("/menu/:id", s.hnd.menu.DeleteMenu)
	adminGrp.GET("/menu/get/:id", s.hnd.menu.GetById)
	adminGrp.GET("/menu", s.hnd.menu.ListMenu)
	adminGrp.POST("/menu/:id", s.hnd.menu.UpdateMenu)
}
