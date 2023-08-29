package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/xu756/qmcy/common/xerr"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AccountModel = (*customAccountModel)(nil)

type (
	// AccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountModel.
	AccountModel interface {
		accountModel
		LoginInsert(ctx context.Context, data *Account) (sql.Result, error)
	}

	customAccountModel struct {
		*defaultAccountModel
	}
)

// NewAccountModel returns a model for the database table.
func NewAccountModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AccountModel {
	return &customAccountModel{
		defaultAccountModel: newAccountModel(conn, c, opts...),
	}
}

// LoginInsert 插入账号信息
func (m *defaultAccountModel) LoginInsert(ctx context.Context, data *Account) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)", m.table, accountRowsExpectAutoSet)
	cacheCtx, err := m.ExecNoCacheCtx(ctx, query, data.Id, data.UserId, data.OpenCode, data.Category, data.Created, data.CreateRpc, data.Edited, data.Editor, data.Deleted)
	if err != nil {
		return nil, xerr.NewDbErr("插入账号信息失败", err)
	}
	return cacheCtx, nil
}
