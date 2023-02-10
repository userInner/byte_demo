package dao

import (
	"errors"
	"log"
	"titok_v1/common"
	"titok_v1/models"
	// "gorm"
)

type Favourite struct {
	UserId      int64
	VideoId     int64
	IsFavourite int64 // 0为关注 1为不关注
}

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

// 这里需要数据库的.sql 在本地试一试  Favourite的表是什么样的？
func GetFavouriteVideoListByUserId(userId int64) ([]int64, error) {
	var favouriteList []int64
	err := common.DB.Model(Favourite{}).Where(map[string]interface{}{"user_id": userId}).Pluck("video_id", &favouriteList).Error

	if err != nil {
		log.Printf(err.Error())
		return nil, errors.New("Method GetFavouriteeVideoListByUserId failed")
	} else {
		return favouriteList, nil
	}
	//这里没有测试 参照他人项目修改了函数 但是目前不确定
	//我觉得这个功能的DAO还要构思一下 近日会把文档发出来讨论
}


// 初步在DAO层写了通过用户ID返回点赞列表的方法
// func FavouriteAction(token string, videoId int64,actionType bool)([]int64 ,err){
// 	var favouriteSet []models.Favourite
// 	if result := DB.Select("video_id","is_favourite").Model(&models.Favourite{}).Where("user_id = ?", userId).Find(&favouriteSet);
// 	result.Error! = nil{
// 		return nil, result.Error
// 	}

// 	favouriteList = make([]int64, 0, len(favouriteSet))
// 	for _, each := range favouriteSet{
// 		favouriteList = append(favouriteSet, each.videoId)
// 	}
// }
