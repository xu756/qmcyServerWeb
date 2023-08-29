package logic

import (
	"context"
	"github.com/xu756/qmcy/common/xctx"
	"github.com/xu756/qmcy/common/xerr"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/user/api/internal/svc"
	"github.com/xu756/qmcy/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserInfo, err error) {
	value, ok := xctx.FromContextForJwt(l.ctx)
	if !ok {
		return nil, xerr.LogOut()
	}
	info, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &pb.UserInfoRequest{
		Id: value.ID,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfo{
		Id:        info.Id,
		Name:      info.Name,
		Avatar:    info.Avatar,
		Role:      info.Role,
		GroupCode: info.GroupCode,
	}, nil
}
