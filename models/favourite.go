package models

// 粉丝
type Favourite struct {
	ID          int64 // 唯一标志物
	UserId      int64
	VideoId     int64
	IsFavourite int64 // 0为非好友 1为好友
}
