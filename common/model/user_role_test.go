package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"log"
	"testing"
)

// FindUserGroup 查询用户组

func TestFindUserRoles(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	UserRoleModel := NewUserRoleModel(conn, sqlCache)
	roles, err := UserRoleModel.FindUserRoles(context.Background(), 1)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(roles)
}
