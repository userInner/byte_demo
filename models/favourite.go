package models

// 粉丝
type Favourite struct {
	UserId      uint64
	VideoId     uint64
	IsFavourite uint64 // 0为关注 1为不关注
}

func (v Favourite) TableName() string {
	return "favourite_tb"
}
