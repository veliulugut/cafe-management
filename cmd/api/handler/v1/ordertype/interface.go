package ordertype

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateOrderType(c *gin.Context)
	DeleteOrderType(c *gin.Context)
	UpdateOrderType(c *gin.Context)
	ListOrderType(c *gin.Context)
}
