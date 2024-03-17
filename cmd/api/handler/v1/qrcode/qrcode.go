package qrcode

import (
	"cafe-management/service/qrcode"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*QRCode)(nil)

func New(s qrcode.ServiceQrCode) *QRCode {
	return &QRCode{
		qrService: s,
	}
}

type QRCode struct {
	qrService qrcode.ServiceQrCode
}

func (q *QRCode) DeleteQRCode(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = q.qrService.DeleteQRCode(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}

/*

// GenerateQRCode
// @Summary create generate qr code
// @Schemes
// @Description generate qr code
// @Tags qrcode
// @Accept json
// @Produce json
// @Param user body qrcode.CreateQrCodeModel true "Generate a new qr code"
// @Success 201
// @Router /qrcode [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (q *QRCode) GenerateQRCode(c *gin.Context) {
	var (
		req qrcode.CreateQrCodeModel
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, image, err := q.qrService.GenerateQRCode(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url":   url,
		"image": image,
	})
}*/

// GenerateQRCode
// @Summary Generate QR Code
// @Description Generate a new QR code
// @Tags qrcode
// @Accept json
// @Produce image/png
// @Param user body qrcode.CreateQrCodeModel true "Generate a new qr code"
// @Success 201 {file} file "QR code image"
// @Router /qrcode [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (q *QRCode) GenerateQRCode(c *gin.Context) {
	var (
		req qrcode.CreateQrCodeModel
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, imageBytes, err := q.qrService.GenerateQRCode(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "image/png", imageBytes)
}

/*
// ListQRCode godoc
// @Summary List all QR codes
// @Description Get all QR codes
// @Tags qrcode
// @Accept json
// @Produce image/png
// @Success 200 {file} file "OK"
// @Router /qrcode [get]
func (q *QRCode) ListQRCode(c *gin.Context) {
	qrCodes, err := q.qrService.ListQRCode(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"qr_codes": qrCodes,
	})

}*/
// ListQRCode godoc
// @Summary List all QR codes
// @Description Get all QR codes
// @Tags qrcode
// @Accept json
// @Produce json
// @Success 200 {array} string "OK"
// @Router /qrcode [get]
func (q *QRCode) ListQRCode(c *gin.Context) {
	qrCodes, err := q.qrService.ListQRCode(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var qrCodeImages []string
	for _, qrCode := range qrCodes {
		qrCodeImages = append(qrCodeImages, qrCode.Image)
	}

	c.JSON(http.StatusOK, qrCodeImages)
}

func (q *QRCode) UpdateQRCode(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um qrcode.UpdateQrCodeModel

	if err = q.qrService.UpdateQRCode(c.Request.Context(), int(id), &um); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "QR Code updated succesfully",
	})

}
