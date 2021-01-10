package handler

import (
	"golang-side-project/common/app"

	"github.com/gin-gonic/gin"
)

var c *gin.Context

//ResponseHTTP response http
func ResponseHTTP(httpCode int, msg string, data interface{}) {
	response := app.Gin{C: c}
	response.ResponseHttp(httpCode, msg, data)
}
