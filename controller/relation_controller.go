package controller

import (
	"TikTok/middleware"
	"TikTok/model"
	"TikTok/service"
	"TikTok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Follow 关注、取关操作
func Follow(c *gin.Context) {
	// 解析请求参数
	token := c.PostForm("token")
	toUserID := c.PostForm("to_user_id")

	actionType := c.PostForm("action_type")
	information, _ := middleware.CheckToken(token)

	//调用service层处理，并返回响应
	err := service.FollowService(information.UserId, toUserID, actionType)
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

// FollowList 关注列表
func FollowList(c *gin.Context) {
	// 解析请求参数
	toUserID := c.Query("user_id")

	//调用service层处理，并返回响应
	followList, err := service.FollowListService(toUserID)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1,
			StatusMsg:  utils.ErrorFollowList.Error(),
		})

	} else {
		c.JSON(http.StatusOK, utils.UserResponse{
			StatusCode: 0,
			StatusMsg:  "获取关注列表成功",
			UserList:   followList,
		})
	}

}

// FollowerList 粉丝列表
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, model.User{
		ID:              2,
		Name:            "王景鹏",
		FollowCount:     3,
		FollowerCount:   10,
		BackgroundImage: "王八蛋",
		Signature:       "111",
		TotalFavorited:  123,
		WorkCount:       213,
		FavoriteCount:   23,
	})

}

// FriendList 好友列表
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, model.User{
		ID:              2,
		Name:            "王景鹏",
		FollowCount:     3,
		FollowerCount:   10,
		BackgroundImage: "王八蛋",
		Signature:       "111",
		TotalFavorited:  123,
		WorkCount:       213,
		FavoriteCount:   23,
	})

}
