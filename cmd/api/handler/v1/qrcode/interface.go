package qrcode

import "github.com/gin-gonic/gin"

type Handler interface {
	GenerateQRCode(c *gin.Context)
	DeleteQRCode(c *gin.Context)
	UpdateQRCode(c *gin.Context)
	ListQRCode(c *gin.Context)
}
