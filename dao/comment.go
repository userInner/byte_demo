package dao

import (
	"errors"
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

// 插入评论
func InsertComment(comment *models.Comment) (*models.Comment, error) {
	err := common.GetDB().Create(comment).Error
	if err != nil {
		return nil, errors.New("数据库错误" + err.Error())
	}
	return comment, nil
}

// 判断是否有该条评论
func GetCommentIsExist(comment *models.Comment) (bool, error) {
	resComment := &models.Comment{}
	err := common.GetDB().Where("author_id=? and video_id=? and id=?", comment.AuthorID, comment.VideoID, comment.ID).
		First(resComment).Error
	if err != nil {
		return false, err
	}
	if resComment.ID == 0 {
		return false, nil
	}
	return true, nil
}

// 删除评论
func DeleteComment(comment *models.Comment) (bool, error) {
	exist, err := GetCommentIsExist(comment)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	err = common.GetDB().
		Delete(comment, comment.ID).
		Error
	if err != nil {
		return false, err
	}

	return true, nil
}

// 获取视频所有评论
func GetCommentByVideoID(video *models.Video) ([]models.Comment, error) {
	comment := []models.Comment{}
	err := common.GetDB().Preload("Author").
		Where("video_id=?", video.ID).
		Find(&comment).Error
	if err != nil {
		return nil, errors.New("数据库错误" + err.Error())
	}
	return comment, nil

}
