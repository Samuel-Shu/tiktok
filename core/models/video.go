package models

import "gorm.io/gorm"

type Video struct {
	ID            int64  `json:"id"` // 视频唯一标识
	Author        User   `json:"author" gorm:"foreignKey:UserId"`
	UserId        int64  `json:"user_id"`
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	//IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL    string `json:"play_url"` // 视频播放地址
	Title      string `json:"title"`    // 视频标题
	CreateDate string `json:"create_date"`
}

func (Video) TableName() string {
	return "tb_video"
}

type DefaultVideoModel struct {
	Db *gorm.DB
}

func NewVideoModel(db *gorm.DB) *DefaultVideoModel {
	return &DefaultVideoModel{
		Db: db,
	}
}

func (m *DefaultVideoModel) List() ([]Video, error) {
	videoList := make([]Video, 0)
	err := m.Db.Model(&Video{}).Preload("Author").Find(&videoList).Error
	return videoList, err
}

func (m *DefaultVideoModel) Create(userId int64, playURL string, coverURL string) error {
	video := &Video{
		UserId:   userId,
		PlayURL:  playURL,
		CoverURL: coverURL,
	}
	err := m.Db.Create(video).Error
	return err
}
