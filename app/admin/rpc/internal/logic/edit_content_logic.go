package logic

import (
	"context"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/admin/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type EditContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditContentLogic {
	return &EditContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditContentLogic) EditContent(in *pb.Content) (*pb.Ok, error) {
	// todo: add your logic here and delete this line

	return &pb.Ok{}, nil
}
