package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	resp "titok_v1/response"
	"titok_v1/service"
)

var (
	InvalidParams = "非法用户名或密码"
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

	res := userService.Login(c)
	c.JSON(http.StatusOK, res)
}
