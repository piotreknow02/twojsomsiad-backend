package controller

import (
	"net/http"

	"twojsomsiad/model"
	"twojsomsiad/service"
	"twojsomsiad/utils"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login
// @Description Login Process
// @Tags auth
// @Accept json
// @Produce json
// @Param Login body model.AuthLoginDTO true "Login"
// @Success 200 "Success"
// @Router /auth/login [post]

func (base *Controller) IndentifyHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &model.UserClaims{
		ID: claims["id"].(uint),
	}
}

func (base *Controller) Autheticator(c *gin.Context) (interface{}, error) {
	var data model.AuthLoginDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	authorized, user := service.FindUserByCredentials(base.DB, &data)
	if authorized {
		return user, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

func (base *Controller) Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*model.User); ok && v != nil && *v != (model.User{}) {
		return true
	}
	return false
}

func (base *Controller) Unathorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

// Register godoc
// @Summary Register
// @Description Register Process
// @Tags auth
// @Accept json
// @Produce json
// @Param Login body model.AuthRegisterDTO true "Register"
// @Success 200 "Success"
// @Router /auth/register [post]
func (base *Controller) Register(c *gin.Context) {
	var data model.AuthRegisterDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.SendError(c, http.StatusBadRequest)
		return
	}
	err := service.CreateUser(base.DB, &data)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, "")
}
