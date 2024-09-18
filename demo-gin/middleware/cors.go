package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CorsMiddleware 允许跨域：允许当前域被其他域访问
func CorsMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")  //*表示所有
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")   //最长时间
		context.Writer.Header().Set("Access-Control-Allow-Methods", "*") //允许访问的方法
		context.Writer.Header().Set("Access-Control-Allow-Headers", "*") //允许的请求头
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 放通 OPTIONS METHOD
		if context.Request.Method == http.MethodOptions {
			context.AbortWithStatus(http.StatusAccepted)
		}
		context.Next()
	}
}
