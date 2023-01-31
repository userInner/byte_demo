package controllers

import (
	"net/http"
	"strconv"
	"time"
	"titok_v1/dao"
	"titok_v1/dto"
	"titok_v1/middleware"
	"titok_v1/models"
	"titok_v1/response"
	"titok_v1/utils"

	"github.com/gin-gonic/gin"
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
		timeUnix = int64(resTime)
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

	respFeed := dto.BuildFeed(last_time, 0, "查询成功", videos)

	// 校验token
	u_id, err := middleware.VerifyToken(token)
	if err != nil { //无效token
		c.JSON(http.StatusOK, respFeed)
		return
	}
	// 关注
	u := &models.User{ID: u_id}
	for k, v := range videos {
		// 点赞
		video_u := &models.User{ID: v.Author.ID}
		respFeed.VideoList[k].IsFavorite, _ = dao.GetFavourite(&models.User{ID: u_id}, &models.Video{ID: v.ID})

		// 关注视频用户
		respFeed.VideoList[k].Author.IsFollow = dao.IsUserFollow(u, video_u)
	}
	c.JSON(http.StatusOK, respFeed)
}

// 查询用户投稿视频
func GetUserVideo(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")
	if len(token) == 0 {
		c.JSON(401, dto.BuildUserFeed(1001, "参数无效", nil))
		return
	}
	u_id, err := middleware.VerifyToken(token)
	if err != nil {
		c.JSON(401, dto.BuildUserFeed(1001, "暂无权限", nil))
		return
	}
	target_u_id, err := strconv.Atoi(user_id)
	if err != nil {
		c.JSON(401, dto.BuildUserFeed(1001, "暂无权限", nil))
		return
	}
	my_user := &models.User{ID: u_id}
	target_user := &models.User{ID: int64(target_u_id)}
	videos, err := dao.GetVideoByUser(*target_user)
	for _, v := range videos {
		v.Author.IsFollow = dao.IsUserFollow(my_user, target_user)
		v.IsFavorite, _ = dao.GetFavourite(my_user, &v)
	}
	c.JSON(200, dto.BuildUserFeed(0, "查询成功", videos))
}
