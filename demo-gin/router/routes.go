package router

import (
	v1 "go-gin/api/v1"
	"go-gin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	routes := r.Group("/api/v1")

	RegisterUserRouter(routes)

	UserBusinessRouter(routes)
}

func RegisterUserRouter(r *gin.RouterGroup) {
	r.POST("/user/sign/in", v1.Login)
	r.POST("/user/sign/up", v1.Register)
}

func UserBusinessRouter(r *gin.RouterGroup) {
	r.Use(middleware.AuthMiddleware())
	r.POST("/user/update/phone", v1.UpdatePhone)
}
