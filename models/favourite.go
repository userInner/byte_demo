package models

// 粉丝
type Favourite struct {
	UserId      int64
	VideoId     int64
	IsFavourite bool // 0为非好友 1为好友
}
