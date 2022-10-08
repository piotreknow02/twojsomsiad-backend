package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, code int) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": http.StatusText(code),
	})
}
