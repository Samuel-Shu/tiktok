package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"mini-tiktok/core/helper"
	"mini-tiktok/core/pb/follow"

	"mini-tiktok/core/internal/svc"
	"mini-tiktok/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingLogic {
	return &GetFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowingLogic) GetFollowing(req *types.GetFollowingRequest) (resp *types.GetFollowingResponse, err error) {
	resp = new(types.GetFollowingResponse)
	result, err := helper.FollowClient.GetFollowingList(l.ctx, &follow.GetFollowingListRequest{
		UserId: uint64(req.UserId),
	})
	logx.Error(err)
	if err != nil {
		resp.StatusCode = 1
	}
	copier.Copy(&resp.UserList, result.UserList)

	return
}
