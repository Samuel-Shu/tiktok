package logic

import (
	"context"
	"log"
	"mini-tiktok/service/core/helper"
	"mini-tiktok/service/core/internal/svc"
	"mini-tiktok/service/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishRequest, userId int64) (resp *types.PublishResponse, err error) {
	resp = new(types.PublishResponse)
	ffmpeg, err := helper.Ffmpeg(req.PlayURL, 1)
	if err != nil {
		log.Fatal(err)
	}
	picUrl, err := helper.FileUploadToJpg(&ffmpeg)
	if err != nil {
		log.Fatal(err)
	}

	err = l.svcCtx.VideoModel.Create(userId, req.PlayURL, picUrl, req.Title)
	if err != nil {
		return nil, err
	}
	return
}
