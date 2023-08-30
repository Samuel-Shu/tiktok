package model

import "time"

type Message struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ToUserID   uint      `gorm:"column:to_user_id" json:"to_user_id"`
	FromUserID uint      `gorm:"column:from_user_id" json:"from_user_id"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}
