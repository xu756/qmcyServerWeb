package logic

import (
	"context"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/admin/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelContentLogic {
	return &DelContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelContentLogic) DelContent(in *pb.ContentReq) (*pb.Ok, error) {
	// todo: add your logic here and delete this line

	return &pb.Ok{}, nil
}
