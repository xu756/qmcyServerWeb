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
	var resp = new(pb.Ok)
	result, err := l.svcCtx.MiniContentModel.DelContent(l.ctx, in.Id, in.ContentClass)
	if err != nil {
		return resp, err
	}
	ok, err := result.RowsAffected()
	if err != nil {
		return resp, err
	}
	resp.Ok = ok == 1
	return resp, nil
}
