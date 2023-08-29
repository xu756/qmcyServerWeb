package logic

import (
	"context"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/user/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.UserInfoRequest) (*pb.UserInfo, error) {
	var resp = &pb.UserInfo{}
	user, err := l.svcCtx.UserModel.FindOneUser(l.ctx, in.Id)
	if err != nil {
		return resp, err
	}
	resp.Id = user.Id
	resp.Name = user.Name
	resp.Avatar = user.HeadImgUrl
	roles, err := l.svcCtx.UserRoleModel.FindUserRoles(l.ctx, in.Id)
	if err != nil {
		return resp, err
	}
	resp.Role = roles
	userGroupId, err := l.svcCtx.UserGroupModel.FindUserGroup(l.ctx, in.Id)
	if err != nil {
		return resp, err
	}
	code, err := l.svcCtx.GroupModel.FindGroupCode(l.ctx, userGroupId)
	if err != nil {
		return resp, err
	}
	resp.GroupCode = code
	return resp, nil
}
