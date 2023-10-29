package service

import (
	"log"
	"titok_v1/common"
	"titok_v1/dao"
	"titok_v1/models"
	resp "titok_v1/response"
)

type favouriteService struct {
	user  models.User
	video models.Video
}

// 初步的getFavouriteStatus方法
func (f *favouriteService) GetFavouriteStatus(userId int64, videoId int64) bool {
	favouriteStatus, err := dao.GetFavourite(&f.user, &f.video)
	if err == nil {
		return favouriteStatus
	} else {
		log.Printf("Method getFavoriteStatus failed :%v", err)
		return false
	}
}

// 点赞的操作 根据传入的actionType区分是点赞和取消
func (f *favouriteService) FavoriteAction(userId int64, videoId int64, actionType int32) *resp.FavouriteActionResp {

	flag := f.GetFavouriteStatus(userId, videoId)

	if actionType == 1 {
		if flag == true {
			return &resp.FavouriteActionResp{
				StatusCode: 0,
				StatusMsg:  "successed",
			}
		} else {

			var favourite models.Favourite
			favourite.UserId = userId
			favourite.VideoId = videoId
			favourite.IsFavourite = 1
			if err := common.DB.Create(&favourite).Error; err != nil {
				log.Printf("Insert favourite log failed :%v", err)
				return &resp.FavouriteActionResp{
					StatusCode: -1,
					StatusMsg:  "failed",
				}
			}
		}
	}
	if actionType == 0 {
		if flag == false {
			return &resp.FavouriteActionResp{
				StatusCode: 0,
				StatusMsg:  "successed",
			}
		} else {

			var favourite models.Favourite
			favourite.UserId = userId
			favourite.VideoId = videoId
			favourite.IsFavourite = 1

			if err := common.DB.Where("user_id = ? and video_id = ?", userId, videoId).Update("IsFavourite", 0); err != nil {
				log.Printf("Delete favourite log failed :%v", err)
				return &resp.FavouriteActionResp{
					StatusCode: -1,
					StatusMsg:  "failed",
				}
			}
		}
	}

	return &resp.FavouriteActionResp{
		StatusCode: 0,
		StatusMsg:  "successed",
	}

	//点赞逻辑：查询列表用户是否已经点赞，如果已经点赞则不修改，如果未点赞则加入。
}

// 在DAO层实现了 这里直接调用DAO层的函数
func GetFavouriteVideoListByUserId(userId int64) []models.Favourite {
	favouriteList, err := dao.GetFavouriteVideoListByUserId(userId)
	if err != nil {
		log.Printf("Get Favouritee Video List By UserID failed :%v", err)
		return nil
	}

	//从video_id中拿author信息

	//从video_id中拿play_url, cover_url, 点赞数，评论数和标题。

	//要用GetCover

	//要用GetMinioUrl

	//拿video的Favourite Count

	//拿video的标题

	//拿Video的评论数 没有方法要写

	return favouriteList
}
