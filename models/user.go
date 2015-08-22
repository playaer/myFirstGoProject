package models

type User struct {
	Id int64
	FullName string
	Address string
	Phone string
	Email string
	Password string
	Hash string
	IsActive bool
	Token string
}
