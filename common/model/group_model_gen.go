// Code generated by goctl. DO NOT EDIT.

package model

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
	groupFieldNames          = builder.RawFieldNames(&Group{}, true)
	groupRows                = strings.Join(groupFieldNames, ",")
	groupRowsExpectAutoSet   = strings.Join(stringx.Remove(groupFieldNames), ",")
	groupRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(groupFieldNames, "id"))

	cachePublicGroupIdPrefix = "cache:public:group:id:"
)

type (
	groupModel interface {
		Insert(ctx context.Context, data *Group) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Group, error)
		Update(ctx context.Context, data *Group) error
		Delete(ctx context.Context, id int64) error
	}

	defaultGroupModel struct {
		sqlc.CachedConn
		table string
	}

	Group struct {
		Id       int64  `db:"id"`        // ID
		ParentId int64  `db:"parent_id"` // 所属父级用户组ID
		Name     string `db:"name"`      // 用户组名称
		Code     string `db:"code"`      // 用户组CODE唯一代码
		Intro    string `db:"intro"`     // 用户组介绍
		Created  int64  `db:"created"`   // 创建时间
		Creator  int64  `db:"creator"`   // 创建人
		Edited   int64  `db:"edited"`    // 修改时间
		Editor   int64  `db:"editor"`    // 修改人
		Deleted  int64  `db:"deleted"`   // 逻辑删除:0=未删除,1=已删除
	}
)

func newGroupModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultGroupModel {
	return &defaultGroupModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."group"`,
	}
}

func (m *defaultGroupModel) withSession(session sqlx.Session) *defaultGroupModel {
	return &defaultGroupModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      `"public"."group"`,
	}
}

func (m *defaultGroupModel) Delete(ctx context.Context, id int64) error {
	publicGroupIdKey := fmt.Sprintf("%s%v", cachePublicGroupIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicGroupIdKey)
	return err
}

func (m *defaultGroupModel) FindOne(ctx context.Context, id int64) (*Group, error) {
	publicGroupIdKey := fmt.Sprintf("%s%v", cachePublicGroupIdPrefix, id)
	var resp Group
	err := m.QueryRowCtx(ctx, &resp, publicGroupIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", groupRows, m.table)
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

func (m *defaultGroupModel) Insert(ctx context.Context, data *Group) (sql.Result, error) {
	publicGroupIdKey := fmt.Sprintf("%s%v", cachePublicGroupIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", m.table, groupRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.ParentId, data.Name, data.Code, data.Intro, data.Created, data.Creator, data.Edited, data.Editor, data.Deleted)
	}, publicGroupIdKey)
	return ret, err
}

func (m *defaultGroupModel) Update(ctx context.Context, data *Group) error {
	publicGroupIdKey := fmt.Sprintf("%s%v", cachePublicGroupIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, groupRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.ParentId, data.Name, data.Code, data.Intro, data.Created, data.Creator, data.Edited, data.Editor, data.Deleted)
	}, publicGroupIdKey)
	return err
}

func (m *defaultGroupModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicGroupIdPrefix, primary)
}

func (m *defaultGroupModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", groupRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGroupModel) tableName() string {
	return m.table
}
