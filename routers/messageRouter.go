package routers

import (
	"github.com/gin-gonic/gin"
	"titok_v1/controllers"
)

func MessageRouters(r *gin.RouterGroup) {
	message := r.Group("message")
	{
		message.POST("/action/", controllers.SendMessage)
		message.GET("/chat/", controllers.MessageList)
	}
}
