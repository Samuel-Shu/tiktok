package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"tiktok/api"
	"tiktok/db"
)

func InitRouter() {
	r := gin.Default()

	//定义无需鉴权路由组,并初始化配置信息，写进*gin.context里
	NoAuthAPI := r.Group("/douyin", db.InitDb)
	//定义需鉴权路由组
	AuthAPI := r.Group("/douyin")

	//基础接口，抖音基本功能实现：视频流、登录注册、投稿等

	//用户注册
	NoAuthAPI.POST("/user/register/", api.UserRegister)
	//用户登录
	NoAuthAPI.POST("/user/login/", api.UserLogin)
	//视频流 feed
	NoAuthAPI.GET("/feed/", api.Feed)
	//获取用户信息
	AuthAPI.GET("/user/", api.GetUserInfo)
	//用户投稿（发布视频）
	AuthAPI.POST("/publish/action/", api.VideoPublish)
	//用户的视频发布列表，直接列出用户所有投稿过的视频
	AuthAPI.GET("/publish/list/", api.GetVideoList)

	//互动接口：点赞操作、获取喜欢列表、评论等

	//登录用户对视频的点赞和取消点赞操作
	AuthAPI.POST("/favorite/action/")
	//用户的所有点赞视频
	AuthAPI.GET("/favorite/list/")
	//登录用户对视频进行评论
	AuthAPI.POST("/comment/action/")
	//查看视频的所有评论，按发布时间倒序
	AuthAPI.GET("/comment/list/")

	//社交接口：用户间的相互关注、聊天等

	//关注操作
	AuthAPI.POST("/relation/action/")
	//获取关注列表
	AuthAPI.GET("/relation/follow/list/")
	//获取粉丝列表
	AuthAPI.GET("/relation/follower/list/")
	//获取好友列表（相互关注即为好友位）
	AuthAPI.GET("/relation/friend/list/")
	//发送消息（好友间的聊天功能）
	AuthAPI.POST("/message/action/")
	//获取好友间的聊天的记录
	AuthAPI.GET("/message/chat/")

	err := r.Run(":3000")
	if err != nil {
		log.Fatal("http server run failed!")
	}
}
