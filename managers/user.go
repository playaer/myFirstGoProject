package managers

import (
	"github.com/playaer/myFirstGoProject/models"
	"database/sql"
	"github.com/playaer/myFirstGoProject/utils"
)

type UserManager struct {
	db *sql.DB
}

func (self *UserManager) SetDb(db *sql.DB) {
	self.db = db
}

func (self *UserManager) FindAll() ([]*models.User, error) {

	db := self.db
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := new(models.User)
		err = rows.Scan(&user.Id, &user.FullName, &user.Address, &user.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (self *UserManager) FindById(id string) *models.User {
	db := self.db
	user := new(models.User)
	stmt, err := db.Prepare("SELECT * FROM users WHERE id =  ?")
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	err = stmt.QueryRow(id).Scan(&user.Id, &user.FullName, &user.Address, &user.Phone)
	if err != nil {
		utils.CheckErr(err, err.Error())
	}

	if err == sql.ErrNoRows {
		return nil
	} else {
		return user
	}
}

func (self *UserManager) Create(user *models.User) int64 {
	db := self.db

	stmt, err := db.Prepare("INSERT INTO users(full_name, address, phone) VALUES(?, ?, ?)")
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	res, err := stmt.Exec(user.FullName, user.Address, user.Phone)
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		utils.CheckErr(err, err.Error())
	}

	return lastId
}

func (self *UserManager) Update(user *models.User) {
	db := self.db

	stmt, err := db.Prepare("UPDATE users SET full_name=?, address=?, phone=? WHERE id = ?")
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	_, err = stmt.Exec(user.FullName, user.Address, user.Phone, user.Id)
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
}

func (self *UserManager) NewUser() *models.User {
	return new(models.User)
}
