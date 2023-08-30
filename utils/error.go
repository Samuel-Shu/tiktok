package utils

import "errors"

var (
	ErrorUserNameNull    = errors.New("用户名为空")
	ErrorUserNameExtend  = errors.New("用户名长度不符合规范")
	ErrorPasswordNull    = errors.New("密码为空")
	ErrorPasswordLength  = errors.New("密码长度不符合规范")
	ErrorUserExit        = errors.New("用户已存在")
	ErrorFullPossibility = errors.New("用户不存在，账号或密码出错")
	ErrorNullPointer     = errors.New("空指针异常")
	ErrorPasswordFalse   = errors.New("密码错误")
	ErrorRelationExit    = errors.New("关注已存在")
	ErrorRelationNull    = errors.New("关注不存在")
	ErrorUserID          = errors.New("用户ID错误")
	ErrorUser            = errors.New("无法关注自己")
	ErrorMessageUser     = errors.New("无法向自己发送消息")
	ErrorFollow          = errors.New("关注失败")
	ErrorActionType      = errors.New("action_type错误")
	ErrorFollowList      = errors.New("获取关注列表失败")
	ErrorSendMessage     = errors.New("发送消息失败")
	ErrorChatHistory     = errors.New("查找聊天记录失败")
)
