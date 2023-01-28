package routers

import (
	"github.com/gin-gonic/gin"
	"titok_v1/middleware"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())

	douyin := r.Group("douyin")
	{
		// 用户相关
		UserRoutes(douyin)
		// 视频
		FeedRoutes(douyin)
	}

	return r
}
