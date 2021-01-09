package app

import (
	"golang-side-project/constant"

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

// ResponseHttp data api
func (g Gin) ResponseHttp(httpCode int, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  constant.GetMsg(errCode),
		Data: data,
	})
}
