package routers

import (
	"github.com/gin-gonic/gin"
	"titok_v1/controllers"
)

func CommentRouters(r *gin.RouterGroup) {
	comment := r.Group("comment")
	{
		comment.POST("/action/", controllers.CommentAction)
		comment.GET("/list/", controllers.CommentList)
	}
}
