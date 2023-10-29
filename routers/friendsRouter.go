package routers

import (
	"github.com/gin-gonic/gin"
	"titok_v1/controllers"
)

func FriendsRouters(r *gin.RouterGroup) {
	friend := r.Group("friend")
	{
		friend.GET("/list/", controllers.FriendsList)
	}
}
