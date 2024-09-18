package model

import "time"

type Article struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	Tag        string    `gorm:"column:tag; type: varchar(256)"`
	Title      string    `gorm:"column:title; type: varchar(256)"`
	Content    string    `gorm:"column:content; type: longtext" `
	UserId     int64     `gorm:"column:user_id; type: bigint(64)"`
	State      int       `gorm:"column:state; type: int(1)"`
	CreateTime time.Time `gorm:"column:create_time; type: timestamp;not null"`
	UpdateTime time.Time `gorm:"column:update_time; type: timestamp;not null"`
}

func (Article) TableName() string {
	return "t_article"
}
