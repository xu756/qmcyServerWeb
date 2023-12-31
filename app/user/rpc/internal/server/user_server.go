// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/user/rpc/internal/logic"
	"github.com/xu756/qmcy/app/user/rpc/internal/svc"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfo, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}
