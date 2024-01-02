package logic

import (
	"context"
	"mini-tiktok/core/define"
	"mini-tiktok/core/helper"
	"mini-tiktok/core/internal/svc"
	"mini-tiktok/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserRegisterResponse)
	l.Logger.Info(req.Username)
	l.Logger.Info(req.Password)

	user, err := l.svcCtx.UserModel.GetByName(req.Username)
	if err == nil {
		resp.StatusMsg = "此用户名已注册"
		resp.StatusCode = 1
		return
	}

	err = l.svcCtx.UserModel.Create(req.Username, helper.Md5(req.Password))
	if err != nil {
		resp.StatusMsg = "注册失败"
		resp.StatusCode = 1
		return
	}

	resp.UserId = 1
	resp.Token, err = helper.GenerateToken(user.ID, user.Name, define.TokenExpire)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "生成token失败"
		return
	}
	resp.StatusCode = 0
	return
}
