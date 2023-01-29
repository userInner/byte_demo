package routers

import (
	"github.com/gin-gonic/gin"
	"titok_v1/controllers"
)

func FeedRoutes(r *gin.RouterGroup) {
	feed := r.Group("feed")
	{
		feed.GET("", controllers.GetFeed)
		publish := feed.Group("publish")
		{
			publish.GET("/list/", controllers.GetUserVideo)
		}
	}
}
