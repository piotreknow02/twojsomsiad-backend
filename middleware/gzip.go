package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func GZIP() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression)
}
