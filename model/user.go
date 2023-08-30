package model

import "time"

// User /*
//
//	"signature": "string",
//	"total_favorited": "string",
//	"work_count": 0,
//	"favorite_count": 0
//
// */
type User struct {
	ID              uint      `gorm:"primaryKey;column:id" json:"id"`                  // 用户 ID
	Name            string    `gorm:"column:name" json:"name,omitempty"`               // 用户名
	FollowCount     int       `gorm:"column:follow_count" json:"follow_count"`         // 关注总数
	FollowerCount   int       `gorm:"column:follower_count" json:"follower_count"`     // 粉丝总数
	BackgroundImage string    `gorm:"column:background_image" json:"background_image"` // 用户个人页顶部大图URL
	Signature       string    `gorm:"column:signature" json:"signature"`               // 个人简介
	TotalFavorited  int       `gorm:"column:total_favorited" json:"total_favorited"`   // 获赞数量
	WorkCount       int       `gorm:"column:work_count" json:"work_count"`             // 作品数量
	FavoriteCount   int       `gorm:"column:favorite_count" json:"favorite_count"`     // 点赞数量
	Password        string    `gorm:"column:password" json:"password"`                 // 密码
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`           // 创建时间
	Avatar          string    `gorm:"column:avatar" json:"avatar,omitempty"`           // 用户头像
	IsFollow        bool      `gorm:"-" json:"is_follow"`
}
