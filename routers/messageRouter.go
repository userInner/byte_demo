package routers

import (
	"github.com/gin-gonic/gin"
	"titok_v1/controllers"
)

func MessageRouters(r *gin.RouterGroup) {
	message := r.Group("message")
	{
		// message.POST("/send/", controllers.SendMessage)
		message.GET("/list/", controllers.MessageList)
	}
}
