package logic

import (
	"context"
	"github.com/xu756/qmcy/app/admin/api/internal/svc"
	"github.com/xu756/qmcy/app/admin/api/internal/types"
	"github.com/xu756/qmcy/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupsLogic {
	return &GetGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupsLogic) GetGroups(req *types.GetGroupReq) (resp *types.GroupList, err error) {
	grouplist, err := l.svcCtx.AdminRpc.GetGroups(l.ctx, &pb.GetGroupReq{
		Id:       req.Id,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var list = make([]types.Group, 0)

	for _, v := range grouplist.List {
		list = append(list, *convertGroup(v))
	}
	return &types.GroupList{
		Total: grouplist.Total,
		List:  list,
	}, nil
}

func convertGroup(g *pb.Group) *types.Group {
	var c = &types.Group{
		Name:    g.Name,
		Code:    g.Code,
		Intro:   g.Intro,
		Created: g.Created,
		Creator: types.User{
			Id:         g.Creator.Id,
			State:      g.Creator.State,
			Name:       g.Creator.Name,
			HeadImgUrl: g.Creator.HeadImgUrl,
			Mobile:     g.Creator.Mobile,
			Deleted:    g.Creator.Deleted,
		},
		Edited: g.Edited,
		Editor: types.User{
			Id:         g.Editor.Id,
			State:      g.Editor.State,
			Name:       g.Editor.Name,
			HeadImgUrl: g.Editor.HeadImgUrl,
			Mobile:     g.Editor.Mobile,
			Deleted:    g.Editor.Deleted,
		},
		Level: g.Level,
	}
	for _, v := range g.Children {
		c.Children = append(c.Children, *convertGroup(v))
	}
	return c
}
