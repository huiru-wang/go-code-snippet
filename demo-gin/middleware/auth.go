package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin/common/consts"
	"go-gin/common/response"
	"go-gin/service/tokenservice"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			response.AuthError(ctx, consts.NoAuthorizationMsg)
		}
		if isValid := tokenservice.CreateTokenService().ValidateUserToken(token); !isValid {
			response.AuthError(ctx, consts.AuthorizationFailMsg)
			return
		}
		ctx.Next()
	}
}
