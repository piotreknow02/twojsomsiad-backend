package utils

import (
	"net/http"
	"twojsomsiad/config"

	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, code int, err error) {
	if config.Conf.IsDev {
		c.JSON(code, gin.H{
			"code":    code,
			"message": http.StatusText(code),
			"error":   err.Error(),
		})
	} else {
		c.JSON(code, gin.H{
			"code":    code,
			"message": http.StatusText(code),
		})
	}
}
