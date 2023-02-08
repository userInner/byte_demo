package routers

import (
	// "net/http"
	// "titok_v1/middleware"

	"github.com/gin-gonic/gin"
)

func RelationRouters(r *gin.RouterGroup) {
	relation := r.Group("relation")
	{
		FriendsRouters(relation)
	}
}
