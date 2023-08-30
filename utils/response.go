package utils

import "TikTok/model"

type Response struct {
	StatusCode  int             `json:"status_code"`
	StatusMsg   string          `json:"status_msg"`
	UserList    []model.User    `json:"user_list,omitempty"`
	MessageList []model.Message `json:"message_list,omitempty"`
}

type UserResponse struct {
	StatusCode int          `json:"status_code"`
	StatusMsg  string       `json:"status_msg"`
	UserList   []model.User `json:"user_list"`
}

type MessageResponse struct {
	StatusCode  int             `json:"status_code"`
	StatusMsg   string          `json:"status_msg"`
	MessageList []model.Message `json:"message_list"`
}
