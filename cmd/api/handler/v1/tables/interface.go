package tables

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateTable(c *gin.Context)
	UpdateTable(c *gin.Context)
	DeleteTable(c *gin.Context)
	ListTable(c *gin.Context)
}
