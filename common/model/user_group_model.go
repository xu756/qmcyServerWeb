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

var _ UserGroupModel = (*customUserGroupModel)(nil)
var cacheUserGroupPrefix = "cache:public:userGroup:id:"

type (
	// UserGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserGroupModel.
	UserGroupModel interface {
		userGroupModel
		FindUserGroup(ctx context.Context, userId int64) (int64, error)
	}

	customUserGroupModel struct {
		*defaultUserGroupModel
	}
)

// NewUserGroupModel returns a model for the database table.
func NewUserGroupModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserGroupModel {
	return &customUserGroupModel{
		defaultUserGroupModel: newUserGroupModel(conn, c, opts...),
	}
}

// FindUserGroup 查询用户组
func (m *customUserGroupModel) FindUserGroup(ctx context.Context, userId int64) (int64, error) {
	userGroupsKey := fmt.Sprintf("%s%v", cacheUserGroupPrefix, userId)
	var resp int64
	err := m.QueryRowCtx(ctx, &resp, userGroupsKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select id from %s where user_id = $1 limit 1", m.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
	})
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return 0, ErrNotFound
	default:
		return 0, xerr.NewDbErr("用户组查询失败", err)
	}
}
