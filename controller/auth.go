package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}
