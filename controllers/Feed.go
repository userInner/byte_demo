package controllers

import (
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
	if len(latest_time) != 0 {
		// 校验token
		if utils.VaildToken(token) {

			return
		}
		timeUnix, _ := strconv.Atoi(latest_time)
		t := time.Unix(int64(timeUnix), 0).Format("2006-01-02 15:04:05")

		videos, err := dao.GetVideo(t)
		if err != nil {
			response.Fail(c, gin.H{
				"msg": "查询失败",
			}, "3001")
		}

		// 得到最新投稿时间
		last := len(videos)
		last_time := utils.GetTimeInt64(videos[last-1])
		feed := dto.BuildFeed(last_time, http.StatusOK, "查询成功", videos)
		c.JSON(http.StatusOK, feed)
	}
}
