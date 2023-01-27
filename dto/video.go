package dto

import "titok_v1/models"

// Video
type VideoDto struct {
	Author        models.User `json:"author"`         // 视频作者信息
	CommentCount  int64       `json:"comment_count"`  // 视频的评论总数
	CoverURL      string      `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64       `json:"favorite_count"` // 视频的点赞总数
	ID            int64       `json:"id"`             // 视频唯一标识
	IsFavorite    bool        `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string      `json:"play_url"`       // 视频播放地址
	Title         string      `json:"title"`          // 视频标题
}
