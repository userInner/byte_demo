package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"titok_v1/dao"
	"titok_v1/middleware"
	"titok_v1/models"
	resp "titok_v1/response"
	"titok_v1/utils"
)

var (
	secret = "gegege"

	ErrExistUser      = "该用户已存在"
	ErrUsername       = "用户名或密码错误"
	ErrServerInternal = "服务端错误"
)

type UserService struct {
	Name   string `form:"username" json:"username"`
	Passwd string `form:"password" json:"password"`
}

func (u *UserService) Login(c *gin.Context) *resp.UserLoginResp {
	u.Passwd = utils.Md5Crypt(u.Passwd, secret)
	if !checkUser(u.Name, u.Passwd) {
		log.Println("checkUser(u.Name, u.Passwd) not correct")
		resp.Fail(c, nil, ErrUsername)
		return nil
	}

	user := dao.GetUser(u.Name, u.Passwd)
	token, err := middleware.GenToken(user.ID, user.UserName)
	if err != nil {
		resp.Fail(c, nil, ErrServerInternal)
		return nil
	}

	return &resp.UserLoginResp{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     user.ID,
		Token:      token,
	}

}

func (user *UserService) Register(c *gin.Context) *resp.UserLoginResp {
	if exist := dao.IsExistUser(user.Name); !exist {
		return &resp.UserLoginResp{
			StatusCode: 1,
			StatusMsg:  ErrExistUser,
		}
	}

	newUser := &models.User{
		UserName:      user.Name,
		Password:      utils.Md5Crypt(user.Passwd, secret),
		FollowCount:   0,
		FollowerCount: 0,
		Create_time:   time.Now().Local(),
		Update_time:   time.Now().Local(),
	}

	err := dao.InsertUser(newUser)
	if err != nil {
		return &resp.UserLoginResp{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  ErrServerInternal,
		}
	}

	token, err := middleware.GenToken(newUser.ID, newUser.UserName)
	if err != nil {
		return &resp.UserLoginResp{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  ErrServerInternal,
		}
	}

	return &resp.UserLoginResp{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     newUser.ID,
		Token:      token,
	}
}

func checkUser(name, password string) bool {
	user := dao.GetUser(name, password)
	return user.UserName == name && user.Password == password
}
