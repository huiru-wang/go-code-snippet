package model

import "time"

type Tag struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	Tag        string    `gorm:"column:tag; type: varchar(256)"`
	UserId     int64     `gorm:"column:user_id;type:bigint(64);not null"`
	CreateTime time.Time `gorm:"column:create_time; type: bigint(64);not null"`
}

func (Tag) TableName() string {
	return "t_tag"
}
