package dao

import (
<<<<<<< HEAD
	"errors"
=======
>>>>>>> d09eb2a798bda77aae1211fa5abdc6322e2a6e2f
	"titok_v1/common"
	"titok_v1/models"
)

<<<<<<< HEAD
func GetUser(u *models.User) (*models.User, error) {
	user := new(models.User)
	err := common.GetDB().
		Where("id=?", u.ID).
		Find(&user).
		Error
	if err != nil {
		return nil, errors.New("数据查询失败" + err.Error())
	}
	return user, nil
=======
func GetUser(username, password string) models.User {
	var user models.User
	common.DB.Table("user_tb").Where("username = ?", username).
		Where("password = ?", password).
		Find(&user)
	return user
>>>>>>> d09eb2a798bda77aae1211fa5abdc6322e2a6e2f
}
