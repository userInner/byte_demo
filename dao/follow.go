package dao

import (
	"log"
	"titok_v1/common"
	"titok_v1/models"
)

// 是否关注对方
func IsUserFollow(u *models.User, to_user *models.User) bool {
	follow := &models.Follow{}
	err := common.GetDB().
		Where("user_id=? and to_user_id=? and is_follow=1", u.ID, to_user.ID).
		Find(follow).Error
	if err != nil {
		log.Println("数据库错误", err.Error)
		return false
	}
	if !follow.IsFollow {
		return false
	}
	return true
}

// 关注对方
func FollowUser(u *models.User, to_user *models.User) {
	common.GetDB().Model(&models.Follow{}).Create(map[string]interface{}{
		"UserID":   u.ID,
		"ToUserID": to_user.ID,
		"IsFollow": true,
	})
	u.FollowCount++
	to_user.FollowerCount++
}

// 返回关注列表
func FollowTB(u *models.User, to_user *models.User) (user []models.User) {
	var users []models.User
	if IsUserFollow(u, to_user) {
		common.GetDB().First(&users, to_user.ID)

	}
	return users
}

// 返回粉丝列表
func FollowersTB(u *models.User, to_user *models.User) (user []models.User) {
	var users []models.User
	if IsUserFollow(u, to_user) {
		common.GetDB().First(&users, u.ID)
	}
	return users
}
