package models

import (
	"time"
)

type User struct {
	BaseModel
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func UserRegister(username string, email string, pwd string) (ok bool) {
	queryString := "INSERT INTO user(username, password, email, created_on, modified_on) VALUES (?, ?, ?, ?, ?)"
	_, err := DB.Exec(queryString, username, pwd, email, time.Now(), time.Now())
	if err != nil {
		return false
	}
	return true
}

func FindUser(e string, pwd string) (bool, *User) {
	queryString := "SELECT * FROM users WHERE email = ?"
	row := DB.QueryRow(queryString, e)
	var username, email, password, createdOn, modifiedOn string
	var id int
	err := row.Scan(&id, &username, &password, &email, &createdOn, &modifiedOn)
	if err != nil || password != pwd {
		return false, &User{}
	}
	u := &User{
		BaseModel: BaseModel{
			ID: id,
			CreateOn: createdOn,
			ModifiedOn: modifiedOn,
		},
		Username: username,
		Email: email,
	}
	return true, u
}