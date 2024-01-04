package logic

import (
	"context"
	"mini-tiktok/core/internal/svc"
	"mini-tiktok/core/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type VideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoListLogic) VideoList(userId uint) (resp *types.VideoListResponse, err error) {
	resp = new(types.VideoListResponse)
	list, err := l.svcCtx.VideoModel.ListByUserId(userId)
	copier.Copy(&resp.VideoList, &list)
	user, err := l.svcCtx.UserModel.GetById(userId)
	if err != nil {
		return
	}
	for _, video := range resp.VideoList {
		video.Author.ID = int(user.ID)
		video.Author.Name = user.Name
	}

	return
}
