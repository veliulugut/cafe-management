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

// CreateProduct
// @Summary create product
// @Schemes
// @Description create product
// @Tags product
// @Accept json
// @Produce json
// @Param user body product.CreateProductModel true "Create a new product"
// @Success 201
// @Router /product [post]
func (p *Product) CreateProduct(c *gin.Context) {
	var (
		req product.CreateProductModel
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = p.productService.CreateProduct(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteProduct godoc
// @Summary delete product by id
// @Schemes
// @Description
// @Tags product
// @Param id   path int true "id"
// @Success 204
// @Router /product/{id} [delete]
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

// GetById godoc
// @Summary get product by id
// @Schemes
// @Description
// @Tags product
// @Accept  json
// @Produce  json
// @Param id   path int true "id"
// @Success 200 {object} product.ProductModel "ok"
// @Router /product/get/{id} [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (p *Product) GetById(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um *product.ProductModel

	if um, err = p.productService.GetById(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, um)
}

// ListProduct godoc
// @Summary list all product
// @Description get all product
// @Tags product
// @Accept  json
// @Produce  json
// @Success 200 {object} product.ProductModel "ok"
// @Router /product [get]
func (p *Product) ListProduct(c *gin.Context) {
	list, err := p.productService.ListProduct(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, list)
}

// UpdateProduct godoc
// @Summary update product by id
// @Schemes
// @Description
// @Tags product
// @Accept  json
// @Produce  json
// @Param user body product.UpdateProductModel  true "Update a product"
// @Param id   path int true "id"
// @Success 204
// @Router /product/{id} [post]
func (p *Product) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um product.UpdateProductModel

	if err = p.productService.UpdateProduct(c.Request.Context(), int(id), &um); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}
