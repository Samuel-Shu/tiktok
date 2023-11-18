package model

import "tiktok/db"

// GiveLike 视频点赞操作
func GiveLike(userid, videoid int64) {
	//1、将tb_user 表的favorite_count进行加一
	AddFavoriteCountInTbUser(userid)
	//2、将关联信息插入tb_favorite表中
	InsertFavoriteMes(userid, videoid)
	//3、将tb_video 表的favorite_count 进行加一
	AddFavoriteCountInTbVideo(videoid)
}

// CancelLike 视频取消点赞操作
func CancelLike(userid, videoid int64) {
	//1、将tb_user 表的favorite_count 进行减一
	DeleteFavoriteCountInTbUser(userid)
	//2、把对应的信息从tb_favorite表中硬删除(连同数据库记录一起删除)
	DeleteFavoriteMes(userid, videoid)
	//3、将tb_video 表的favorite_count 进行加一
	DeleteFavoriteCountInTbVideo(videoid)
}

// InsertFavoriteMes 将关联信息插入tb_favorite表中
func InsertFavoriteMes(userid, videoid int64) {
	mes := tbFavorite{
		Username: GetUserData(userid).Name,
		UserId:   userid,
		VideoId:  videoid,
	}
	db.Db.Mysql.Create(&mes)
}

// DeleteFavoriteMes 把对应的信息从tb_favorite表中硬删除
func DeleteFavoriteMes(userid, videoid int64) {
	db.Db.Mysql.Unscoped().Where("user_id = ? AND video_id = ?", userid, videoid).Delete(&tbFavorite{})
}

func FindFavoriteVideoId(userId int64) []int64 {
	var videoIdList []int64
	db.Db.Mysql.Model(&tbFavorite{}).Where("user_id = ?", userId).Select("video_id").Scan(&videoIdList)
	return videoIdList
}
