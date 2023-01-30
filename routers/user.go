package routers

import (
	"github.com/gin-gonic/gin"
	"titok_v1/controllers"
)

func UserRoutes(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		user.POST("/register/", controllers.UserRegister)
		user.POST("/login/", controllers.UserLogin)
		// TODO 需要鉴权后才能访问 UserInfo
		user.GET("/", controllers.UserInfo)
	}
}
