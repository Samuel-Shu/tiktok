package model

import (
	"time"
)

type Relation struct {
	ID          uint64    `gorm:"primaryKey;column:id"` // 关注 ID
	FollowerID  uint64    `gorm:"column:follower_id"`   // 粉丝 ID
	FollowingID uint64    `gorm:"column:following_id"`  // 博主 ID
	IsDeleted   int       `gorm:"column:isdeleted"`     // 是否取消关注
	CreateTime  time.Time `gorm:"column:create_time"`   // 创建时间
}
