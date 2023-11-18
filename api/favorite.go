package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tiktok/model"
)

type FavoriteVideoList struct {
	model.HttpStatus
	VideoList []VideoList `json:"video_list"`
}

// GiveLikeOrCancelLike 点赞或取消点赞接口
func GiveLikeOrCancelLike(c *gin.Context) {
	videoId, actionType := c.Query("video_id"), c.Query("action_type")

	userId, exists := c.Get("userId")
	if !exists {
		log.Fatal("从token中获取userId失败！")
	}

	videoid, err := strconv.Atoi(videoId)
	if err != nil {
		log.Fatal(err)
	}

	actiontype, err := strconv.Atoi(actionType)
	if err != nil {
		log.Fatal(err)
	}

	if actiontype != 1 && actiontype != 2 {
		c.JSON(http.StatusOK, model.HttpStatus{
			StatusCode: -1,
			StatusMsg:  "点赞操作错误！",
		})

		return
	}
	//actontype == 1 表示点赞
	if actiontype == 1 {
		model.GiveLike(userId.(int64), int64(videoid))
		c.JSON(http.StatusOK, model.HttpStatus{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		})
		return
	}

	model.CancelLike(userId.(int64), int64(videoid))
	c.JSON(http.StatusOK, model.HttpStatus{
		StatusCode: 0,
		StatusMsg:  "取消点赞成功",
	})

}

// GetFavoriteList 获取用户喜欢列表
func GetFavoriteList(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Fatal(err)
	}

	//用户id错误或不存在
	if !model.FindUserWithId(int64(userId)) {
		c.JSON(http.StatusOK, FavoriteVideoList{
			HttpStatus: model.HttpStatus{
				StatusCode: -1,
				StatusMsg:  "用户id错误或不存在",
			},
		})
		return
	}

	videoIdList := model.FindFavoriteVideoId(int64(userId))
	VideoList := make([]VideoList, len(videoIdList))
	for i := 0; i < len(videoIdList); i++ {
		Video := model.GetVideoWithVideoId(videoIdList[i])
		VideoList[i].Title = Video.Title
		VideoList[i].ID = Video.ID
		VideoList[i].PlayURL = Video.PlayURL
		VideoList[i].CommentCount = Video.CommentCount
		VideoList[i].CoverURL = Video.CoverURL
		VideoList[i].FavoriteCount = Video.FavoriteCount
		VideoList[i].Author = model.GetUserData(Video.UserId)
		VideoList[i].IsFavorite = true
	}
	fmt.Println(VideoList)
	c.JSON(http.StatusOK, FavoriteVideoList{
		HttpStatus: model.HttpStatus{
			StatusCode: 0,
			StatusMsg:  "获取喜欢列表成功",
		},
		VideoList: VideoList,
	})

}
