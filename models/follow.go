package models

// 关注
type Follow struct {
	UserID   int64
	ToUserID int64
	IsFollow bool //  false为不关注 true为关注
}

func (v Follow) TableName() string {
	return "follow_tb"
}
