package logic

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"pandownload/dao/redis"
	"time"
)

const (
	tokenSecret  = "pan-download"
	InvalidToken = "invalid token"
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

func getSecret(token *jwt.Token) (interface{}, error) {
	return []byte(tokenSecret), nil
}

func GenerateToken(username string) (string, error) {
	timeExpire := time.Now().Add(5 * time.Hour).Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: timeExpire,
			Issuer:    "pan-download",
			Subject:   "fileProcess",
		},
	}).SignedString([]byte(tokenSecret))
	err = redis.SaveuserToken(username)
	return token, err
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, getSecret)
	if err != nil {
		return nil, err
	}
	if calims, ok := token.Claims.(*Claims); ok && token.Valid {
		return calims, nil
	}
	return nil, errors.New(InvalidToken)
}

func GetUsername(token string) string {
	c, err := ParseToken(token)
	if err != nil {
		return ""
	}
	return c.Username
}
