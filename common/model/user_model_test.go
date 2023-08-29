package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"log"
	"testing"
)

// LoginByWx 登录
func TestLoginByWx(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	UserModel := NewUserModel(conn, sqlCache)
	user, err := UserModel.LoginByWx(context.Background(), "1524992333")
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("userId:  ", user.Id)
}
