package model

import "time"

//User management model
type User struct {
	UserID    string    `json:"user_id" pg:"user_id"`
	UserName  string    `json:"user_name" pg:"user_name"`
	Password  string    `json:"password" pg:"password" `
	Email     string    `json:"email" pg:"email"`
	CreatedOn time.Time `pg:"created_on"`
	LastLogin time.Time `pg:"last_login"`
	Token     string    `pg:"token"`
}

// LoginModel struc
type LoginModel struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
