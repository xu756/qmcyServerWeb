package logic

import (
	"context"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/admin/api/internal/svc"
	"github.com/xu756/qmcy/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelContentLogic {
	return &DelContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelContentLogic) DelContent(req *types.ContentReq) (resp *types.Ok, err error) {
	ok, err := l.svcCtx.AdminRpc.DelContent(l.ctx, &pb.ContentReq{
		Id:           req.Id,
		ContentClass: req.ContentClass,
	})
	if err != nil {
		return nil, err
	}
	return &types.Ok{
		Ok: ok.Ok,
	}, nil
}
