package logic

import (
	"context"
	"github.com/xu756/qmcy/common/miniModel"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/admin/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddContentLogic {
	return &AddContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddContentLogic) AddContent(in *pb.Content) (*pb.Ok, error) {
	result, err := l.svcCtx.MiniContentModel.AddContent(l.ctx, &miniModel.Contents{
		Id:           in.Id,
		Title:        in.Title,
		DescText:     in.DescText,
		ImgUrl:       in.ImgUrl,
		Path:         in.Path,
		Percent:      in.Percent,
		ContentClass: in.ContentClass,
		ContentType:  in.ContentType,
		ContentText:  in.ContentText,
		ContentImg:   in.ContentImg,
		Grade:        in.Created,
		Created:      in.Created,
		Edited:       in.Edited,
		IsEdit:       in.IsEdit,
		Deleted:      in.Deleted,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Ok{}, nil
}
