package miniModel

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/xu756/qmcy/common/xerr"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContentsModel = (*customContentsModel)(nil)

type (
	// ContentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContentsModel.
	ContentsModel interface {
		contentsModel
		FindContentsByContentClass(ctx context.Context, ContentClass string, PageNum int64, PageSize int64) ([]*Contents, error)
		AddContent(ctx context.Context, data *Contents) (sql.Result, error)
		EditContent(ctx context.Context, data *Contents) (sql.Result, error)
		FindContent(ctx context.Context, id int64, ContentClass string) (*Contents, error)
		DelContent(ctx context.Context, id int64, ContentClass string) (sql.Result, error)
	}

	customContentsModel struct {
		*defaultContentsModel
	}
)

// NewContentsModel returns a model for the database table.
func NewContentsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ContentsModel {
	return &customContentsModel{
		defaultContentsModel: newContentsModel(conn, c, opts...),
	}
}

// FindContentsByContentClass 根据内容分类查找内容
func (m *defaultContentsModel) FindContentsByContentClass(ctx context.Context, ContentClass string, PageNum int64, PageSize int64) ([]*Contents, error) {
	var res []*Contents
	query := fmt.Sprintf("select %s from %s where content_class = $1 limit $2 offset $3", contentsRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &res, query, ContentClass, PageSize, (PageNum-1)*PageSize)
	switch {
	case err == nil:
		return res, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, xerr.NewDbErr("数据库查询失败", err)
	}
}

// AddContent 添加内容
func (m *defaultContentsModel) AddContent(ctx context.Context, data *Contents) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)", m.table, contentsRowsExpectAutoSet)
	return m.ExecNoCacheCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Title, data.DescText, data.ImgUrl, data.Path, data.Percent, data.ContentClass, data.ContentType, data.ContentText, data.ContentImg, data.Grade, data.Created, data.Edited, data.IsEdit, data.Deleted)
}

// EditContent 编辑内容
func (m *defaultContentsModel) EditContent(ctx context.Context, data *Contents) (sql.Result, error) {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, contentsRowsWithPlaceHolder)
	return m.ExecNoCacheCtx(ctx, query, data.Id, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Title, data.DescText, data.ImgUrl, data.Path, data.Percent, data.ContentClass, data.ContentType, data.ContentText, data.ContentImg, data.Grade, data.Created, data.Edited, data.IsEdit, data.Deleted)
}

// FindContent 查找单个内容
func (m *defaultContentsModel) FindContent(ctx context.Context, id int64, ContentClass string) (*Contents, error) {
	var res Contents
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1 and  content_class = $2", contentsRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &res, query, id, contentsRows)
	switch {
	case err == nil:
		return &res, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// DelContent
func (m *defaultContentsModel) DelContent(ctx context.Context, id int64, ContentClass string) (sql.Result, error) {
	query := fmt.Sprintf("update %s set deleted = 1 where id = $1 and  content_class = $2", m.table)
	return m.ExecNoCacheCtx(ctx, query, id, ContentClass)

}
