package logic

import (
	"context"
	"fmt"
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

func (l *VideoListLogic) VideoList() (resp *types.VideoListResponse, err error) {
	resp = new(types.VideoListResponse)
	list, err := l.svcCtx.VideoModel.List()

	copier.Copy(&resp.VideoList, &list)

	for _, video := range resp.VideoList {
		fmt.Printf("%+v\n", video)
	}

	return
}
