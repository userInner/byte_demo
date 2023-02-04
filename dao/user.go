package dao

import (
	"errors"
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

func GetUserByID(user_id int64) (*models.User, error) {
	var u models.User
	err := common.GetDB().Where("id=?", user_id).Find(&u).Error
	if err != nil {
		return nil, errors.New("数据库查询失败" + err.Error())
	}
	return &u, nil
}
