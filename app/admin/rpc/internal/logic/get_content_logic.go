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
	content, err := l.svcCtx.MiniContentModel.FindContent(l.ctx, in.Id)
	if err != nil {
		return new(pb.Content), err
	}
	return &pb.Content{
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
