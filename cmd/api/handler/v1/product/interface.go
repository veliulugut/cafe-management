package product

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	ListProduct(c *gin.Context)
	GetById(c *gin.Context)
}
