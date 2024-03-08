package product

import (
	"cafe-management/service/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*Product)(nil)

func New(product product.ServiceProduct) *Product {
	return &Product{
		productService: product,
	}
}

type Product struct {
	productService product.ServiceProduct
}

func (p *Product) CreateProduct(c *gin.Context) {
	panic("unimplemented")
}

// DeleteProduct godoc
// @Summary delete menu by id
// @Schemes
// @Description
// @Tags menu
// @Param id   path int true "id"
// @Success 204
// @Router /menu/{id} [delete]
func (p *Product) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = p.productService.DeleteProduct(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (p *Product) GetById(c *gin.Context) {
	panic("unimplemented")
}

func (p *Product) ListProduct(c *gin.Context) {
	panic("unimplemented")
}

func (p *Product) UpdateProduct(c *gin.Context) {
	panic("unimplemented")
}
