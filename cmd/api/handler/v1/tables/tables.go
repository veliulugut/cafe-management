package tables

import (
	"cafe-management/service/tables"
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

func (t *Tables) CreateTable(c *gin.Context) {
	panic("unimplemented")
}

func (t *Tables) DeleteTable(c *gin.Context) {
	panic("unimplemented")
}

func (t *Tables) ListTable(c *gin.Context) {
	panic("unimplemented")
}

func (t *Tables) UpdateTable(c *gin.Context) {
	panic("unimplemented")
}
