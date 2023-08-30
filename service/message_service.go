package service

import (
	"TikTok/db"
	"TikTok/model"
	"TikTok/utils"
	"fmt"
	"strconv"
	"time"
)

func SendMessageService(userID uint64, toUserID string, actionType string, content string) error {
	if actionType != "1" {
		return utils.ErrorActionType
	}
	toID, err := strconv.ParseUint(toUserID, 10, 64)
	if err != nil {
		return utils.ErrorUserID
	}

	if userID == toID {
		return utils.ErrorMessageUser
	}
	if !utils.CheckToUserID(toID) {
		return utils.ErrorUserID
	}
	err = db.DB.Table("tb_message").Create(&model.Message{
		ToUserID:   toID,
		FromUserID: userID,
		Content:    content,
		CreateTime: time.Now(),
	}).Error
	if err != nil {
		return utils.ErrorSendMessage
	}
	return nil

}

func ChatHistoryService(userID uint64, toUserID string) (messageList []model.Message, err error) {
	toID, err := strconv.ParseUint(toUserID, 10, 64)
	if err != nil {
		return nil, utils.ErrorUserID
	}

	if userID == toID {
		return nil, utils.ErrorMessageUser
	}
	if !utils.CheckToUserID(toID) {
		return nil, utils.ErrorUserID
	}

	err = db.DB.Table("tb_message").
		Where("to_user_id = ? AND from_user_id = ?", toID, userID).
		Find(&messageList).Error
	fmt.Println(messageList)

	if err != nil {
		return nil, utils.ErrorSendMessage
	}
	return messageList, nil
}
