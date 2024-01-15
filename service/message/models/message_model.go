package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ToUserId   uint   `json:"to_user_id"`
	FromUserId uint   `json:"from_user_id"`
	Content    string `json:"content"`
}

func (Message) TableName() string {
	return "tb_message"
}

type DefaultMessageModel struct {
	Db *gorm.DB
}

func NewMessageModel(db *gorm.DB) *DefaultMessageModel {
	return &DefaultMessageModel{
		Db: db,
	}
}

func (m *DefaultMessageModel) Create(toUserId, formUserId uint, content string) (*Message, error) {
	message := &Message{
		ToUserId:   toUserId,
		FromUserId: formUserId,
		Content:    content,
	}
	err := m.Db.Create(message).Error
	return message, err
}

func (m *DefaultMessageModel) GetByToUserIdAndUserId(toUserId, formUserId uint) ([]Message, error) {
	list := make([]Message, 0)
	err := m.Db.Where("to_user_id = ? and from_user_id = ?", toUserId, formUserId).Find(&list).Error
	return list, err
}
