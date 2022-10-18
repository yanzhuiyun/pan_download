package redis

import (
	"context"
	"time"
)

const (
	tokenKey = "recent:"
)

func SaveauthCode(emailKey string, authCode string) error {
	err := client.Set(context.Background(), emailKey, authCode, time.Minute*2).Err()
	return err
}

func GetauthCode(email string) (string, error) {
	stmt := client.Get(context.Background(), email)
	return stmt.Result()
}

func SaveuserToken(username string) error {
	return client.SAdd(context.Background(), tokenKey, username).Err()
}

func CheckuserToken(username string) {
	
}
