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
	contentsFieldNames          = builder.RawFieldNames(&Contents{}, true)
	contentsRows                = strings.Join(contentsFieldNames, ",")
	contentsRowsExpectAutoSet   = strings.Join(stringx.Remove(contentsFieldNames, "id"), ",")
	contentsRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(contentsFieldNames, "id"))

	cachePublicContentsIdPrefix = "cache:public:contents:id:"
)

type (
	contentsModel interface {
		Insert(ctx context.Context, data *Contents) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Contents, error)
		Update(ctx context.Context, data *Contents) error
		Delete(ctx context.Context, id int64) error
	}

	defaultContentsModel struct {
		sqlc.CachedConn
		table string
	}

	Contents struct {
		Id           int64        `db:"id"`
		CreatedAt    sql.NullTime `db:"created_at"`
		UpdatedAt    sql.NullTime `db:"updated_at"`
		DeletedAt    sql.NullTime `db:"deleted_at"`
		Title        string       `db:"title"`
		DescText     string       `db:"desc_text"`
		ImgUrl       string       `db:"img_url"`
		Path         string       `db:"path"`
		Percent      int64        `db:"percent"`
		ContentClass string       `db:"content_class"`
		ContentType  int64        `db:"content_type"`
		ContentText  string       `db:"content_text"`
		ContentImg   string       `db:"content_img"`
		Grade        int64        `db:"grade"`
		Created      int64        `db:"created"`
		Edited       int64        `db:"edited"`
		IsEdit       int64        `db:"is_edit"`
		Deleted      int64        `db:"deleted"`
	}
)

func newContentsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultContentsModel {
	return &defaultContentsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."contents"`,
	}
}

func (m *defaultContentsModel) withSession(session sqlx.Session) *defaultContentsModel {
	return &defaultContentsModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      `"public"."contents"`,
	}
}

func (m *defaultContentsModel) Delete(ctx context.Context, id int64) error {
	publicContentsIdKey := fmt.Sprintf("%s%v", cachePublicContentsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicContentsIdKey)
	return err
}

func (m *defaultContentsModel) FindOne(ctx context.Context, id int64) (*Contents, error) {
	publicContentsIdKey := fmt.Sprintf("%s%v", cachePublicContentsIdPrefix, id)
	var resp Contents
	err := m.QueryRowCtx(ctx, &resp, publicContentsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", contentsRows, m.table)
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

func (m *defaultContentsModel) Insert(ctx context.Context, data *Contents) (sql.Result, error) {
	publicContentsIdKey := fmt.Sprintf("%s%v", cachePublicContentsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)", m.table, contentsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Title, data.DescText, data.ImgUrl, data.Path, data.Percent, data.ContentClass, data.ContentType, data.ContentText, data.ContentImg, data.Grade, data.Created, data.Edited, data.IsEdit, data.Deleted)
	}, publicContentsIdKey)
	return ret, err
}

func (m *defaultContentsModel) Update(ctx context.Context, data *Contents) error {
	publicContentsIdKey := fmt.Sprintf("%s%v", cachePublicContentsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, contentsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Title, data.DescText, data.ImgUrl, data.Path, data.Percent, data.ContentClass, data.ContentType, data.ContentText, data.ContentImg, data.Grade, data.Created, data.Edited, data.IsEdit, data.Deleted)
	}, publicContentsIdKey)
	return err
}

func (m *defaultContentsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicContentsIdPrefix, primary)
}

func (m *defaultContentsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", contentsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultContentsModel) tableName() string {
	return m.table
}
