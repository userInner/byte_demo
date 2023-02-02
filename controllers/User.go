package controllers

import (
	"log"
	"net/http"
	"strconv"
	"titok_v1/dao"
	"titok_v1/dto"
	"titok_v1/middleware"
	"titok_v1/models"
	resp "titok_v1/response"
	"titok_v1/service"

	"github.com/gin-gonic/gin"
)

var (
	InvalidParams = "非法用户名或密码"
	DataError     = "数据库错误"
	SuccessData   = "数据库查询成功"
)

func UserRegister(c *gin.Context) {
	//username := c.Query("username")
	//password := c.Query("password")
	var userService service.UserService
	if err := c.ShouldBindQuery(&userService); err != nil {
		log.Printf("c.ShouldBind(&userService): %s\n", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}

	log.Printf("userService: %v\n", userService)
	if len(userService.Name) > 32 || len(userService.Passwd) > 32 { //最长32位字符
		resp.Fail(c, nil, InvalidParams)
		return
	}

	if userService.Name == "" || userService.Passwd == "" {
		resp.Fail(c, nil, InvalidParams)
		return
	}

	log.Println("register...")
	res := userService.Register(c)
	c.JSON(http.StatusOK, res)
}

func UserLogin(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err != nil {
		log.Println("UserLogin: ", userService)
		resp.Fail(c, nil, "登录失败")
		return
	}

	res, _ := userService.Login(c)
	c.JSON(http.StatusOK, res)
}

/*
参数；user_id,token
*/
func UserInfo(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	if len(user_id) == 0 {
		c.JSON(http.StatusOK, dto.UserInfoDto{}.BuildUserInfoDto(
			1001, InvalidParams, nil,
		))
		return
	}
	target_u_id, err := strconv.Atoi(user_id)
	if err != nil {
		c.JSON(http.StatusOK, dto.UserInfoDto{}.BuildUserInfoDto(
			1001, "参数无效"+err.Error(), nil,
		))
		return
	}
	// 判断token是否有效
	u_id, err := middleware.VerifyToken(token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadGateway, dto.UserInfoDto{
			1001, &InvalidParams, nil,
		})
		return
	}
	// 查询token所属用户
	u, err := dao.GetUserByID(u_id)
	if err != nil {
		c.JSON(401, dto.UserInfoDto{
			1001, &DataError, nil,
		})
		return
	}
	if u.ID == 0 {
		c.JSON(200, dto.UserInfoDto{ // 查询成功，但是没有数据
			1001, &SuccessData, nil,
		})
		return
	}
	// 查询用户是否关注user_id
	respUserInfo := dto.UserInfoDto{}.BuildUserInfoDto(0, "查询成功", u)
	target_u := &models.User{ID: int64(target_u_id)}
	is_follow := dao.IsUserFollow(u, target_u)
	if !is_follow {
		c.JSON(http.StatusOK, respUserInfo)
		return
	}
	//respUserInfo.User.IsFollow = dao.GetUserFollow()
	c.JSON(http.StatusOK, respUserInfo)
}
