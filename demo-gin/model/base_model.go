package model

type BaseModel struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	CreatedAt string `gorm:"column:create_time; type: timestamp;not null"`
	UpdatedAt string `gorm:"column:update_time; type: timestamp;not null"`
	//DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
