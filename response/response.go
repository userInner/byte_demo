package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	1000～1999 区间表示参数错误
	2000～2999 区间表示用户错误
	3000～3999 区间表示接口异常
*/

func Response(c *gin.Context, httpStatus, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 400, data, msg)
}
