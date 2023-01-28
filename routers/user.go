package routers

import (
	"github.com/gin-gonic/gin"
	"titok_v1/controllers"
)

func UserRoutes(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		user.GET("", controllers.UserInfo)
		user.POST("/register/", controllers.UserRegister)
		user.POST("/login/", controllers.UserLogin)
	}
}
