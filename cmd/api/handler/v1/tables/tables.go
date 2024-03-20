package tables

import (
	"cafe-management/service/tables"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*Tables)(nil)

func New(table tables.ServiceTables) *Tables {
	return &Tables{
		tableService: table,
	}
}

type Tables struct {
	tableService tables.ServiceTables
}

// CreateTable godoc
// @Summary create a new table
// @Schemes
// @Description create table
// @Tags table
// @Accept json
// @Produce json
// @Param user body tables.CreateTableModel true "Create a new table"
// @Success 201
// @Router /table [post]
func (t *Tables) CreateTable(c *gin.Context) {
	var (
		req tables.CreateTableModel
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = t.tableService.CreateTable(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteTable godoc
// @Summary delete table by id
// @Schemes
// @Description
// @Tags table
// @Param id   path int true "id"
// @Success 204
// @Router /table/{id} [delete]
func (t *Tables) DeleteTable(c *gin.Context) {
	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = t.tableService.DeleteTable(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListTable godoc
// @Summary list all table
// @Description get all table
// @Tags table
// @Accept  json
// @Produce  json
// @Success 200 {object} tables.TableModel "ok"
// @Router /table [get]
func (t *Tables) ListTable(c *gin.Context) {
	list, err := t.tableService.ListTable(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}

// UpdateTable godoc
// @Summary update table by id
// @Schemes
// @Description
// @Tags table
// @Accept  json
// @Produce  json
// @Param user body tables.UpdateTableModel  true "Update a table"
// @Param id   path int true "id"
// @Success 204
// @Router /table/{id} [post]
func (t *Tables) UpdateTable(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um tables.UpdateTableModel

	if err = t.tableService.UpdateTable(c.Request.Context(), int(id), &um); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
