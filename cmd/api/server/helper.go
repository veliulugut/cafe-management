package server

import (
	loginhnd "cafe-management/cmd/api/handler/v1/login"
	menuhnd "cafe-management/cmd/api/handler/v1/menu"
	ordertypehnd "cafe-management/cmd/api/handler/v1/ordertype"
	producthnd "cafe-management/cmd/api/handler/v1/product"
	qrhnd "cafe-management/cmd/api/handler/v1/qrcode"
	"cafe-management/cmd/api/handler/v1/tables"
	tablestypehnd "cafe-management/cmd/api/handler/v1/tablestype"
	userhnd "cafe-management/cmd/api/handler/v1/user"
	"cafe-management/cmd/api/middlewares/auth"
	_ "cafe-management/docs"
	"cafe-management/ent"
	golangjwt "cafe-management/pkg/jwt/golang_jwt"
	"cafe-management/pkg/passwd/bcrypt"
	"cafe-management/pkg/repository/entadp"
	loginsrv "cafe-management/service/login"
	menusrv "cafe-management/service/menu"
	ordertypesrv "cafe-management/service/ordertype"
	productsrv "cafe-management/service/product"
	qrsrv "cafe-management/service/qrcode"
	tablesrv "cafe-management/service/tables"
	tablestypesrv "cafe-management/service/tablestype"
	usersrv "cafe-management/service/user"
	"context"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*func (s *Server) initDB() error {
	var err error
	s.dbClient, err = ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return fmt.Errorf("ent open :%w", err)
	}

	if err = s.dbClient.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("schema create  :%w", err)
	}

	return nil

}*/

func (s *Server) initDB() error {
	var err error
	s.dbClient, err = ent.Open("mysql", "root:123turkTR562@tcp(localhost)/cafedb")
	if err != nil {
		return fmt.Errorf("ent open: %w", err)
	}

	if err = s.dbClient.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("schema create: %w", err)
	}

	return nil
}

func (s *Server) initMiddlewares() {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	jwt := golangjwt.New("verysecretkey", 24)
	s.mw = auth.New(jwt)
}

func (s *Server) initHandlers() error {
	secret := "verysecretkey"
	repo := entadp.NewRepository(s.dbClient)

	//service
	bc := bcrypt.NewBcrypt(secret, 10)
	j := golangjwt.New(secret, 24)
	userService := usersrv.New(repo, bc)
	loginService := loginsrv.New(repo, j, bc)
	menuService := menusrv.New(repo)
	qrService := qrsrv.New(repo)
	tableService := tablesrv.New(repo)
	productService := productsrv.New(repo)
	tablestypeService := tablestypesrv.New(repo)
	ordertypeService := ordertypesrv.New(repo)

	//handlers
	s.hnd.user = userhnd.NewUser(userService)
	s.hnd.login = loginhnd.New(loginService)
	s.hnd.menu = menuhnd.NewMenu(menuService)
	s.hnd.qrcode = qrhnd.New(qrService)
	s.hnd.tables = tables.New(tableService)
	s.hnd.product = producthnd.New(productService)
	s.hnd.tabletype = tablestypehnd.New(tablestypeService)
	s.hnd.ordertype = ordertypehnd.New(ordertypeService)
	return nil
}
