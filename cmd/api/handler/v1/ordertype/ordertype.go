package ordertype

import (
	"cafe-management/service/ordertype"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func New(ordertype ordertype.ServiceOrderType) *OrderType {
	return &OrderType{
		orderType: ordertype,
	}
}

var _ Handler = (*OrderType)(nil)

type OrderType struct {
	orderType ordertype.ServiceOrderType
}

// CreateOrderType godoc
// @Summary create a new order type
// @Schemes
// @Description create order type
// @Tags order type
// @Accept json
// @Produce json
// @Param user body ordertype.CreateOrderTypeModel true "Create a new order type"
// @Success 201
// @Router /ordertype [post]
func (o *OrderType) CreateOrderType(c *gin.Context) {
	var (
		req ordertype.CreateOrderTypeModel
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = o.orderType.CreateOrderType(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteOrderType godoc
// @Summary delete order type by id
// @Schemes
// @Description
// @Tags order type
// @Param id   path int true "id"
// @Success 204
// @Router /ordertype/{id} [delete]
func (o *OrderType) DeleteOrderType(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = o.orderType.DeleteOrderType(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListOrderType godoc
// @Summary list all ordertype
// @Description get all ta
// @Tags order type
// @Accept  json
// @Produce  json
// @Success 200 {object} ordertype.ListOrderType "ok"
// @Router /ordertype [get]
func (o *OrderType) ListOrderType(c *gin.Context) {
	list, err := o.orderType.ListOrderType(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}

// UpdateOrderType godoc
// @Summary update order type by id
// @Schemes
// @Description
// @Tags order type
// @Accept  json
// @Produce  json
// @Param user body ordertype.UpdateOrderType  true "Update a order type"
// @Param id   path int true "id"
// @Success 204
// @Router /ordertype/{id} [post]
func (o *OrderType) UpdateOrderType(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um ordertype.UpdateOrderType

	if err = o.orderType.UpdateOrderType(c.Request.Context(), int(id), &um); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
