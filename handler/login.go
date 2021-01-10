package handler

import (
	"encoding/json"
	"golang-side-project/common/app"
	"golang-side-project/constant"
	"golang-side-project/model"
	"golang-side-project/repo"
	util "golang-side-project/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var ResponseMap map[string]interface{}

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
	user.CreatedOn = time.Now()
	user.LastLogin = time.Now()
	isExistUserID, err := checkIsUserIDExists(user.UserID)
	if isExistUserID == true {
		util.Debug(map[string]interface{}{}, "user id exists")
		ResponseHTTP(constant.SUCCESS, "user id exists", map[string]interface{}{})
	}
	err = repo.CreateAccount(user)
	if err != nil {
		responseData.ResponseHttp(constant.ERROR, err.Error(), "")
		return
	}
	responseData.ResponseHttp(constant.SUCCESS, "successful to create account", "")
}

// LoginByUser func
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
		responseData.ResponseHttp(constant.ERROR, err.Error(), "")
		return
	}
	if token == "" {
		responseData.ResponseHttp(constant.ERROR, "User or Password is not correct", "")
		return
	}
	responseData.ResponseHttp(constant.SUCCESS, "successful to create account", token)
	return
}

func checkIsUserIDExists(UserID string) (bool, error) {
	isUserExists, err := repo.CheckIsUserIDExists(UserID)
	if err != nil {
		ResponseHTTP(constant.ERROR, err.Error(), "")
		return false, err
	}
	return isUserExists, nil
}
