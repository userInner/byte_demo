package models

// 粉丝
type Favourite struct {
	UserId      int64
	VideoId     int64
	IsFavourite int64 // 0为关注 1为不关注
}

func (v Favourite) TableName() string {
	return "favourite_tb"
}
