package dto

import (
	// "time"
	"titok_v1/models"
)

type FriendsDto struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

type FriendsListDto struct {
	FriendsList []FriendsDto `json:"user_list"` // 好友列表
	StatusCode  string       `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string       `json:"status_msg"`   // 返回状态描述
}

func BuildFriendsDto(friends *models.User) *FriendsDto {
	if friends == nil {
		return nil
	}
	return &FriendsDto{
		FollowCount:   friends.FollowerCount,
		FollowerCount: friends.FollowerCount,
		ID:            friends.ID,
		IsFollow:      false,
		Name:          friends.UserName,
	}
}

func BuildFriendsListDto(code string, msg string, friends []models.User) *FriendsListDto {
	if len(friends) == 0 ||friends == nil {
		return &FriendsListDto{
			FriendsList: nil,
			StatusCode:  code,
			StatusMsg:   msg,
		}
	}
	resFriendsDtos := make([]FriendsDto, len(friends))
	for k, _ := range friends {
		resFriendsDtos[k].ID = friends[k].ID
		resFriendsDtos[k].FollowCount = friends[k].FollowCount
		resFriendsDtos[k].FollowerCount = friends[k].FollowerCount
		resFriendsDtos[k].IsFollow = friends[k].IsFollow
		resFriendsDtos[k].Name = friends[k].UserName
	}
	return &FriendsListDto{
		FriendsList: resFriendsDtos,
		StatusCode:  code,
		StatusMsg:   msg,
	}
}
