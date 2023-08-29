// Code generated by goctl. DO NOT EDIT.

package miniModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{}, true)
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "id"), ",")
	usersRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(usersFieldNames, "id"))

	cachePublicUsersIdPrefix = "cache:public:users:id:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id        int64        `db:"id"`
		Username  string       `db:"username"`
		Mobile    string       `db:"mobile"`
		Password  string       `db:"password"`
		Avatar    string       `db:"avatar"`
		Level     int64        `db:"level"`
		Openid    string       `db:"openid"`
		Unionid   string       `db:"unionid"`
		CreatedAt sql.NullTime `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
		Created   int64        `db:"created"`
		Edited    int64        `db:"edited"`
		IsEdit    int64        `db:"is_edit"`
		Deleted   int64        `db:"deleted"`
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."users"`,
	}
}

func (m *defaultUsersModel) withSession(session sqlx.Session) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      `"public"."users"`,
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicUsersIdKey)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, publicUsersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Username, data.Mobile, data.Password, data.Avatar, data.Level, data.Openid, data.Unionid, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Created, data.Edited, data.IsEdit, data.Deleted)
	}, publicUsersIdKey)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, data *Users) error {
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.Username, data.Mobile, data.Password, data.Avatar, data.Level, data.Openid, data.Unionid, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Created, data.Edited, data.IsEdit, data.Deleted)
	}, publicUsersIdKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
