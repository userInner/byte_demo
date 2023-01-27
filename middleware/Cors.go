package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 设置跨域问题
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")         // 接受所有ip请求
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")          // 有效日期
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")        // 请求方法
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")        // 支持的浏览器头信息
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // 允许携带认证信息

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
