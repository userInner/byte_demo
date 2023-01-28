package dto

import (
	"titok_v1/dao"
	"titok_v1/models"
)

type Feed struct {
	NextTime   int64      `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64      `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string     `json:"status_msg"`  // 返回状态描述
	VideoList  []VideoDto `json:"video_list"`  // 视频列表
}

func BuildFeed(nextTime int64, statusCode int64, statusMsg string, videoList []models.Video) *Feed {
	feed := new(Feed)
	feed.NextTime = nextTime
	feed.StatusCode = statusCode
	feed.StatusMsg = statusMsg
	tempVideoList := make([]VideoDto, 0)
	for _, v := range videoList {
		// 获取
		author := VideoDto{
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

		tempVideoList = append(tempVideoList, author)
	}
	feed.VideoList = tempVideoList
	return feed
}
