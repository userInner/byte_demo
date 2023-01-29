package routers

import (
	"github.com/gin-gonic/gin"
	"titok_v1/controllers"
)

func VideoRoutes(r *gin.RouterGroup) {
	publish := r.Group("publish")
	{
		publish.GET("/list/", controllers.GetUserVideo)
	}

}
