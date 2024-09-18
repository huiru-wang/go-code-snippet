package userservice

import (
	"go-gin/common/consts"
	"go-gin/common/syserror"
	"go-gin/model"
	"go-gin/service/dto"
	"go-gin/service/tokenservice"
	"go-gin/utils/encryptutil"
	"time"
)

type UserService struct {
}

func CreateUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) Register(userDto *dto.UserDto) *syserror.SysError {
	// TODO 用户名判重

	user := &model.User{
		UserId:     1321387128937, // TODO 用户唯一ID生成
		UserName:   userDto.UserName,
		Password:   encryptutil.Sha1(userDto.Password),
		Phone:      userDto.Phone,
		Status:     consts.UserStatusNormal,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return user.InsertOne()
}

func (userService *UserService) Login(userDto *dto.UserDto) (string, *syserror.SysError) {
	user := &model.User{
		UserName: userDto.UserName,
		Status:   consts.UserStatusNormal,
	}
	if err := user.SelectOne(); err != nil {
		return "", err
	}
	encryptPass := encryptutil.Sha1(userDto.Password)
	if encryptPass != user.Password {
		return "", syserror.UserPassNotMatch
	}
	// 创建token返回
	if token, err := tokenservice.CreateTokenService().CreateUserToken(user); err == nil {
		return token, nil
	}
	return "", syserror.CreateTokenError
}

func (userService *UserService) UpdatePhone(userDto *dto.UserDto) *syserror.SysError {
	user := &model.User{
		UserName: userDto.UserName,
		Password: encryptutil.Sha1(userDto.Password),
		Status:   consts.UserStatusNormal,
		Phone:    userDto.Phone,
	}
	if err := user.UpdatePhone(); err != nil {
		return err
	}
	return nil
}
