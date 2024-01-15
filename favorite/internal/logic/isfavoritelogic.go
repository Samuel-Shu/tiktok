package logic

import (
	"context"

	"favorite/favorite"
	"favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFavoriteLogic {
	return &IsFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFavoriteLogic) IsFavorite(in *favorite.IsFavoriteRequest) (*favorite.IsFavoriteResponse, error) {
	res, err := l.svcCtx.FavoriteModel.IsFavor(uint(in.VideoId), uint(in.UserId))

	return &favorite.IsFavoriteResponse{
		IsFavorite: res,
	}, err
}
