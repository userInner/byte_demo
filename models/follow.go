package models

type Follow struct {
<<<<<<< HEAD
	ID        uint64
=======
>>>>>>> b2cb668523580da494ed0f502e9f763dc42b5086
	UserID    uint64
	User_tb   User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	ToUserID  uint64
	ToUser_tb User `gorm:"ForeignKey:ToUserID;AssociationForeignKey:ID"`
	IsFollow  bool
}

func (v Follow) TableName() string {
	return "follow_tb"
}
