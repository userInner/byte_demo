package routers

import (
	"github.com/gin-gonic/gin"
	"titok_v1/controllers"
)

func PublishRouters(r *gin.RouterGroup) {
	publish := r.Group("publish")
	{
		publish.POST("/action/", controllers.PublishVideoByUser)

		publish.GET("/list/", controllers.GetUserVideo)

	}
}
