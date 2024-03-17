package price

import (
	"cafe-management/service/price"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*Price)(nil)

func New(s price.ServicePrice) *Price {
	return &Price{
		priceService: s,
	}
}

type Price struct {
	priceService price.ServicePrice
}

// CreatePrice
// @Summary create price
// @Schemes
// @Description create price
// @Tags price
// @Accept json
// @Produce json
// @Param user body price.CreatePriceModel true "Create a new price"
// @Success 201
// @Router /price [post]
func (p *Price) CreatePrice(c *gin.Context) {
	var (
		req price.CreatePriceModel
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = p.priceService.CreatePrice(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// DeletePrice godoc
// @Summary delete price by id
// @Schemes
// @Description
// @Tags price
// @Param id   path int true "id"
// @Success 204
// @Router /price/{id} [delete]
func (p *Price) DeletePrice(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = p.priceService.DeletePrice(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListPrice godoc
// @Summary list all price
// @Description get all price
// @Tags price
// @Accept  json
// @Produce  json
// @Success 200 {object} price.PriceModel "ok"
// @Router /price [get]
func (p *Price) ListPrice(c *gin.Context) {
	list, err := p.priceService.ListPrice(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, list)
}

// UpdatePrice godoc
// @Summary update price by id
// @Schemes
// @Description
// @Tags price
// @Accept  json
// @Produce  json
// @Param user body price.UpdatePriceModel  true "Update a price"
// @Param id   path int true "id"
// @Success 204
// @Router /price/{id} [post]
func (p *Price) UpdatePrice(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um price.UpdatePriceModel

	if err = p.priceService.UpdatePrice(c.Request.Context(), int(id), &um); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
