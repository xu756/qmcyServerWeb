package miniModel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"log"
	"testing"
)

// 查找内容
func TestFindContents(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	ContentsModel := NewContentsModel(conn, sqlCache)

	res, err := ContentsModel.FindContentsByContentClass(context.Background(), "banner", 1, 1)
	if err != nil {
		log.Print("err:", err)
	}

	fmt.Print(res)

}

// 添加内容
func TestAddContent(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	ContentsModel := NewContentsModel(conn, sqlCache)

	res, err := ContentsModel.AddContent(context.Background(), &Contents{
		Title:        "测试",
		ContentClass: "content",
	})
	if err != nil {
		log.Print("err:", err)
	}
	fmt.Print(res.RowsAffected())

}

// 编辑内容
func TestEditContent(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	ContentsModel := NewContentsModel(conn, sqlCache)

	res, err := ContentsModel.EditContent(context.Background(), &Contents{
		Id:           21,
		Title:        "测试qqq",
		ContentClass: "edit",
	})
	if err != nil {
		log.Print("err:", err)
	}
	fmt.Print(res.RowsAffected())

}

// 查找一个
func TestFindContent(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	ContentsModel := NewContentsModel(conn, sqlCache)

	res, err := ContentsModel.FindContent(context.Background(), 1)
	if err != nil {
		log.Print("err:", err)
	}
	fmt.Print(res)

}

// 删除
func TestDelContent(t *testing.T) {
	conn := postgres.New(DbSource)
	sqlCache := RedisConf
	ContentsModel := NewContentsModel(conn, sqlCache)
	res, err := ContentsModel.DelContent(context.Background(), 2)
	if err != nil {
		log.Print("err:", err)
	}
	fmt.Print(res.RowsAffected())
}
