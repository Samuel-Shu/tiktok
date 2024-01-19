package models

import "gorm.io/gorm"

type Relation struct {
	gorm.Model
	FollowerId    uint   `json:"follower_id"`
	FollowingId   uint   `json:"following_id"`
	FollowingName string `json:"following_name"`
}

func (Relation) TableName() string {
	return "tb_relation"
}

type DefaultRelationModel struct {
	Db *gorm.DB
}

func NewCommentModel(db *gorm.DB) *DefaultRelationModel {
	return &DefaultRelationModel{
		Db: db,
	}
}

func (m *DefaultRelationModel) Create(followerId, followingId uint, followingName string) (*Relation, error) {
	relation := &Relation{
		FollowerId:    followerId,
		FollowingId:   followingId,
		FollowingName: followingName,
	}
	err := m.Db.Create(relation).Error
	return relation, err
}

func (m *DefaultRelationModel) DeleteById(followerId, followingId uint) error {
	err := m.Db.Unscoped().Where("follower_id = ? and following_id = ?", followerId, followingId).Delete(&Relation{}).Error
	return err
}

func (m *DefaultRelationModel) GetByFollowerId(followerId uint) ([]Relation, error) {
	relationList := make([]Relation, 0)
	err := m.Db.Where("follower_id = ?", followerId).Find(&relationList).Error
	return relationList, err
}

func (m *DefaultRelationModel) GetByFollowingId(followingId uint) ([]Relation, error) {
	relationList := make([]Relation, 0)
	err := m.Db.Where("following_id = ?", followingId).Find(&relationList).Error
	return relationList, err
}

func (m *DefaultRelationModel) GetFriendList(userId uint) ([]Relation, error) {
	relationList := make([]Relation, 0)
	err := m.Db.Raw(`SELECT A.*
		FROM tb_relation A
		INNER JOIN tb_relation B ON A.follower_id = B.following_id AND A.following_id = B.follower_id
		WHERE A.follower_id = ?`, userId).Scan(&relationList).Error

	return relationList, err
}
