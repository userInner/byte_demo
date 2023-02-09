package models

import (
	"time"
)

type User struct {
	ID            int64    `json:"id" gorm:"primary_key; column:id"`            // 用户id
	UserName      string    `json:"name" gorm:"column:username"`                 // 用户名称
	Password      string    `gorm:"column:password"`                             // 用户密码
	CreateTime    time.Time `gorm:"column:create_time"`                          // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time"`                          // 更新时间
	FollowCount   int64    `json:"follow_count" gorm:"column:follow_count"`     // 关注总数
	FollowerCount int64    `json:"follower_count" gorm:"column:follower_count"` // 粉丝总数
	IsFollow      bool
}

func (v User) TableName() string {
	return "user_tb"
}
