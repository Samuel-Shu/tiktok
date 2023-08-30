package utils

import (
	"TikTok/db"
	"TikTok/model"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func CheckToUserID(toUserID uint64) bool {
	var user model.User
	result := db.DB.Table("tb_user").First(&user, toUserID)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true

}

func CheckActionType(actionType string, toUserID string) bool {
	if actionType != "1" && actionType != "2" {
		return false
	}

	userID, err := strconv.ParseUint(toUserID, 10, 64)
	if err != nil {
		return false
	}

	if actionType == "1" {
		var relation model.Relation
		result := db.DB.First(&relation, userID)
		if result.Error != nil {
			return false
		}
	}
	if actionType == "2" {

	}
	return false
}
