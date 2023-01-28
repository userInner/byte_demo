package models

import "time"

// comment_tb
type Comment struct {
	ID         int64
	AuthorID   int64
	Content    string
	CreateDate time.Time
	CreateTime time.Time
	UpdateTime time.Time
	VideoId    int64
}

func (v Comment) TableName() string {
	return "comment_tb"
}
