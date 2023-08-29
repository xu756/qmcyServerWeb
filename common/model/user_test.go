package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"log"
	"testing"
)

// LoginByPassword 登录
func TestLoginByPassword(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	UserModel := NewUserModel(conn, sqlCache)
	user, err := UserModel.LoginByPassword(context.Background(), "admin")
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(user)
}

// FindOneUser 查询用户
func TestFindOneUser(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	UserModel := NewUserModel(conn, sqlCache)
	user, err := UserModel.FindOneUser(context.Background(), 1)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(user)
}
