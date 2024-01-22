package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	RecordId   *int      `gorm:"column:recordId;primary_key"`
	UserName   string    `gorm:"column:userName"`
	CreateTime time.Time `gorm:"column:createTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = time.Now()
	return
}

func (u *User) TableName() (tableName string) {
	return "go_user"
}

func (u *User) Save() error {
	return nil
}
