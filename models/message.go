package models

import "time"

type Message struct {
	ID          	int64 `gorm:"primary_key; column:id"`
	FromUserID   	int64 `gorm:"ForeignKey:Author_id;AssociationForeignKey:ID"`
	ToUserID 		int64 `gorm:"ForeignKey:Receiver_id;AssociationForeignKey:ID"`
	Content     	string
	CreateTime 	time.Time `gorm:"column:create_time"`
	
}

func (v Message) tableName() string {
	return "message_tb"
}
