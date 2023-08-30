package controller

import (
	"TikTok/middleware"
	"TikTok/service"
	"TikTok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMessage(c *gin.Context) {
	// 解析请求参数
	token := c.PostForm("token")
	toUserID := c.PostForm("to_user_id")
	actionType := c.PostForm("action_type")
	content := c.PostForm("content")
	information, _ := middleware.CheckToken(token)

	//调用service层处理，并返回响应
	err := service.SendMessageService(information.UserId, toUserID, actionType, content)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		})
	}
}

func ChatHistory(c *gin.Context) {
	// 解析请求参数
	token := c.Query("token")
	toUserID := c.Query("to_user_id")
	information, _ := middleware.CheckToken(token)

	//调用service层处理，并返回响应
	messageList, err := service.ChatHistoryService(information.UserId, toUserID)
	if err != nil {
		c.JSON(http.StatusOK, utils.MessageResponse{
			StatusCode:  1,
			StatusMsg:   err.Error(),
			MessageList: nil,
		})

	} else {
		c.JSON(http.StatusOK, utils.MessageResponse{
			StatusCode:  0,
			StatusMsg:   "聊天记录查找成功",
			MessageList: messageList,
		})
	}

}
