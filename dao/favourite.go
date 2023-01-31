package dao

import (
	"errors"
	"log"
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

// 点赞功能空函数

// 这里的实现 可能得取取决于Favourite表的元素 讨论定下来以后后1月31日解决
func AddFavourite() {}

// 取消功能空函数

// 问题同上
func CancelFavourite() {}
