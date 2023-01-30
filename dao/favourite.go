package dao

import (
	"titok_v1/common"
	"titok_v1/models"
)

// 点赞
// 由于业务层需要增加一个error返回值
func GetFavourite(u *models.User, video *models.Video) (bool, error) {
	fav := &models.Favourite{} //favorite
	err := common.GetDB().
		Where("user_id=? and video_id =?", u.ID, video.ID).
		First(fav).Error
	if err != nil {
		return false, err
	}
	if fav.UserId == 0 || fav.IsFavourite == 0 {
		return false, err
	}
	return true, err

}

// DAO层通过ID查询用户赞过视频的空函数
func GetFavouriteVideoListByUserId(userId int64) ([]int64, error) {
	var favouriteList []int64
	return favouriteList, nil
}

// 点赞功能空函数
func AddFavourite() {}

// 取消功能空函数
func CancelFavourite() {}
