package models

type Follow struct {
	UserID   int64
	ToUserID int64
	IsFollow int32
}

func (v Follow) TableName() string {
	return "follow_tb"
}
