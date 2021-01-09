package main

import (
	"golang-side-project/router"
	"net/http"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	routerInit := router.InitRouter()
	server := &http.Server{
		Handler: routerInit,
	}
	server.ListenAndServe()
}
