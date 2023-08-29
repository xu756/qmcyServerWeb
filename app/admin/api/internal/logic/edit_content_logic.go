package logic

import (
	"context"

	"github.com/xu756/qmcy/app/admin/api/internal/svc"
	"github.com/xu756/qmcy/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditContentLogic {
	return &EditContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditContentLogic) EditContent(req *types.Content) (resp *types.Ok, err error) {
	// todo: add your logic here and delete this line

	return
}
