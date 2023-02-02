package models

import "time"

type Message struct {
	Id          uint64 `gorm:"primary_key; column:id"`
	Author_id   uint64
	Author_tb   User `gorm:"ForeignKey:Author_id;AssociationForeignKey:ID"`
	Receiver_id uint64
	Receiver_tb User `gorm:"ForeignKey:Receiver_id;AssociationForeignKey:ID"`
	Content     string
	Create_time time.Time `gorm:"column:create_time"`
	Update_time time.Time `gorm:"column:update_time"`
}

func (v Message) tableName() string {
	return "message_tb"
}
