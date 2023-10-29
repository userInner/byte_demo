package models

import "time"

type Friends struct {
	Id          int64 `gorm:"primary_key"`
	User_id     int64
	User_tb     User `gorm:"ForeignKey:User_id;AssociationForeignKey:ID"`
	Friend_id   int64
	Friend_tb   User      `gorm:"ForeignKey:Friend_id;AssociationForeignKey:ID"`
	Create_time time.Time `gorm:"column:create_time"`
	Update_time time.Time `gorm:"column:update_time"`
}

func (v Friends) TableName() string {
	return "friends_tb"
}
