package dao

import (
	"log"
	"titok_v1/common"
	"titok_v1/models"
)

func GetUserFollow(u *models.User, to_user *models.User) bool {
	follow := &models.Follow{}
	err := common.GetDB().
		Where("user_id=? and to_user_id=? and is_follow=1", u.ID, to_user.ID).
		Find(follow).Error
	if err != nil {
		log.Println("数据库错误", err.Error)
		return false
	}
	if follow.IsFollow == 0 {
		return false
	}
	return true
}
