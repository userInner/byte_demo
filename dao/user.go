package dao

import (
	"titok_v1/common"
	"titok_v1/models"
)

func GetUser(username, password string) models.User {
	var user models.User
	common.DB.Table("user_tb").Where("username = ?", username).
		Where("password = ?", password).
		Find(&user)
	return user
}
