package routers

import (
	"net/http"
	"titok_v1/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 200,
			"msg":  "访问错误",
		})
	})
	douyin := r.Group("douyin")
	{
		// 用户相关
		UserRoutes(douyin)
		// 视频
		FeedRoutes(douyin)
		PublishRouters(douyin)
		// 评论
		CommentRouters(douyin)
	}

	return r
}
