package dao

import (
	"titok_v1/common"
	"titok_v1/models"
)

// 获取用户是否点赞
func GetFavourite(user *models.User, video *models.Video) bool {
	var favourite models.Favourite
	err := common.GetDB().
		Where("user_id = ? and video_id = ? and is_favourite = 1", user.ID, video.ID).
		First(&favourite).Error
	if err != nil {
		return false
	}
	if !favourite.IsFavourite {
		return false
	}
	return true
}
