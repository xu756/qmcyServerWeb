package logic

import (
	"context"
	"github.com/xu756/qmcy/app/admin/rpc/internal/svc"
	"github.com/xu756/qmcy/common/model"
	"github.com/xu756/qmcy/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupsLogic {
	return &GetGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupsLogic) GetGroups(in *pb.GetGroupReq) (*pb.GroupList, error) {
	var resp = &pb.GroupList{}
	groups, err := l.svcCtx.GroupModel.FindGroups(l.ctx, in.Id, in.PageNum, in.PageSize)
	if err != nil {
		return resp, err
	}
	resp.Total = int64(len(groups))
	for _, v := range groups {
		resp.List = append(resp.List, convertGroup(v))
	}
	return resp, nil
}

func convertGroup(g *model.SqlGroup) *pb.Group {
	var c = &pb.Group{
		Name:    g.Name,
		Code:    g.Code,
		Intro:   g.Intro,
		Created: g.Created,
		Creator: &pb.GroupUserInfo{
			Id:         g.Creator.Id,
			State:      g.Creator.State,
			Name:       g.Creator.Name,
			HeadImgUrl: g.Creator.HeadImgUrl,
			Mobile:     g.Creator.Mobile,
			Deleted:    g.Creator.Deleted,
		},
		Edited: g.Edited,
		Editor: &pb.GroupUserInfo{
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
		c.Children = append(c.Children, convertGroup(&v))
	}
	return c
}
