package router

import (
	"golang-side-project/handler"
	"golang-side-project/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.POST("/user/create", handler.CreateAccount)
	r.POST("/user/loginByUser", handler.LoginByUser)
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.GET("/info")
	}
	return r
}
