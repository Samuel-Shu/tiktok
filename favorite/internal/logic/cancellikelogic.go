package logic

import (
	"context"

	"favorite/favorite"
	"favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLikeLogic {
	return &CancelLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelLikeLogic) CancelLike(in *favorite.CancelLikeRequest) (*favorite.Response, error) {
	err := l.svcCtx.FavoriteModel.CancelLike(in.UserId, in.VideoId)
	if err != nil {
		return &favorite.Response{
			Code:    1,
			Message: "取消失败",
		}, nil
	}
	return &favorite.Response{
		Code:    0,
		Message: "取消成功",
	}, nil
}
