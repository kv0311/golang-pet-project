package repo

import (
	"golang-side-project/db"
	"golang-side-project/model"
	util "golang-side-project/utils"

	"time"

	_ "github.com/lib/pq"
)

//CreateAccount create
func CreateAccount(user model.User) error {
	user.CreatedOn = time.Now()
	user.LastLogin = time.Now()

	connection := db.GetConnection()
	_, err := connection.Query("insert into user_management values($1,$2,$3,$4,$5,$6)", user.UserID, user.UserName, user.Password, user.Email, user.CreatedOn, user.LastLogin)
	if err != nil {
		return err
	}
	return nil
}

// LoginByUser func
func LoginByUser(userLogin model.LoginModel) (string, error) {
	isCorrect, err := checkUserPasswordCorrect(userLogin.UserID, userLogin.Password)
	if err != nil {
		return "", err
	}

	if isCorrect == false {
		return "", nil
	}

	token, err := util.GenerateToken(userLogin.UserID, userLogin.Password)
	if err != nil {
		util.Error(map[string]interface{}{}, "fail to generate token")
		return token, err
	}
	return token, nil
	// Inser db
}

//CheckIsUserIDExists check user id exist or not
func CheckIsUserIDExists(UserID string) (bool, error) {
	var user model.User
	connection := db.GetConnection()
	rows, err := connection.Query("select * from user_management where user_id = $1", UserID)
	if err != nil {
		util.Error(map[string]interface{}{}, "fail to query check exists")
	}
	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.UserName, &user.Password, &user.Email, &user.CreatedOn, &user.LastLogin, &user.Token)
	}
	if err != nil {
		util.Error(map[string]interface{}{}, "fail to parse check exists")
	}
	if user.UserID == "" {
		return false, nil
	} else {
		return true, nil
	}
}

func checkUserPasswordCorrect(userID string, password string) (bool, error) {
	connection := db.GetConnection()
	rows, err := connection.Query("select count(*) from user_management where user_id = $1 and password = $2", userID, password)
	if err != nil {
		return false, err
	}
	countRow := checkCount(rows)
	if countRow > 0 {
		return true, nil
	}
	return false, nil
}
