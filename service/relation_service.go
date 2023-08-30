package service

import (
	"TikTok/db"
	"TikTok/model"
	"TikTok/utils"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func FollowService(userID uint64, toUserID string, actionType string) (err error) {
	toID, err := strconv.ParseUint(toUserID, 10, 64)
	if err != nil {
		return utils.ErrorUserID
	}

	if userID == toID {
		return utils.ErrorUser
	}
	if !utils.CheckToUserID(toID) {
		return utils.ErrorUserID
	}

	if actionType != "1" && actionType != "2" {
		return utils.ErrorActionType
	}

	if actionType == "1" {
		tx := db.DB.Begin()
		err := tx.Table("tb_user").Model(&model.User{}).
			Where("id = ?", userID).
			Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error

		if err != nil {
			tx.Rollback()
			return utils.ErrorFollow
		}
		err = tx.Table("tb_user").Model(&model.User{}).
			Where("id = ?", toID).
			Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error

		if err != nil {
			tx.Rollback()
			return utils.ErrorFollow
		}

		err = tx.Table("tb_relation").Create(&model.Relation{
			FollowerID:  userID,
			FollowingID: toID,
			IsDeleted:   0,
			CreateTime:  time.Now(),
		}).Error

		if err != nil {
			tx.Rollback()
			return utils.ErrorFollow
		}

		tx.Commit()
	}
	if actionType == "2" {
		tx := db.DB.Begin()
		err := tx.Table("tb_user").Model(&model.User{}).
			Where("id = ?", userID).
			Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error

		if err != nil {
			tx.Rollback()
			return utils.ErrorFollow
		}
		err = tx.Table("tb_user").Model(&model.User{}).
			Where("id = ?", toID).
			Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error

		if err != nil {
			tx.Rollback()
			return utils.ErrorFollow
		}

		err = tx.Table("tb_relation").Model(&model.Relation{}).
			Where("follower_id = ? AND following_id = ?", userID, toID).
			Update("isdeleted", 1).Error

		if err != nil {
			tx.Rollback()
			return utils.ErrorFollow
		}

		tx.Commit()

	}

	return nil
}

func FollowListService(userID string) ([]model.User, error) {
	ID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, utils.ErrorUserID
	}
	if !utils.CheckToUserID(ID) {
		return nil, utils.ErrorUserID
	}
	var followList []model.User

	// 使用联接查询获取关注列表
	result := db.DB.Table("tb_relation").
		Select("tb_user.*").
		Joins("JOIN tb_user ON tb_relation.following_id = tb_user.id").
		Where("tb_relation.follower_id = ?", userID).
		Find(&followList)
	if result.Error != nil {
		return nil, utils.ErrorFollowList
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	for i := 0; i < len(followList); i++ {
		var relation model.Relation
		db.DB.Table("tb_relation").Limit(1).
			Where("follower_id=? AND following_id=? AND isdeleted = ?", followList[i].ID, ID, 0).
			Find(&relation)
		fmt.Println(relation)
		followList[i].IsFollow = relation.ID != 0
	}
	return followList, nil
}

func FollowingListService(userID string) ([]model.User, error) {
	ID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, utils.ErrorUserID
	}
	if !utils.CheckToUserID(ID) {
		return nil, utils.ErrorUserID
	}
	var followingList []model.User

	// 使用联接查询获取关注列表
	result := db.DB.Table("tb_relation").
		Select("tb_user.*").
		Joins("JOIN tb_user ON tb_relation.follower_id = tb_user.id").
		Where("tb_relation.following_id = ?", userID).
		Find(&followingList)
	if result.Error != nil {
		return nil, utils.ErrorFollowList
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	for i := 0; i < len(followingList); i++ {
		var relation model.Relation
		db.DB.Table("tb_relation").Limit(1).
			Where("follower_id=? AND following_id=? AND isdeleted = ?", ID, followingList[i].ID, 0).
			Find(&relation)
		followingList[i].IsFollow = relation.ID != 0
	}
	return followingList, nil
}

func FriendListService(userID string) ([]model.User, error) {
	ID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, utils.ErrorUserID
	}
	if !utils.CheckToUserID(ID) {
		return nil, utils.ErrorUserID
	}
	var followingList []model.User

	// 使用联接查询获取关注列表
	result := db.DB.Table("tb_relation r1").
		Select("tb_user.*").
		Joins("JOIN tb_relation r2 ON r1.follower_id = r2.following_id"+
			" AND r1.following_id = r2.follower_id").
		Joins("JOIN tb_user ON r2.follower_id = tb_user.id").
		Where("r2.following_id = ?", userID).
		Scan(&followingList)
	if result.Error != nil {
		return nil, utils.ErrorFollowList
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	for i := 0; i < len(followingList); i++ {
		followingList[i].IsFollow = true
	}
	return followingList, nil
}
