package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"titok_v1/dao"
	"titok_v1/dto"
	"titok_v1/response"
	"titok_v1/utils"
)

/*
 api 首页视频
 日期： 2023-1-27 16:56
*/

func GetFeed(c *gin.Context) {
	latest_time := c.Query("latest_time")
	token := c.Query("token")

	// 数据库查询最大的时间戳视频
	var timeUnix int64

	if len(latest_time) != 0 { // 有上传时间戳
		resTime, _ := strconv.Atoi(latest_time)
		timeUnix = (int64(resTime))
	} else { // 无上传时间戳
		timeUnix = time.Now().Unix()
	}
	t := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")

	videos, err := dao.GetVideo(t)
	if err != nil {
		response.Fail(c, gin.H{
			"msg":   "查询失败",
			"error": err,
		}, "3001")
		return
	}

	// 得到最新投稿时间
	last := len(videos)
	last_time := utils.GetTimeInt64(videos[last-1].CreateTime.String())
	// 校验token
	if utils.VaildToken(token) { // 有上传token,并且为合法token
		// 关注
		for _, v := range videos {
			fmt.Println(v)
			// 解析token获取用户，再对比视频作者，用户是否是粉丝
			//v.Is_favorite = dao.GetFavourite()
		}
		return
	}
	feed := dto.BuildFeed(last_time, http.StatusOK, "查询成功", videos)
	c.JSON(http.StatusOK, feed)
}
