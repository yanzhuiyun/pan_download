package mysql

import (
	"fmt"
	"pandownload/model"
)

const (
	tableName = "users"
)

func CheckuserExist(username string) (password string, ok bool) {
	sqlStr := fmt.Sprintf("SELECT password from %s WHERE username=?", tableName)
	stmt, err := db.Prepare(sqlStr)
	result := stmt.QueryRow(username)
	if err != nil {
		return
	}
	result.Scan(&password)
	if password == "" {
		return
	}
	ok = true
	return
}

func Signup(user *model.User) error {
	sqlStr := fmt.Sprintf("INSERT INTO %s(username,password,email) VALUES(?,?,?)", tableName)
	stmt, _ := db.Prepare(sqlStr)
	_, err := stmt.Exec(user.Username, user.Password, user.Email)
	return err
}

func GetEmail(username string) (email string) {
	sqlStr := fmt.Sprintf("SELECT email from %s WHERE username=?", tableName)
	stmt, _ := db.Prepare(sqlStr)
	result := stmt.QueryRow(username)
	if err != nil {
		return ""
	}
	result.Scan(&email)
	return email
}
