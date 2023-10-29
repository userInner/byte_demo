package dto

import (
	"titok_v1/dao"
	"titok_v1/models"
)

// Video
type VideoDto struct {
	Author        UserDto `json:"author"`         // 视频作者信息
	CommentCount  int64   `json:"comment_count"`  // 视频的评论总数
	CoverURL      string  `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64   `json:"favorite_count"` // 视频的点赞总数
	ID            int64   `json:"id"`             // 视频唯一标识
	IsFavorite    bool    `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string  `json:"play_url"`       // 视频播放地址
	Title         string  `json:"title"`          // 视频标题
}

func BuildFeedDto(videoList []models.Video) []VideoDto {
	tempVideoList := make([]VideoDto, len(videoList))
	for k, v := range videoList {
		// 获取
		dto := VideoDto{
			Author: UserDto{
				FollowCount:   v.Author.FollowCount,
				FollowerCount: v.Author.FollowerCount,
				ID:            v.Author.ID,
				IsFollow:      false, // 是否关注
				Name:          v.Author.UserName,
			},
			CommentCount:  dao.GetCommentCount(v),
			CoverURL:      v.CoverURL,
			FavoriteCount: v.FavouriteCount,
			ID:            v.ID,
			IsFavorite:    false,
			PlayURL:       v.PlayURL,
			Title:         v.Title,
		}
		tempVideoList[k] = dto
	}
	return tempVideoList
}
