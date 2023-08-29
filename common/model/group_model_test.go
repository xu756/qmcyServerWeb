package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"log"
	"testing"
)

// FindGroupCode 根据用户组ID查询用户组唯一code
func TestFindGroupCode(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	GroupModel := NewGroupModel(conn, sqlCache)

	code, err := GroupModel.FindGroupCode(context.Background(), 1)
	if err != nil {
		log.Print("err:", err)
	}

	fmt.Print(code)

}

// FindGroups 根据父级ID查询用户组列表
// pageNum是请求的第几页，pageSize是这一页有多少个
func TestFindGroups(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	GroupModel := NewGroupModel(conn, sqlCache)

	groups, err := GroupModel.FindGroups(context.TODO(), 0, 1, 10)
	if err != nil {
		log.Print("err:", err)
	}

	for _, group := range groups {
		fmt.Print(group)
	}
}
