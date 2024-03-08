package user

import (
	"cafe-management/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*User)(nil)

func NewUser(us user.ServiceUser) *User {
	return &User{
		userService: us,
	}
}

type User struct {
	userService user.ServiceUser
}

// CreateUser
// @Summary create user
// @Schemes
// @Description create user
// @Tags user
// @Accept json
// @Produce json
// @Param user body user.CreateUserModel true "Create a new user"
// @Success 201
// @Router /user [post]
func (u *User) CreateUser(c *gin.Context) {
	var (
		userdb user.CreateUserModel
		err    error
	)

	if err = c.ShouldBindJSON(&userdb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = u.userService.CreateUser(c, &userdb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteUser godoc
// @Summary delete user by id
// @Schemes
// @Description
// @Tags user
// @Param id   path int true "id"
// @Success 204
// @Router /user/{id} [delete]
func (u *User) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  uint64
		err error
	)

	if id, err = strconv.ParseUint(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = u.userService.DeleteUser(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}

// GetById godoc
// @Summary get user by id
// @Schemes
// @Description
// @Tags user
// @Accept  json
// @Produce  json
// @Param id   path int true "id"
// @Success 200 {object} user.UserModel "ok"
// @Router /user/get/{id} [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (u *User) GetById(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um *user.UserModel

	if um, err = u.userService.GetById(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, um)
}

// GetByUserName godoc
// @Summary get user by name
// @Schemes
// @Description
// @Tags user
// @Accept  json
// @Produce  json
// @Param name   path string true "name"
// @Success 204 {object} user.UserModel
// @Router /user/{name} [get]
func (u *User) GetByUserName(c *gin.Context) {
	name := c.Param("name")

	names, err := u.userService.GetByUserName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, names)

}

// ListUser godoc
// @Summary list all users
// @Description get all users
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} user.UserModel "ok"
// @Router /user [get]
func (u *User) ListUser(c *gin.Context) {

	ums, err := u.userService.ListUser(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ums)
}

// UpdateUser godoc
// @Summary update user by id
// @Schemes
// @Description
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.UpdateUserModel  true "Update a  user"
// @Param id   path int true "id"
// @Success 204
// @Router /user/{id} [post]
func (u *User) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  uint64
		err error
	)

	if id, err = strconv.ParseUint(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um user.UpdateUserModel

	if err = c.ShouldBindJSON(&um); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = u.userService.UpdateUser(c.Request.Context(), int(id), &um); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}
