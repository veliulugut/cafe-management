package menu

import (
	"cafe-management/service/menu"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*Menu)(nil)

func NewMenu(menu menu.ServiceMenu) *Menu {
	return &Menu{
		menuService: menu,
	}
}

type Menu struct {
	menuService menu.ServiceMenu
}

// CreateMenu
// @Summary create menu
// @Schemes
// @Description create menu
// @Tags menu
// @Accept json
// @Produce json
// @Param user body menu.CreateMenuModel true "Create a new menu"
// @Success 201
// @Router /menu [post]
func (m *Menu) CreateMenu(c *gin.Context) {
	var (
		req menu.CreateMenuModel
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = m.menuService.CreateMenu(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteMenu godoc
// @Summary delete menu by id
// @Schemes
// @Description
// @Tags menu
// @Param id   path int true "id"
// @Success 204
// @Router /menu/{id} [delete]
func (m *Menu) DeleteMenu(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = m.menuService.DeleteMenu(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetById godoc
// @Summary get menu by id
// @Schemes
// @Description
// @Tags menu
// @Accept  json
// @Produce  json
// @Param id   path int true "id"
// @Success 200 {object} menu.MenuModel "ok"
// @Router /menu/get/{id} [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (m *Menu) GetById(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um *menu.MenuModel

	if um, err = m.menuService.GetById(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, um)
}

// ListMenu godoc
// @Summary list all menu
// @Description get all menu
// @Tags menu
// @Accept  json
// @Produce  json
// @Success 200 {object} menu.MenuModel "ok"
// @Router /menu [get]
func (m *Menu) ListMenu(c *gin.Context) {
	list, err := m.menuService.ListMenu(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, list)
}

// UpdateMenu godoc
// @Summary update menu by id
// @Schemes
// @Description
// @Tags menu
// @Accept  json
// @Produce  json
// @Param user body menu.UpdateMenuModel  true "Update a  menu"
// @Param id   path int true "id"
// @Success 204
// @Router /menu/{id} [post]
func (m *Menu) UpdateMenu(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um menu.UpdateMenuModel

	if err = m.menuService.UpdateMenu(c.Request.Context(), int(id), &um); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
