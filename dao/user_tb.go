package dao

import (
	"titok_v1/common"
	"titok_v1/models"
)

func IsExistUser(username string) bool {
	var user models.User
	var count int64

	db := common.GetDB()
	db.Table("user_tb").Where("username = ?", username).Find(&user).Count(&count)
	return count == 0
}

func InsertUser(newUser *models.User) error {
	db := common.GetDB()
	if err := db.Model(&models.User{}).Create(newUser).Error; err != nil {
		return err
	}
	return nil
}
