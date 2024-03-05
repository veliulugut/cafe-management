package user

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetById(c *gin.Context)
	ListUser(c *gin.Context)
	GetByUserName(c *gin.Context)
}
