package models

type Follow struct {
	ID        int64
	UserID    int64
	User_tb   User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	ToUserID  int64
	ToUser_tb User `gorm:"ForeignKey:ToUserID;AssociationForeignKey:ID"`
	IsFollow  bool
}

func (v Follow) TableName() string {
	return "follow_tb"
}
