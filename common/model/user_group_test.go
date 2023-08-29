package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"log"
	"testing"
)

// FindUserGroup 查询用户组

func TestFindUserGroup(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	UserGroupModel := NewUserGroupModel(conn, sqlCache)
	id, err := UserGroupModel.FindUserGroup(context.Background(), 1)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(id)
}
