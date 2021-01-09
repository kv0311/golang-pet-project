package middleware

import (
	"golang-side-project/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var msg string
		token := c.Query("token")
		if token == "" {
			msg = constant.GetMsg(constant.INVALID_PARAMS)
			code = 401
			c.JSON(http.StatusUnauthorized, gin.H{"code": code, "msg": msg})
			c.Abort()
		}
		c.Next()
	}
}
