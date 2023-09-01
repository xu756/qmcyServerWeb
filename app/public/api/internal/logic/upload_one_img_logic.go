package logic

import (
	"context"

	"github.com/xu756/qmcy/app/public/api/internal/svc"
	"github.com/xu756/qmcy/app/public/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadOneImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadOneImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadOneImgLogic {
	return &UploadOneImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadOneImgLogic) UploadOneImg(name string) (resp *types.UploadRes, err error) {

	return &types.UploadRes{
		Url:    "https://qmcy.xu756.top/" + name,
		Status: "success",
	}, nil
}
