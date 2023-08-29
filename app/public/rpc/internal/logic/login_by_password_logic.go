package logic

import (
	"context"
	"github.com/xu756/qmcy/app/public/rpc/internal/svc"
	"github.com/xu756/qmcy/common/model"
	"github.com/xu756/qmcy/common/tool"
	"github.com/xu756/qmcy/common/xerr"
	"github.com/xu756/qmcy/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPasswordLogic {
	return &LoginByPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginByPassword 通过密码登录
func (l *LoginByPasswordLogic) LoginByPassword(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var resp = new(pb.LoginResponse)
	result := l.svcCtx.Captcha.Verify(in.SessionId, in.Code, true)
	if !result {
		return resp, xerr.NewMsgError("验证码错误")
	}
	user, err := l.svcCtx.UserModel.LoginByPassword(l.ctx, in.Username)
	if err != nil {
		return resp, err
	}
	if user.Password != in.Password {
		return resp, xerr.NewMsgError("密码错误")
	}
	roles, err := l.svcCtx.UserRoleModel.FindUserRoles(l.ctx, user.Id)
	if err != nil {
		return resp, err
	}
	userGroupId, err := l.svcCtx.UserGroupModel.FindUserGroup(l.ctx, user.Id)
	if err != nil {
		return resp, err
	}
	code, err := l.svcCtx.GroupModel.FindGroupCode(l.ctx, userGroupId)
	if err != nil {
		return resp, err
	}
	jwt, err := l.svcCtx.Jwt.NewJwt(user.Id, roles, code)
	if err != nil {
		return resp, err
	}
	_, err = l.svcCtx.AccountModel.LoginInsert(l.ctx, &model.Account{
		UserId:    user.Id,
		OpenCode:  "密码登录",
		Category:  1,
		Created:   tool.TimeNowInTimeZoneUnix(),
		CreateRpc: "public.rpc",
		Edited:    tool.TimeNowInTimeZoneUnix(),
		Editor:    0,
		Deleted:   0,
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token:  jwt,
		Expire: 7200,
	}, nil
}
