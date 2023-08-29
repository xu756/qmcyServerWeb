package miniModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContentsModel = (*customContentsModel)(nil)

type (
	// ContentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContentsModel.
	ContentsModel interface {
		contentsModel
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
