package model

import (
	"tiktok/db"
	"tiktok/utils"
	"time"
)

type Video struct {
	//Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	//IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL    string `json:"play_url"` // 视频播放地址
	Title      string `json:"title"`    // 视频标题
	CreateDate string `json:"create_date"`
}

type tbVideo struct {
	Video
	UserId int64 `json:"user_id"`
}

type tbFavorite struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
	VideoId  int64  `json:"videoId"`
}

// PushVideo 上传视频至数据库
func PushVideo(playUrl, coverUrl, title string, userId int64) {
	db.Db.Mysql.Create(&tbVideo{
		Video: Video{
			PlayURL:    playUrl,
			CoverURL:   coverUrl,
			Title:      title,
			CreateDate: utils.TransformUnixToDate(time.Now().Unix()),
		},
		UserId: userId,
	})
	AddWorkCount(userId)
}

// GetVideo 从数据库获取视频完整信息  latest_time 表示此次获取的N个视频中发布最晚的时间
func GetVideo(latestTime int64, N int64) []tbVideo {

	videoList := make([]tbVideo, int(N))
	db.Db.Mysql.Where("create_date < ?", utils.TransformUnixToDate(latestTime)).Order("id desc").Limit(int(N)).Find(&videoList)
	return videoList
}

// IsFavorited 判断该用户是否点赞该视频的接口
func IsFavorited(userId, videoId int64) bool {
	tbFavorite := tbFavorite{}
	db.Db.Mysql.Where("user_id = ? AND video_id = ?", userId, videoId).Find(&tbFavorite)
	if tbFavorite.Id == 0 {
		return false
	}
	return true
}

// GetVideoWithUserId 通过用户id获取用户所发布的所有视频
func GetVideoWithUserId(userid int64) []tbVideo {
	var tb []tbVideo
	db.Db.Mysql.Where("user_id = ?", userid).Find(&tb)
	return tb
}
