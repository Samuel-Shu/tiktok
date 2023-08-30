package router

import (
	"TikTok/controller"
	"TikTok/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//主路由组
	mainGroup := r.Group("/douyin")
	{
		// relation路由组
		relationGroup := mainGroup.Group("/relation")
		{
			relationGroup.POST("/action/", middleware.JwtMiddleware(), controller.Follow)
			relationGroup.GET("/follow/list/", middleware.JwtMiddleware(), controller.FollowList)
			relationGroup.GET("/follower/list/", middleware.JwtMiddleware(), controller.FollowerList)
			relationGroup.GET("/friend/list/", middleware.JwtMiddleware(), controller.FriendList)
		}
		messageGroup := mainGroup.Group("/message")
		{
			messageGroup.GET("/chat/", middleware.JwtMiddleware(), controller.ChatHistory)
			messageGroup.POST("/action/", middleware.JwtMiddleware(), controller.SendMessage)

		}
	}

	return r
}
