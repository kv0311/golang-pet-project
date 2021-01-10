package app

import (
	"github.com/gin-gonic/gin"
)

// Gin struc
type Gin struct {
	C *gin.Context
}

// Response struct
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//ResponseHttp data api
func (g Gin) ResponseHttp(httpCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  msg,
		Data: data,
	})
}
