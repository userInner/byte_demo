package dao

import (
	"log"
	"titok_v1/common"
	"titok_v1/models"
)

// 获取视频评论总数
func GetCommentCount(video models.Video) int64 {
	var cnt int64
	err := common.GetDB().Where("video_id=?", video.ID).Find(&models.Comment{}).Count(&cnt).Error
	if err != nil {
		log.Println("数据库错误")
		return 0
	}
	return cnt
}
