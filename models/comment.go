package models

import "time"

// comment_tb
type Comment struct {
	ID         int64  `gorm:"column:id"`
	AuthorID   int64  `gorm:"column:author_id"`
	Author     User   `gorm:"foreignKey:AuthorID"`
	VideoID    int64  `gorm:"column:video_id"`
	Content    string `gorm:"column:content"`
	CreateDate string
	CreateTime time.Time `gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time"` // 更新时间
}

func (v Comment) TableName() string {
	return "comment_tb"
}
