package logic

import (
	"context"
	"github.com/xu756/qmcy/pb"

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
	contents, err := l.svcCtx.AdminRpc.GetContents(l.ctx, &pb.ContentsReq{
		ContentClass: req.ContentClass,
		Current:      req.Current,
		PageSize:     req.PageSize,
		Title:        req.Title,
		DescText:     req.DescText,
		Path:         req.Path,
	})
	if err != nil {
		return nil, err
	}
	var list = make([]types.Content, 0)
	for _, content := range contents.List {
		list = append(list, types.Content{
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
	return &types.ContentList{
		Total: contents.Total,
		List:  list,
	}, nil
}
