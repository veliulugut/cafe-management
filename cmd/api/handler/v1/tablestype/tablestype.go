package tablestype

import (
	"cafe-management/service/tablestype"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*TablesType)(nil)

func New(t tablestype.ServiceTablestype) *TablesType {
	return &TablesType{
		tablesType: t,
	}
}

type TablesType struct {
	tablesType tablestype.ServiceTablestype
}

// CreateTableType godoc
// @Summary create a new table type
// @Schemes
// @Description create table type
// @Tags table type
// @Accept json
// @Produce json
// @Param user body tablestype.CreateTableTypeModel true "Create a new table"
// @Success 201
// @Router /tabletype [post]
func (t *TablesType) CreateTableType(c *gin.Context) {
	var (
		req tablestype.CreateTableTypeModel
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = t.tablesType.CreateTableType(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteTableType godoc
// @Summary delete table type by id
// @Schemes
// @Description
// @Tags table type
// @Param id   path int true "id"
// @Success 204
// @Router /tabletype/{id} [delete]
func (t *TablesType) DeleteTableType(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := t.tablesType.DeleteTableType(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateTableType godoc
// @Summary update table type by id
// @Schemes
// @Description
// @Tags table type
// @Accept  json
// @Produce  json
// @Param user body tablestype.UpdateTableTypeModel  true "Update a table"
// @Param id   path int true "id"
// @Success 204
// @Router /tabletype/{id} [post]
func (t *TablesType) UpdateTableType(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um tablestype.UpdateTableTypeModel

	if err = t.tablesType.UpdateTableType(c.Request.Context(), int(id), &um); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
