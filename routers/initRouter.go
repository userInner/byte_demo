package routers

import (
	"github.com/gin-gonic/gin"
	"titokl_v1/controllers"
	"titokl_v1/middleware"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.GET("/douyin/feed/", controllers.GetFeed)
	return r
}
