package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"titok_v1/dao"
	"titok_v1/dto"
	"titok_v1/models"
)

/*
获取用户信息
获取用户的 id、昵称，
如果实现社交部分的功能，还会返回关注数和粉丝数

请求方法 query
请求类型 string
请求参数 user_id 用户id，token 鉴权

根据token查找用户，比对user_id用户,是否关注
*/
func GetUserInfo(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")

	// user_id 转整数
	u_id, err := strconv.Atoi(user_id)
	if err != nil {
		c.JSON(404, dto.UserInfoDto{}.BuildUserInfoDto(
			1001, "非法参数", nil,
		))
	}
	// token是否有效
	tagetUser, err := dao.GetUser(&models.User{ID: int64(u_id)})
	if err != nil {
		c.JSON(422, dto.UserInfoDto{}.BuildUserInfoDto(
			3001, "查询失败", nil))
		return
	}
	if tagetUser.ID == 0 {
		c.JSON(200, dto.UserInfoDto{}.BuildUserInfoDto(
			2000, "查询成功", tagetUser))
	}

}
