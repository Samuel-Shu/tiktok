package logic

import (
	"context"

	"favorite/favorite"
	"favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteCountLogic {
	return &GetFavoriteCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFavoriteCountLogic) GetFavoriteCount(in *favorite.GetFavoriteCountRequest) (*favorite.GetFavoriteCountResponse, error) {
	count, err := l.svcCtx.FavoriteModel.CountByVideoId(uint(in.VideoId))

	return &favorite.GetFavoriteCountResponse{
		Count: uint64(count),
	}, err
}
