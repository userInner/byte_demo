package dao

import (
	"errors"
	"titok_v1/common"
	"titok_v1/models"
)

type Favourite struct {
	UserId      int64
	VideoId     int64
	IsFavourite int64 // 0为关注 1为不关注
}

// 点赞
// 由于业务层需要增加一个error返回值
// 这里把数据库表中的0和1映射成true或者false
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

func GetFavouriteVideoListByUserId(userId int64) ([]models.Favourite, error) {
	var favouriteSet []models.Favourite
	err := common.GetDB().Where("user_id = ? and is_favourite = 1").Find(&favouriteSet)

	if err != nil {
		return nil, errors.New("GetFavouritedVideoListByUserId failed." + err.Error.Error())
	}
	return favouriteSet, nil
}
