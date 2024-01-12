// Code generated by goctl. DO NOT EDIT.
// Source: message.proto

package server

import (
	"context"

	"message/internal/logic"
	"message/internal/svc"
	"message/message"
)

type MessageServer struct {
	svcCtx *svc.ServiceContext
	message.UnimplementedMessageServer
}

func NewMessageServer(svcCtx *svc.ServiceContext) *MessageServer {
	return &MessageServer{
		svcCtx: svcCtx,
	}
}

func (s *MessageServer) PostMessage(ctx context.Context, in *message.PostMessageRequest) (*message.Response, error) {
	l := logic.NewPostMessageLogic(ctx, s.svcCtx)
	return l.PostMessage(in)
}

func (s *MessageServer) GetMessage(ctx context.Context, in *message.GetMessageRequest) (*message.GetMessageResponse, error) {
	l := logic.NewGetMessageLogic(ctx, s.svcCtx)
	return l.GetMessage(in)
}
