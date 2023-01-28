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
		return nil, errors.New("get video failed" + err.Error())
	}
	return videos, nil
}

// 获取用户所有投稿视频
func GetVideoByUser(u models.User) ([]models.Video, error) {
	var videos []models.Video
	err := common.GetDB().
		Preload("Author").
		Where("author_id=?", u.ID).
		Find(&videos).Error
	if err != nil {
		return nil, errors.New("数据库错误" + err.Error())
	}
	return videos, nil
}
