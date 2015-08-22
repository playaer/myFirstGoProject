package models

import "time"

type UserUpdate struct {
	Id int64
	UserId int64
	UpdatedAt time.Time
	OldData map[string]string
	NewData map[string]string
}
