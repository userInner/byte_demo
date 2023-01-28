package dao

import (
	"errors"
	"titok_v1/common"
	"titok_v1/models"
)

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
}
