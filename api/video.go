package api

import (
	"github.com/gin-gonic/gin"
	"tiktok/model"
)

/*
VideoList 声明视频列表格式信息
*/
type VideoList struct {
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
	//config, exists := c.Get("ServerConfig")
	//if !exists {
	//	log.Fatal(errors.New("init config message failed"))
	//}
	//N := config.(*conf.ServerConfig).N
}

// VideoPublish todo:发布视频
func VideoPublish(c *gin.Context) {

}

// GetVideoList todo:获取视频发布列表
func GetVideoList(c *gin.Context) {

}
