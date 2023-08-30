package controller

import (
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SendMessage(c *gin.Context) {
	c.JSON(http.StatusOK, "关注成功")
}

func ChatHistory(c *gin.Context) {
	c.JSON(http.StatusOK, model.Message{
		ID:         1,
		ToUserID:   332,
		FromUserID: 213,
		Content:    "你妈喊你回家吃饭",
		CreateTime: time.Now(),
	})
}
