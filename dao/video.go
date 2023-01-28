package dao

import (
	"errors"
	"titok_v1/common"
	"titok_v1/models"
)

// GetVideo 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
func GetVideo(time string) ([]models.Video, error) {
	var videos []models.Video
	err := common.GetDB().
		Preload("Author").
		Limit(30).
		Where("create_time <= ?", time).
		Order("create_time desc").
		Find(&videos).Error
	if err != nil {
		return nil, errors.New("get video failed")
	}
	return videos, nil
}

// 获取登录状态的video
// user 代表本身
// to 所关注的用户
func GetOnVideo(user *models.User, to *models.User, time string) ([]models.Video, error) {
	// 查询 点赞状态 是否关注
	var videos []models.Video
	common.GetDB().
		Preload("User").
		Order("create_time desc").
		Limit(30).
		Where("create_time < time").
		Preload("Favourite").
		Where("user_id = ? and is_favourite = 1 and to_user_id = ? ").
		Find(&videos)

	return nil, nil
}
