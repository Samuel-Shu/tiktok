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

type GetFansLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFansLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFansLogic {
	return &GetFansLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFansLogic) GetFans(req *types.GetFansRequest) (resp *types.GetFansResponse, err error) {
	resp = new(types.GetFansResponse)
	result, err := helper.FollowClient.GetFans(l.ctx, &follow.GetFansRequest{
		UserId: uint64(req.UserId),
	})
	logx.Error(err)
	if err != nil {
		resp.StatusCode = 1
	}
	copier.Copy(&resp.UserList, result.UserList)

	return
}
