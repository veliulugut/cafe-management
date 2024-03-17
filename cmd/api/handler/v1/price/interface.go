package price

import "github.com/gin-gonic/gin"

type Handler interface {
	CreatePrice(c *gin.Context)
	DeletePrice(c *gin.Context)
	UpdatePrice(c *gin.Context)
	ListPrice(c *gin.Context)
}
