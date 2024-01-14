package logic

import (
	"context"
	"mini-tiktok/core/helper"
	"mini-tiktok/core/pb/favorite"

	"mini-tiktok/core/internal/svc"
	"mini-tiktok/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.LikeRequest) (resp *types.BaseResponse, err error) {
	resp = new(types.BaseResponse)
	var response *favorite.Response
	logx.Info(req)
	if req.ActionType == 1 {
		response, err = helper.FavoriteClient.GiveLike(l.ctx, &favorite.GiveLikeRequest{
			UserId:  uint64(req.UserId),
			VideoId: uint64(req.VideoId),
		})
	} else {
		response, err = helper.FavoriteClient.CancelLike(l.ctx, &favorite.CancelLikeRequest{
			UserId:  uint64(req.UserId),
			VideoId: uint64(req.VideoId),
		})
	}

	if err != nil {
		return
	}
	resp.StatusMsg = response.Message
	resp.StatusCode = int(response.Code)
	return
}
