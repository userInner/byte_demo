package dto

import "titok_v1/models"

// 视频作者信息
//
// User
type UserDto struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

type UserInfoDto struct {
	StatusCode int64    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `json:"status_msg"`  // 返回状态描述
	User       *UserDto `json:"user"`        // 用户信息
}

// BuildUserInfoDto 构建userDTO
func (u UserInfoDto) BuildUserInfoDto(code int64, msg string, user *models.User) *UserInfoDto {
	uDto := &UserDto{
		FollowCount:   user.FollowerCount,
		FollowerCount: user.FollowerCount,
		ID:            user.ID,
		IsFollow:      false,
		Name:          user.UserName,
	}
	return &UserInfoDto{
		StatusCode: code,
		StatusMsg:  &msg,
		User:       uDto,
	}
}
