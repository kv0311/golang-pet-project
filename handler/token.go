package handler

import (
	"encoding/json"
	"golang-side-project/common/app"
	"golang-side-project/constant"
	"golang-side-project/model"
	"golang-side-project/repo"

	"github.com/gin-gonic/gin"
)

// CreateAccount by user id and password
func CreateAccount(c *gin.Context) {
	responseData := app.Gin{C: c}
	var userInfo interface{}
	var user model.User
	err := c.BindJSON(&userInfo)
	if err != nil {
		return
	}

	byteData, _ := json.Marshal(userInfo)
	err = json.Unmarshal(byteData, &user)
	err = repo.CreateAccount(user)
	if err != nil {
		responseData.ResponseHttp(constant.ERROR, constant.ERROR, err.Error())
		return
	}
	responseData.ResponseHttp(constant.SUCCESS, constant.SUCCESS, "successful")
}

// Login func
func LoginByUser(c *gin.Context) {
	responseData := app.Gin{C: c}
	var userInfo interface{}
	var user model.LoginModel
	err := c.BindJSON(&userInfo)
	if err != nil {
		return
	}

	byteData, _ := json.Marshal(userInfo)
	err = json.Unmarshal(byteData, &user)
	token, err := repo.LoginByUser(user)
	if err != nil {
		responseData.ResponseHttp(constant.ERROR, constant.ERROR, err.Error())
		return
	}
	if token == "" {
		responseData.ResponseHttp(constant.ERROR, constant.ERROR, "User or Password is not correct")
		return
	}
	responseData.ResponseHttp(constant.SUCCESS, constant.SUCCESS, token)
	return
}
