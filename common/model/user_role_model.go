package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/xu756/qmcy/common/xerr"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserRoleModel = (*customUserRoleModel)(nil)
var cacheUserRolesPrefix = "cache:public:userRoles:id:"

type (
	// UserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRoleModel.
	UserRoleModel interface {
		userRoleModel
		FindUserRoles(ctx context.Context, userId int64) ([]int64, error)
	}

	customUserRoleModel struct {
		*defaultUserRoleModel
	}
)

// NewUserRoleModel returns a model for the database table.
func NewUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserRoleModel {
	return &customUserRoleModel{
		defaultUserRoleModel: newUserRoleModel(conn, c, opts...),
	}
}

// FindUserRoles 查询用户权限
func (m *customUserRoleModel) FindUserRoles(ctx context.Context, userId int64) ([]int64, error) {
	userRolesKey := fmt.Sprintf("%s%v", cacheUserRolesPrefix, userId)
	var resp []int64
	err := m.QueryRowCtx(ctx, &resp, userRolesKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select id from %s where user_id = $1 limit 1", m.table)
		return conn.QueryRowsCtx(ctx, v, query, userId)
	})
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return []int64{}, ErrNotFound
	default:
		return nil, xerr.NewDbErr("权限查询失败", err)
	}

}
