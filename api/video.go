package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	config2 "tiktok/config"
	"tiktok/model"
	"tiktok/utils"
	"time"
)

/*
VideoList 声明视频列表格式信息
*/
type VideoList struct {
	IsFavorite bool
	model.Video
	Author model.User `json:"author"`
}

/*
VideoMes 定义视频流返回信息
*/
type VideoMes struct {
	model.HttpStatus
	NextTime  int64
	VideoList []VideoList `json:"video_list"`
}

// Feed todo:获取视频流信息（播放视频）
func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	Date, err := strconv.Atoi(latestTime)
	if err != nil {
		log.Fatal(err)
	}
	/*
		获取单次视频推流量 N ：每一次从数据库中获取N个视频进行推流播放
	*/
	config, exists := c.Get("ServerConfig")
	if !exists {
		log.Fatal(errors.New("init config message failed"))
	}
	N := config.(*config2.ServerConfig).N

	fmt.Println(Date)
	if Date == 0 {
		Date = int(time.Now().Unix())
	}
	Video := model.GetVideo(int64(Date), int64(N))
	VideoList := make([]VideoList, len(Video))
	for i := 0; i < len(Video); i++ {
		VideoList[i].Title = Video[i].Title
		VideoList[i].ID = Video[i].ID
		VideoList[i].PlayURL = Video[i].PlayURL
		VideoList[i].CommentCount = Video[i].CommentCount
		VideoList[i].CoverURL = Video[i].CoverURL
		VideoList[i].FavoriteCount = Video[i].FavoriteCount
		VideoList[i].Author = model.GetUserData(Video[i].UserId)
		VideoList[i].IsFavorite = model.IsFavorited(Video[i].UserId, Video[i].ID)
	}
	c.JSON(http.StatusOK, VideoMes{
		HttpStatus: model.HttpStatus{
			StatusCode: 0,
			StatusMsg:  "获取视频成功",
		},

		NextTime:  utils.TransformDateToUnix(Video[0].CreateDate),
		VideoList: VideoList,
	})
}

// VideoPublish todo:发布视频
func VideoPublish(c *gin.Context) {
	title := c.PostForm("title")
	o := utils.NewOSSConfig(c)
	videoStream, err := c.FormFile("data")
	if err != nil {
		log.Println("① push video is failed", err)
	}
	videoFile, err1 := videoStream.Open()
	if err1 != nil {
		log.Println("② push video is failed", err)
	}
	defer func(videoFile multipart.File) {
		err := videoFile.Close()
		if err != nil && err != io.EOF {
			log.Println(err)
		}
	}(videoFile)
	if err != nil {
		log.Println("③ push video is failed", err)
	}

	file, err := ioutil.ReadAll(videoFile)
	if err != nil && err == io.EOF {
		fmt.Println(err)
	}

	// videoYes : 视频是否上传成功
	videoYes := o.PushVideo(fmt.Sprintf("%s.mp4", title), file)
	playUrl := o.GetVideo(fmt.Sprintf("%s.mp4", title))
	//fmt.Println("one teo ", playUrl)
	ffmpeg, err := utils.Ffmpeg(playUrl, 1)
	if err != nil {
		log.Println("ffmpeg failed")
	}
	//coverYes : 封面是否上传成功
	coverYes := o.PushVideoCover(fmt.Sprintf("%s.jpg", title), ffmpeg)

	coverUrl := o.GetCover(fmt.Sprintf("%s.jpg", title))

	userId, exists := c.Get("userId")
	if !exists {
		log.Println("the userId is not exist")
	}

	if !videoYes || !coverYes {
		c.JSON(http.StatusOK, model.HttpStatus{
			StatusCode: -1,
			StatusMsg:  "上传视频失败",
		})
		return
	}

	model.PushVideo(playUrl, coverUrl, title, userId.(int64))
	c.JSON(http.StatusOK, model.HttpStatus{
		StatusCode: 0,
		StatusMsg:  "上传视频成功",
	})

}

// GetVideoList todo:获取视频发布列表
func GetVideoList(c *gin.Context) {
	userId := c.Query("user_id")
	atoi, err := strconv.Atoi(userId)
	if err != nil {
		log.Fatal(err)
	}

	Video := model.GetVideoWithUserId(int64(atoi))
	VideoList := make([]VideoList, len(Video))
	for i := 0; i < len(Video); i++ {
		VideoList[i].Title = Video[i].Title
		VideoList[i].ID = Video[i].ID
		VideoList[i].PlayURL = Video[i].PlayURL
		VideoList[i].CommentCount = Video[i].CommentCount
		VideoList[i].CoverURL = Video[i].CoverURL
		VideoList[i].FavoriteCount = Video[i].FavoriteCount
		VideoList[i].Author = model.GetUserData(Video[i].UserId)
		VideoList[i].IsFavorite = model.IsFavorited(Video[i].UserId, Video[i].ID)
	}
	model.UpdateWorkCount(int64(atoi), int64(len(Video)))
	c.JSON(http.StatusOK, VideoMes{
		HttpStatus: model.HttpStatus{
			StatusCode: 0,
			StatusMsg:  "获取投稿列表成功",
		},

		VideoList: VideoList,
	})
}
