package dao

import (
	"titok_v1/common"
	"titok_v1/models"
)

// 点赞
func GetFavourite(u *models.User, video *models.Video) bool {
	fav := &models.Favourite{}
	err := common.GetDB().
		Where("user_id=? and video_id =?", u.ID, video.ID).
		First(fav).Error
	if err != nil {
		return false
	}
	if fav.UserId == 0 || fav.IsFavourite == 0 {
		return false
	}
	return true

}
