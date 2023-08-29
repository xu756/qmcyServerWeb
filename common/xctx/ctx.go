package xctx

import (
	"context"
	"github.com/xu756/qmcy/common/xjwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type key string

var userKey key

func NewContextForJwt(ctx context.Context, u *xjwt.AuthInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContextForJwt(ctx context.Context) (*xjwt.AuthInfo, bool) {
	if u := ctx.Value(userKey); u != nil {
		if value, ok := u.(*xjwt.AuthInfo); ok {
			return value, true
		}
	}
	logx.WithContext(ctx).Error("没有找到用户信息")
	return nil, false
}
