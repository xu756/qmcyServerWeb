package logic

import (
	"context"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/public/api/internal/svc"
	"github.com/xu756/qmcy/app/public/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha(req *types.GetCodeReq) (resp *types.GetCodeRes, err error) {
	result, err := l.svcCtx.PublicRpc.GetCode(l.ctx, &pb.GetCodeReq{
		Sign:      req.Sign,
		Timestamp: req.Timestamp,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetCodeRes{
		Expire:    result.Expire,
		Img:       result.Img,
		SessionId: result.SessionId,
	}, nil
}
