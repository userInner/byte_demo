package service

import (
	"log"
	"titok_v1/dao"
	"titok_v1/models"
)

type favouriteService struct {
	user  models.User
	video models.Video
}

// 初步的getFavouriteStatus方法
func (f *favouriteService) GetFavouriteStatus(userId int64, videoId int64) (bool, error) {
	favouriteStatus, err := dao.GetFavourite(&f.user, &f.video)
	if err == nil {
		return favouriteStatus, nil
	} else {
		log.Printf("Method getFavoriteStatus failed :%v", err)
		return false, err
	}
}

// 点赞的操作 根据传入的actionType区分是点赞和取消
func (f *favouriteService) FavoriteAction(userId int64, videoId int64, actionType int32) (bool, error) {

	//首先获取用户的点赞列表
	favouriteList, err := dao.GetFavouriteVideoListByUserId(f.user.ID)

	flag := IsInArray(favouriteList)

	//点赞逻辑：查询列表用户是否已经点赞，如果已经点赞则不修改，如果未点赞则加入。

	if flag == true {
	} else {
	}

	//取消逻辑：查询列表用户是否已经点赞，如果未点赞则不修改，如果已点赞则取消。
	if flag == true {
	} else {
	}
	return true, err

}

// 根据用用户ID获取点赞列表的空函数
func GetFavouriteVideoListByUserId(userId int64) []int64 {
	favouriteList, err := dao.GetFavouriteVideoListByUserId(userId)
	if err != nil {
		return nil
	}

	return favouriteList
}

// 判断元素是否在切片中的空函数
func IsInArray([]int64) bool {
	return true
}
