package tablestype

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateTableType(c *gin.Context)
	DeleteTableType(c *gin.Context)
	UpdateTableType(c *gin.Context)
}
