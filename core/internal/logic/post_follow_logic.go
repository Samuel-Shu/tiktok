package logic

import (
	"context"
	"mini-tiktok/core/helper"
	"mini-tiktok/core/internal/svc"
	"mini-tiktok/core/internal/types"
	"mini-tiktok/core/pb/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostFollowLogic {
	return &PostFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostFollowLogic) PostFollow(req *types.PostFollowRequest) (resp *types.BaseResponse, err error) {
	resp = new(types.BaseResponse)
	user, err := l.svcCtx.UserModel.GetById(req.ToUserId)
	if err != nil {
		return
	}

	result, err := helper.FollowClient.PostFollow(context.Background(), &follow.PostFollowRequest{
		UserId:     uint64(req.UserId),
		ToUserId:   uint64(req.ToUserId),
		Username:   user.Name,
		ActionType: uint64(req.ActionType),
	})
	resp.StatusMsg = result.Message
	resp.StatusCode = int(result.Code)
	return
}
