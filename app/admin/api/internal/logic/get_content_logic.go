package logic

import (
	"context"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/admin/api/internal/svc"
	"github.com/xu756/qmcy/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentLogic {
	return &GetContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContentLogic) GetContent(req *types.ContentReq) (resp *types.Content, err error) {
	content, err := l.svcCtx.AdminRpc.GetContent(l.ctx, &pb.ContentReq{
		Id:           req.Id,
		ContentClass: req.ContentClass,
	})
	if err != nil {
		return nil, err
	}
	return &types.Content{
		Id:           content.Id,
		Title:        content.Title,
		DescText:     content.DescText,
		ImgUrl:       content.ImgUrl,
		Path:         content.Path,
		Percent:      content.Percent,
		ContentClass: content.ContentClass,
		ContentType:  content.ContentType,
		ContentText:  content.ContentText,
		ContentImg:   content.ContentImg,
		Grade:        content.Created,
		Created:      content.Created,
		Edited:       content.Edited,
		IsEdit:       content.IsEdit,
		Deleted:      content.Deleted,
	}, nil
}
