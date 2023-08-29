package logic

import (
	"context"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/admin/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentLogic {
	return &GetContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContentLogic) GetContent(in *pb.ContentReq) (*pb.Content, error) {
	// todo: add your logic here and delete this line

	return &pb.Content{}, nil
}
