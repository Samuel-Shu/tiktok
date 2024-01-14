package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"mini-tiktok/core/helper"
	"mini-tiktok/core/pb/favorite"

	"mini-tiktok/core/internal/svc"
	"mini-tiktok/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentLogic) GetComment(req *types.GetCommentRequest) (resp *types.GetCommentResponse, err error) {
	resp = new(types.GetCommentResponse)
	fmt.Printf("token:%+v", req)
	result, err := helper.FavoriteClient.GetCommentList(l.ctx, &favorite.GetCommentRequest{
		VideoId: uint64(req.VideoId),
	})
	if err != nil {
		resp.StatusCode = uint(result.Code)
		return
	}
	copier.Copy(&resp.CommentList, &result.CommentList)
	return
}
