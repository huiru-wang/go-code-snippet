package response

import (
	"go-gin/common/consts"
	"go-gin/common/syserror"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func buildBody(ctx *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

func SuccessWithToken(ctx *gin.Context, token string, data interface{}) {
	ctx.Writer.Header().Set("Authorization", token)
	buildBody(ctx, http.StatusOK, consts.SuccessCode, consts.SuccessMsg, data)
}

func Success(ctx *gin.Context, data interface{}) {
	buildBody(ctx, http.StatusOK, consts.SuccessCode, consts.SuccessMsg, data)
}

func ParamError(ctx *gin.Context, msg string) {
	buildBody(ctx, http.StatusBadRequest, consts.ParamErrorCode, consts.ParmaErrorMsg, "")
	ctx.Abort()
}

func AuthError(ctx *gin.Context, msg string) {
	buildBody(ctx, http.StatusUnauthorized, consts.NoAuthorizationCode, msg, "")
	ctx.Abort()
}

func Error(ctx *gin.Context, error *syserror.SysError) {
	buildBody(ctx, http.StatusOK, error.ErrorCode, error.ErrorMsg, "")
	ctx.Abort()
}
