package database

import "time"

type User struct {
	UserId     string `gorm:"primaryKey"`
	UserName   string
	CreateTime time.Time
}

func (User) TableName() string {
	return "user"
}
