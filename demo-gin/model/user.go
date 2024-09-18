package model

import (
	"go-gin/common/consts"
	mysql "go-gin/common/database"
	"go-gin/common/syserror"
	"time"
)

type User struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	UserId     int64     `gorm:"column:user_id;type:bigint(64);not null;unique"`
	UserName   string    `gorm:"column:user_name;type:varchar(20);not null"`
	Phone      string    `gorm:"column:phone; type:varchar(11); not null;unique"`
	Password   string    `gorm:"column:password; type: varchar(64) size:255;not null"`
	Status     int       `gorm:"column:status; type: int(1);not null"`
	CreateTime time.Time `gorm:"column:create_time; type: timestamp;not null"`
	UpdateTime time.Time `gorm:"column:create_time; type: timestamp;not null"`
}

func (user *User) TableName() string {
	return "t_user"
}

func (user *User) InsertOne() *syserror.SysError {
	db := mysql.Session()
	var count int64
	db.Model(user).Where("user_name = ?", user.UserName).Count(&count)
	if count > 0 {
		return syserror.UserNameExits
	}
	if result := db.Model(user).Create(user); result.RowsAffected != 1 {
		return syserror.RegisterError
	}
	return nil
}

func (user *User) SelectOne() *syserror.SysError {
	db := mysql.Session()
	if result := db.Model(user).Where("user_name = ? and status = ?", user.UserName, user.Status).First(user); result.RowsAffected < 1 {
		return syserror.UserNameNotExits
	}
	return nil
}

func (user *User) UpdatePhone() *syserror.SysError {
	db := mysql.Session()
	sql := "update t_user set phone = ? where user_name = ? and status = ?"
	result := db.Model(user).Exec(sql, user.Phone, user.UserName, user.Status)
	if result.Error == nil && result.RowsAffected == 1 {
		return nil
	}
	return syserror.Error(consts.UserPhoneUpdateFailCode, consts.UserPhoneUpdateFailMsg)
}

func (user *User) Validate() bool {
	db := mysql.Session()
	var count int64
	db.Model(user).Where("user_id = ? and user_name = ?", user.UserId, user.UserName).Count(&count)
	if count < 1 {
		return false
	}
	return true
}
