package models

type Follow struct {
	ID       int64
	UserID   int64
	ToUserID int64
}
