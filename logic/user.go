package logic

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
	"pandownload/dao/mysql"
	"pandownload/dao/redis"
	"pandownload/model"
	email2 "pandownload/utils/email"
)

const (
	// ErrUsername 用户不存在
	ErrUsername = "username error"
	// ErrPassword 密码错误
	ErrPassword    = "password error"
	ErruserExisted = "user existed"
	ErrServicebusy = "service busy"
	ErrauthCode    = "invalid authcode"
)

func CheckuserInfo(userinfo map[string]string) (err error) {
	//获取用户名
	username := userinfo["username"]
	//判断用户是否存在
	password, ok := mysql.CheckuserExist(username)
	if !ok {
		return errors.New(ErrUsername)
	}
	opassword, _ := userinfo["password"]
	if opassword != password {
		return errors.New(ErrPassword)
	}
	return
}

func SignupUser(username, password, email string) error {
	err := mysql.Signup(&model.User{
		Username: username,
		Password: password,
		Email:    email,
	})
	return err
}

func GetuserEmail(username string) string {
	return mysql.GetEmail(username)
}

func SendmailTo(email string, body string) error {
	return email2.SendmailTo([]string{email}, body)
}

func GenerateauthCode(email string) (string, error) {
	num, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	authCode := fmt.Sprintf("%d", num.Int64())[:6]
	err := redis.SaveauthCode(email, authCode)
	return authCode, err
}

func VerifyauthCode(email string, authCode string) error {
	code, err := redis.GetauthCode(email)
	if err != nil {
		return err
	}
	if code != authCode {
		return errors.New(ErrauthCode)
	}
	return nil
}
