package controller

import (
	"net/http"
	"strconv"
	"twojsomsiad/model"
	"twojsomsiad/service"
	"twojsomsiad/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GetUser godoc
// @Summary Get user
// @Description Get user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 "Success"
// @Router /user/{id} [get]
func (base *Controller) GetUser(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest)
		return
	}
	user, err := service.FindUserById(base.DB, uint(id))
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user.ToView())
}

// Register godoc
// @Summary Register
// @Description Register Process
// @Tags user
// @Accept json
// @Produce json
// @Success 200 "Success"
// @Router /user [get]
func (base *Controller) GetMyUser(c *gin.Context) {
	id, is := c.Get("id")
	if !is {
		utils.SendError(c, http.StatusBadRequest)
		return
	}
	user, err := service.FindUserById(base.DB, id.(uint))
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user.ToView())
}

// UpdateUser godoc
// @Summary Update user
// @Description Update username/name/surname/password for current user
// @Tags user
// @Accept json
// @Produce json
// @Param Login body model.UserUpdateDTO true "UpdateUser"
// @Success 200 "Success"
// @Router /user [post]
func (base *Controller) UpdateUser(c *gin.Context) {
	id, is := c.Get("id")
	if !is {
		utils.SendError(c, http.StatusBadRequest)
		return
	}
	var data model.UserUpdateDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.SendError(c, http.StatusBadRequest)
		return
	}
	if err := validator.New().Struct(data); err != nil {
		utils.SendError(c, http.StatusBadRequest)
		return
	}
	err := service.UpdateUser(base.DB, id.(uint), data)
	if err != nil {
		utils.SendError(c, http.StatusConflict)
		return
	}
	c.JSON(http.StatusOK, "")
}
