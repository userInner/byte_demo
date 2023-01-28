package routers

import (
	"github.com/gin-gonic/gin"
	"titok_v1/controllers"
	"titok_v1/middleware"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())

	douyin := r.Group("douyin")
	{
		douyin.GET("/feed/", controllers.FeedDemo)
		// 用户相关
		UserRoutes(douyin)
		// 视频
		VideoRoutes(douyin)
	}

	return r
}
