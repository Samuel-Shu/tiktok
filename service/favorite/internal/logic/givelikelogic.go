package logic

import (
	"context"

	"mini-tiktok/service/favorite/favorite"
	"mini-tiktok/service/favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GiveLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGiveLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiveLikeLogic {
	return &GiveLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GiveLikeLogic) GiveLike(in *favorite.GiveLikeRequest) (*favorite.Response, error) {
	logx.Info(in)
	logx.Info("userId:", in.UserId)
	err := l.svcCtx.FavoriteModel.GiveLike(in.UserId, in.VideoId)
	if err != nil {
		return &favorite.Response{
			Code:    1,
			Message: "点赞失败",
		}, nil
	}
	return &favorite.Response{
		Code:    0,
		Message: "点赞成功",
	}, nil
}
