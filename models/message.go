package models

import "time"

type Message struct {
	ID          int64 `gorm:"primary_key; column:id"`
	Author_id   int64 `gorm:"ForeignKey:Author_id;AssociationForeignKey:ID"`
	
	Receiver_id int64 `gorm:"ForeignKey:Receiver_id;AssociationForeignKey:ID"`

	Content     string
	CreateDate  string
	Create_Time time.Time `gorm:"column:create_time"`
	//Update_Time time.Time `gorm:"column:update_time"`
}

func (v Message) tableName() string {
	return "message_tb"
}
