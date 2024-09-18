package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin/common/response"
	"go-gin/common/syserror"
	"go-gin/service/dto"
	"go-gin/service/userservice"
)

func Register(ctx *gin.Context) {
	userDto := &dto.UserDto{}
	if err := ctx.BindJSON(userDto); err != nil {
		response.Error(ctx, syserror.ParamError)
	}
	userService := userservice.CreateUserService()
	if err := userService.Register(userDto); err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, "")
}

func Login(ctx *gin.Context) {
	userDto := &dto.UserDto{}
	if err := ctx.BindJSON(userDto); err != nil {
		response.Error(ctx, syserror.ParamError)
	}
	token, err := userservice.CreateUserService().Login(userDto)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.SuccessWithToken(ctx, token, "")
}

func UpdatePhone(ctx *gin.Context) {
	userDto := &dto.UserDto{}
	if err := ctx.BindJSON(userDto); err != nil {
		response.Error(ctx, syserror.ParamError)
	}
	if err := userservice.CreateUserService().UpdatePhone(userDto); err != nil {
		response.Error(ctx, err)
	}
	response.Success(ctx, "")
}
