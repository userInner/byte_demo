package models

import (
	"time"
)

type Video struct {
	AuthorID       int64     `-`                                         // 视频作者ID
	Author         User      `json:"author" gorm:"foreignKey:AuthorID;"` // 视频作者信息
	CoverURL       string    `json:"cover_url" gorm:"column:cover_url"`  // 视频封面地址
	ID             int64     `json:"id" gorm:"primary_key; column:id"`   // 视频唯一标识
	PlayURL        string    `json:"play_url" gorm:"column:play_url"`    // 视频播放地址
	Title          string    `json:"title" gorm:"column:title"`          // 视频标题
	FavouriteCount int64     `json:"favorite_count"`
	CreateTime     time.Time `json:"create_time" gorm:"column:create_time;type:datetime"` // 创建时间
	UpdateTime     time.Time `json:"update_time" gorm:"column:update_time;type:datetime"` // 更新时间
	IsFavorite     bool      `gorm:"-:all"`
}

func (v Video) TableName() string {
	return "video_tb"
}
