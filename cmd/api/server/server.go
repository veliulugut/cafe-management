package server

import (
	"cafe-management/cmd/api/handler/v1/login"
	"cafe-management/cmd/api/handler/v1/menu"
	"cafe-management/cmd/api/handler/v1/qrcode"
	"cafe-management/cmd/api/handler/v1/user"
	"cafe-management/cmd/api/middlewares/auth"
	"cafe-management/ent"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func New(port int) *Server {
	return &Server{
		port: port,
	}
}

type Server struct {
	router   *gin.Engine
	dbClient *ent.Client
	port     int
	hnd      Handlers
	mw       auth.Auth
}

type Handlers struct {
	user   user.Handler
	login  login.Handler
	menu   menu.Handler
	qrcode qrcode.Handler
}

func (s *Server) Init() error {
	s.router = gin.New()
	err := s.initDB()
	if err != nil {
		return err
	}
	err = s.initHandlers()
	if err != nil {
		return fmt.Errorf("init :%w", err)
	}
	s.initMiddlewares()
	s.linkRoutes()

	return nil
}

func (s *Server) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := s.router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return fmt.Errorf("server run :%w", err)
	}

	return nil

}
