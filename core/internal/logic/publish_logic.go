package logic

import (
	"context"
	"log"
	"mini-tiktok/core/helper"
	"mini-tiktok/core/internal/svc"
	"mini-tiktok/core/internal/types"

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

func (l *PublishLogic) Publish(req *types.PublishRequest) (resp *types.PublishResponse, err error) {
	resp = new(types.PublishResponse)
	log.Println("url:", req.PlayURL)
	ffmpeg, err := helper.Ffmpeg(req.PlayURL, 1)
	if err != nil {
		log.Fatal(err)
	}
	picUrl, err := helper.FileUploadToJpg(&ffmpeg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("jpg url:", picUrl)

	return
}
