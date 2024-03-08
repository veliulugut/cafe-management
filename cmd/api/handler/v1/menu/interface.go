package menu

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateMenu(c *gin.Context)
	DeleteMenu(c *gin.Context)
	UpdateMenu(c *gin.Context)
	ListMenu(c *gin.Context)
	GetById(c *gin.Context)
}
