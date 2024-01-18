package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"mini-tiktok/service/core/define"
	"mini-tiktok/service/core/internal/svc"
	"mini-tiktok/service/core/internal/types"
	"mini-tiktok/service/favorite/favorite"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedRequest) (resp *types.FeedResponse, err error) {
	resp = new(types.FeedResponse)
	if req.LatestTime == 0 {
		req.LatestTime = uint(time.Now().Unix())
	}
	videos, _ := l.svcCtx.VideoModel.ListByCreatedAt(int64(req.LatestTime), uint(define.N))
	copier.Copy(&resp.VideoList, &videos)

	len := len(videos)
	resp.NextTime = uint64(videos[len-1].CreatedAt.Unix())

	for i, item := range resp.VideoList {
		if req.UserId == 0 {
			resp.VideoList[i].IsFavorite = false
		} else {
			res, err := l.svcCtx.FavoritePb.IsFavorite(l.ctx, &favorite.IsFavoriteRequest{
				UserId:  uint64(req.UserId),
				VideoId: uint64(item.ID),
			})
			if err != nil {
				logx.Error(err)
			}
			resp.VideoList[i].IsFavorite = res.IsFavorite
		}

		res2, err := l.svcCtx.FavoritePb.GetFavoriteCount(l.ctx, &favorite.GetFavoriteCountRequest{VideoId: uint64(item.ID)})
		if err != nil {
			logx.Error(err)
		}
		resp.VideoList[i].FavoriteCount = int64(res2.Count)

		res3, err := l.svcCtx.FavoritePb.GetCommentCount(l.ctx, &favorite.GetCommentCountRequest{VideoId: uint64(item.ID)})
		if err != nil {
			logx.Error(err)
		}
		resp.VideoList[i].CommentCount = int(res3.Count)
	}

	return
}
