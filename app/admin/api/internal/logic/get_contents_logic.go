package logic

import (
	"context"

	"github.com/xu756/qmcy/app/admin/api/internal/svc"
	"github.com/xu756/qmcy/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentsLogic {
	return &GetContentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContentsLogic) GetContents(req *types.ContentsReq) (resp *types.ContentList, err error) {
	// todo: add your logic here and delete this line

	return
}
