package service

import (
	"gorm"
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
func (f *favouriteService) FavoriteAction(userId int64, videoId int64, actionType int32) (json, error) {

	flag, err := GetFavouriteStatus(userId, videoId)
	var status_msg string
	var status_code int32

	if actionType == 1 {
		if flag == true {
			c.JSON(http.status_code, http.status_msg)
			return true, nil
		} else {
			var favourite models.favourite

			favourite.UserId = userId
			favourite.VideoId = videoId
			favourite.IsFavourite = true
			c.JSON(http.status_code, http.status_msg)

			if err := DB.Create(&favourite).Error; err != nil {
				log.Printf("Insert favourite log failed :%v", err)
				c.JSON(http.status_code, http.status_msg)
				return false, err
			}
		}
	}
	if actionType == 0 {
		if flag == false {
			c.JSON(http.status_code, http.status_msg)
			return true, nil
		} else {

			var favourite models.favourite
			favourite.UserId = userId
			favourite.VideoId = videoId
			favourite.IsFavourite = true

			//这里应该要再改一下
			if err := DB.Where("user_id = ? and video_id = ?", userId, videoId).Delete(); err != nil {
				log.Printf("Delete favourite log failed :%v", err)

				c.JSON(http.status_code, http.status_msg)
				return false, err
			}
		}
	}

	return true, nil //默认返回

	//点赞逻辑：查询列表用户是否已经点赞，如果已经点赞则不修改，如果未点赞则加入。
}

// 在DAO层实现了 这里直接调用DAO层的函数
func GetFavouriteVideoListByUserId(userId int64) []int64 {
	favouriteList, err := dao.GetFavouriteVideoListByUserId(userId)
	if err != nil {
		log.Printf("Get Favouritee Video List By UserID failed :%v", err)
		return err
	}
	return favouriteList
}
