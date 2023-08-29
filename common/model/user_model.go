package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/xu756/qmcy/common/tool"
	"github.com/xu756/qmcy/common/xerr"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)
var cachePublicUserMobildPrefix = "cache:public:user:mobile:"

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		LoginByPassword(ctx context.Context, username string) (*User, error)
		FindOneUser(ctx context.Context, id int64) (*User, error)
		LoginByWx(ctx context.Context, mobile string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

// LoginByPassword 登录
func (m *defaultUserModel) LoginByPassword(ctx context.Context, username string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where name = $1 or mobile = $2 ", userRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, username, username)
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, xerr.NewMsgError("用户不存在,请注册")
	default:
		return nil, xerr.NewDbErr("查询失败", err)
	}
}

// FindOneUser 查询用户
func (m *defaultUserModel) FindOneUser(ctx context.Context, id int64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cachePublicUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, xerr.NewDbErr("查询失败", err)
	}
}

// LoginByWx 登录
func (m *defaultUserModel) LoginByWx(ctx context.Context, mobile string) (*User, error) {
wxLog:
	publicUserMobileKey := fmt.Sprintf("%s%v", cachePublicUserMobildPrefix, mobile)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, publicUserMobileKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where mobile = $1 limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, mobile)
	})
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		var user = User{
			State:      0,
			Name:       "微信用户" + mobile,
			HeadImgUrl: "https://qmcy.xu756.top/img/avatar.jpeg",
			Mobile:     mobile,
			Salt:       "",
			Password:   "",
			Created:    tool.TimeNowInTimeZoneUnix(),
			Creator:    0,
			Edited:     tool.TimeNowInTimeZoneUnix(),
			Editor:     0,
			Deleted:    0,
		}
		_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
			query := fmt.Sprintf("insert into %s (state,name,head_img_url,mobile,salt,password,created,creator,edited,editor,deleted) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", m.table)
			return conn.ExecCtx(ctx, query, user.State, user.Name, user.HeadImgUrl, user.Mobile, user.Salt, user.Password, user.Created, user.Creator, user.Edited, user.Editor, user.Deleted)
		}, publicUserMobileKey)
		if err != nil {
			return nil, xerr.NewDbErr("创建用户失败", err)
		}
		goto wxLog
	default:
		return nil, xerr.NewDbErr("查询失败", err)
	}
}
