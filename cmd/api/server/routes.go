package server

func (s *Server) linkRoutes() {
	v1Grp := s.router.Group("/v1")
	v1Grp.POST("/login", s.hnd.login.Login)
	v1Grp.POST("/register", s.hnd.login.Register)
	v1Grp.POST("/qrcode", s.hnd.qrcode.GenerateQRCode)
	v1Grp.GET("/qrcode", s.hnd.qrcode.ListQRCode)

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
	adminGrp.GET("/table", s.hnd.tables.ListTable)
	adminGrp.POST("/table/:id", s.hnd.tables.UpdateTable)
	adminGrp.POST("/table", s.hnd.tables.CreateTable)
	adminGrp.DELETE("/table/:id", s.hnd.tables.DeleteTable)
	adminGrp.POST("/product", s.hnd.product.CreateProduct)
	adminGrp.DELETE("/product/:id", s.hnd.product.DeleteProduct)
	adminGrp.GET("/product/get/:id", s.hnd.product.GetById)
	adminGrp.GET("product", s.hnd.product.ListProduct)
	adminGrp.POST("/product/:id", s.hnd.product.UpdateProduct)
	adminGrp.POST("/tabletype", s.hnd.tabletype.CreateTableType)
	adminGrp.DELETE("/tabletype/:id", s.hnd.tabletype.DeleteTableType)
	adminGrp.POST("/tabletype/:id", s.hnd.tabletype.UpdateTableType)
	adminGrp.POST("/ordertype", s.hnd.ordertype.CreateOrderType)
	adminGrp.GET("/ordertype", s.hnd.ordertype.ListOrderType)
	adminGrp.POST("/ordertype/:id", s.hnd.ordertype.UpdateOrderType)
	adminGrp.DELETE("/ordertype/:id", s.hnd.ordertype.DeleteOrderType)
}
