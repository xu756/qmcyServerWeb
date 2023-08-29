package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xu756/qmcy/common/xerr"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupModel = (*customGroupModel)(nil)
var cacheGroupCode = "cache:public:groupCode:id:"

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
		FindGroupCode(ctx context.Context, id int64) (string, error)
		FindGroups(ctx context.Context, parentId, pageNum, pageSize int64) ([]*SqlGroup, error)
	}

	customGroupModel struct {
		*defaultGroupModel
	}
)

// NewGroupModel returns a model for the database table.
func NewGroupModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupModel {
	return &customGroupModel{
		defaultGroupModel: newGroupModel(conn, c, opts...),
	}
}

// FindGroupCode 根据用户组ID查询用户组唯一code
func (m *customGroupModel) FindGroupCode(ctx context.Context, id int64) (string, error) {
	userGroupCodeKey := fmt.Sprintf("%s%v", cacheGroupCode, id)
	var resp string
	err := m.QueryRowCtx(ctx, &resp, userGroupCodeKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select code from %s where id = $1 limit 1", m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return "", ErrNotFound
	default:
		return "", xerr.NewDbErr("用户组查询失败", err)
	}
}

// FindGroups 根据父级ID查询用户组列表
// pageNum是请求的第几页，pageSize是这一页有多少个
func (m *customGroupModel) FindGroups(ctx context.Context, parentId, pageNum, pageSize int64) ([]*SqlGroup, error) {
	result := make([]*sqlResult, 0)
	query := "SELECT jsonb_build_object ('id',parent.ID,'code',parent.code,'name',parent.NAME,'intro',parent.intro,'created',parent.created,'creator',GetUserInfoByID (parent.creator),'edited',parent.edited,'editor',GetUserInfoByID (parent.editor),'level',1,'children',(\nSELECT COALESCE (jsonb_agg (jsonb_build_object ('id',child.ID,'code',child.code,'name',child.NAME,'intro',child.intro,'created',child.created,'creator',GetUserInfoByID (parent.creator),'edited',parent.edited,'editor',GetUserInfoByID (parent.editor),'level',2,'children',(\nSELECT COALESCE (jsonb_agg (jsonb_build_object ('id',subchild.ID,'code',subchild.code,'name',subchild.NAME,'intro',subchild.intro,'created',subchild.created,'creator',GetUserInfoByID (parent.creator),'edited',parent.edited,'editor',GetUserInfoByID (parent.editor),'level',3)),'[]' :: JSONB) FROM \"group\" AS subchild WHERE subchild.parent_id=child.ID AND subchild.deleted=0))),'[]' :: JSONB) FROM \"group\" AS child WHERE child.parent_id=parent.ID AND child.deleted=0)) AS RESULT FROM \"group\" AS parent LEFT JOIN \"group\" AS child ON child.parent_id=parent.ID AND child.deleted=0 WHERE parent.parent_id=$1 AND parent.deleted=0 GROUP BY parent.ID LIMIT $2 OFFSET $3;"
	err := m.QueryRowsNoCacheCtx(ctx, &result, query, parentId, pageSize, (pageNum-1)*pageSize)
	switch {
	case err == nil:
		groups := make([]*SqlGroup, 0)
		for _, item := range result {
			group := &SqlGroup{}
			err := json.Unmarshal(item.Result, group)
			if err != nil {
				return nil, xerr.NewDbErr("序列化用户组查询失败", err)
			}
			groups = append(groups, group)
		}
		return groups, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, xerr.NewDbErr("用户组查询失败", err)
	}
}
