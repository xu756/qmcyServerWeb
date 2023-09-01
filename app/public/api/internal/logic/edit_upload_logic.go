package logic

import (
	"context"
	"github.com/xu756/qmcy/app/public/api/internal/svc"
	"github.com/xu756/qmcy/app/public/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type EditUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUploadLogic {
	return &EditUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditUploadLogic) EditUpload(name string) (resp *types.EditUploadRes, err error) {

	return &types.EditUploadRes{
		Url:  "https://qmcy.xu756.top/images/" + name,
		Href: "",
		Alt:  "",
	}, nil
}
