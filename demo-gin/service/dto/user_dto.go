package dto

type UserDto struct {
	UserId   int64  `form:"userId" json:"userId"`
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Phone    string `form:"phone" json:"phone"`
}
