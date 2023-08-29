package logic

import (
	"context"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/admin/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentsLogic {
	return &GetContentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContentsLogic) GetContents(in *pb.ContentsReq) (*pb.ContentList, error) {
	var resp = &pb.ContentList{}
	contents, err := l.svcCtx.MiniContentModel.FindContentsByContentClass(l.ctx, in.ContentClass, in.PageNum, in.PageSize)
	if err != nil {
		return resp, err
	}
	resp.Total = int64(len(contents))
	for _, content := range contents {
		resp.List = append(resp.List, &pb.Content{
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
		})
	}
	return resp, nil
}
