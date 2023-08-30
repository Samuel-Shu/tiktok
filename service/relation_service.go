package service

import (
	"TikTok/db"
	"TikTok/model"
	"TikTok/utils"
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
			return utils.ErrorFollw
		}
		err = tx.Table("tb_user").Model(&model.User{}).
			Where("id = ?", toID).
			Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error

		if err != nil {
			tx.Rollback()
			return utils.ErrorFollw
		}

		err = tx.Table("tb_relation").Create(&model.Relation{
			FollowerID:  userID,
			FollowingID: toID,
			IsDeleted:   0,
			CreateTime:  time.Now(),
		}).Error

		if err != nil {
			tx.Rollback()
			return utils.ErrorFollw
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
			return utils.ErrorFollw
		}
		err = tx.Table("tb_user").Model(&model.User{}).
			Where("id = ?", toID).
			Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error

		if err != nil {
			tx.Rollback()
			return utils.ErrorFollw
		}

		err = tx.Table("tb_relation").Model(&model.Relation{}).
			Where("follower_id = ? AND following_id = ?", userID, toID).
			Update("isdeleted", 1).Error

		if err != nil {
			tx.Rollback()
			return utils.ErrorFollw
		}

		tx.Commit()

	}

	return nil
}
