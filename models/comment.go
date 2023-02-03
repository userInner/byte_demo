package models

import "time"

// comment_tb
type Comment struct {
	ID         uint64 `gorm:"column:id"`
	AuthorID   uint64 `gorm:"column:author_id"`
	Author     User   `gorm:"foreignKey:AuthorID"`
	VideoID    uint64 `gorm:"column:video_id"`
	Content    string `gorm:"column:content"`
	CreateDate string
	CreateTime time.Time `gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time"` // 更新时间
}

func (v Comment) TableName() string {
	return "comment_tb"
}
