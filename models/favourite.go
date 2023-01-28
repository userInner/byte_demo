package models

// 点赞
type Favourite struct {
	ID          int64
	UserId      int64
	VideoId     int64
	IsFavourite bool // 0为非好友 1为好友
}

func (v Favourite) TableName() string {
	return "favourite_tb"
}
