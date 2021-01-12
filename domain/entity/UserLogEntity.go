package entity

import "time"

type UserLogEntity struct {
	Id int
	UserId int
	UserName string
	LogType int
	LogComment string
	CreatedTime time.Time
}
