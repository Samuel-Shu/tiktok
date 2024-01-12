package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"message/internal/svc"
	"message/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogic {
	return &GetMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageLogic) GetMessage(in *message.GetMessageRequest) (*message.GetMessageResponse, error) {
	resp := new(message.GetMessageResponse)
	list, err := l.svcCtx.MessageModel.GetByToUserId(uint(in.ToUserId))
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		return resp, err
	}
	err = copier.Copy(&resp.MessageList, &list)

	for i, item := range list {
		resp.MessageList[i].CreatedAt = item.CreatedAt.Format("2006-01-02 15-04-05")
	}

	if err != nil {
		panic(err)
	}
	return resp, nil
}
