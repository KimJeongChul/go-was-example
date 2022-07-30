package datatype

import "time"

type User struct {
	UserId     string
	UserName   string
	CreateTime time.Time
}
