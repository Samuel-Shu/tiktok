package models

import "gorm.io/gorm"

type Favorite struct {
	UserId  uint `json:"userId"`
	VideoId uint `json:"videoId"`
	gorm.Model
}

func (Favorite) TableName() string {
	return "tb_favorite"
}

type DefaultFavoriteModel struct {
	Db *gorm.DB
}

func NewUserModel(db *gorm.DB) *DefaultFavoriteModel {
	return &DefaultFavoriteModel{
		Db: db,
	}
}

func (m *DefaultFavoriteModel) GiveLike(userId, videoId uint64) error {
	favorite := &Favorite{UserId: uint(userId), VideoId: uint(videoId)}
	err := m.Db.Create(favorite).Error
	return err
}

func (m *DefaultFavoriteModel) CancelLike(userId, videoId uint64) error {
	err := m.Db.Unscoped().Where("user_id = ? and video_id = ?", userId, videoId).Delete(&Favorite{}).Error
	return err
}

func (m *DefaultFavoriteModel) GetByUserId(userId uint64) (*[]Favorite, error) {
	favorite := make([]Favorite, 0)
	err := m.Db.Where("user_id = ?", userId).Find(&favorite).Error
	return &favorite, err
}
