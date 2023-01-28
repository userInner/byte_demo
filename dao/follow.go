package dao

import (
	"titok_v1/common"
	"titok_v1/models"
)

// 获取用户是否关注
func GetFollow(user *models.User, to_user *models.User) bool {
	var follow models.Follow
	err := common.GetDB().
		Where("user_id = ? and to_user_id = ? and is_follow = 1", user.ID, to_user.ID).
		First(&follow).Error
	if err != nil {
		return false
	}
	if !follow.IsFollow {
		return false
	}
	return true
}
