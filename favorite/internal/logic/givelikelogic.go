package logic

import (
	"context"

	"favorite/favorite"
	"favorite/internal/svc"

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
	// todo: add your logic here and delete this line

	return &favorite.Response{
		Code:    1,
		Message: "点赞成功",
	}, nil
}
