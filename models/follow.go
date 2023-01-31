package models

type Follow struct {
	UserID    uint64
	User_tb   User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	ToUserID  uint64
	ToUser_tb User `gorm:"ForeignKey:ToUserID;AssociationForeignKey:ID"`
	IsFollow  bool
}

func (v Follow) TableName() string {
	return "follow_tb"
}
