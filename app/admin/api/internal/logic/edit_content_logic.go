package logic

import (
	"context"
	"github.com/xu756/qmcy/pb"

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
	ok, err := l.svcCtx.AdminRpc.EditContent(l.ctx, &pb.Content{
		Id:           req.Id,
		Title:        req.Title,
		DescText:     req.DescText,
		ImgUrl:       req.ImgUrl,
		Path:         req.Path,
		Percent:      req.Percent,
		ContentClass: req.ContentClass,
		ContentType:  req.ContentType,
		ContentText:  req.ContentText,
		ContentImg:   req.ContentImg,
		Grade:        req.Created,
		Created:      req.Created,
		Edited:       req.Edited,
		IsEdit:       req.IsEdit,
		Deleted:      req.Deleted,
	})
	if err != nil {
		return nil, err
	}
	return &types.Ok{
		Ok: ok.Ok,
	}, nil
}
