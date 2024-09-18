package tokenservice

import (
	"go-gin/common/syserror"
	"go-gin/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenService struct {
}

func CreateTokenService() *TokenService {
	return &TokenService{}
}

var jwtKey = []byte("my key")

// UserClaims userToken 载荷
type UserClaims struct {
	UserId   int64
	UserName string
	jwt.StandardClaims
}

func (TokenService) CreateUserToken(user *model.User) (string, *syserror.SysError) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	userClaims := &UserClaims{
		UserId:   user.UserId,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mess",
			Subject:   "User Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	// 生成 userToken 签名
	if userToken, err := token.SignedString(jwtKey); err == nil {
		return userToken, nil
	}
	return "", syserror.SystemError
}

func (TokenService) ValidateUserToken(userToken string) bool {
	userClaims := &UserClaims{}
	token, err := jwt.ParseWithClaims(userToken, userClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return false
	}
	user := &model.User{UserId: userClaims.UserId, UserName: userClaims.UserName}
	// 用户是否存在
	return user.Validate()
}

func (TokenService) FreshUserToken() {

}

func (TokenService) DisableUserToken() {

}
