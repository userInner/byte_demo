package dao

import (
	"titok_v1/common"
	"titok_v1/models"
)

// 获取用户的关注信息
func GetFavourite(user models.User, to_User models.User) (bool) {
	var favourite models.Favourite
	err := common.GetDB().
		Where("user_id = ? and to_user_id = ? and is_follow = 1").
		First(favourite).Error
	if err != nil {
		return false
	}
	if !favourite.IsFavourite {
		return false
	}
	return true

}
