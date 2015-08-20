package models

type User struct {
	Id int64     `db:"id, primarykey, autoincrement"`
	FullName string `db:"full_name"`
	Address string `db:"address"`
	Phone string `db:"phone"`
}
