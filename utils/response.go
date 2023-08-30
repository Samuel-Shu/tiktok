package utils

import "TikTok/model"

type Response struct {
	StatusCode  int             `json:"status_code"`
	StatusMsg   string          `json:"status_msg"`
	UserList    []model.User    `json:"user_list,omitempty"`
	MessageList []model.Message `json:"message_list,omitempty"`
}
