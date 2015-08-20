package managers

import (
	"github.com/playaer/myFirstGoProject/models"
	"database/sql"
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

func (self *UserManager) FindById(id string) models.User {
	return models.User{}
}
