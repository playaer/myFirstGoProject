package managers

import (
	"github.com/playaer/myFirstGoProject/models"
	"database/sql"
	"github.com/playaer/myFirstGoProject/utils"
	"crypto/md5"
	"encoding/hex"
)

type UserManager struct {
	db *sql.DB
}

func (self *UserManager) SetDb(db *sql.DB) {
	self.db = db
}

// Find all users
func (self *UserManager) FindAll() ([]*models.User, error) {

	db := self.db
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := new(models.User)
		err = rows.Scan(&user.Id, &user.FullName, &user.Address, &user.Phone, &user.Email, &user.Password)
		if err != nil {
			utils.CheckErr(err, err.Error())
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		utils.CheckErr(err, err.Error())
	}

	return users, nil
}

// Find user by id
func (self *UserManager) FindById(id string) *models.User {
	db := self.db
	user := new(models.User)
	stmt, err := db.Prepare("SELECT * FROM users WHERE id =  ?")
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	err = stmt.QueryRow(id).Scan(&user.Id, &user.FullName, &user.Address, &user.Phone, &user.Email, &user.Password)
	if err != nil {
		utils.CheckErr(err, err.Error())
	}

	if err == sql.ErrNoRows {
		return nil
	} else {
		return user
	}
}

// Find user by email
func (self *UserManager) FindByEmail(email string) *models.User {
	db := self.db
	user := new(models.User)
	stmt, err := db.Prepare("SELECT * FROM users WHERE email =  ?")
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	err = stmt.QueryRow(email).Scan(&user.Id, &user.FullName, &user.Address, &user.Phone, &user.Email, &user.Password)
	if err != nil {
		utils.CheckErr(err, err.Error())
	}

	if err == sql.ErrNoRows {
		return nil
	} else {
		return user
	}
}

// Create user
func (self *UserManager) Create(user *models.User) int64 {
	db := self.db

	stmt, err := db.Prepare("INSERT INTO users(full_name, address, phone, email, password) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	res, err := stmt.Exec(user.FullName, user.Address, user.Phone, user.Email, user.Password)
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		utils.CheckErr(err, err.Error())
	}

	return lastId
}

// Update user
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

// Get empty user
func (self *UserManager) NewUser() *models.User {
	return new(models.User)
}

// Crypt password
// Low security and w/o salt
func (self *UserManager) CryptPassword(rawPassword string) string {
	hasher := md5.New()
	hasher.Write([]byte(rawPassword))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Check password
func (self *UserManager) CheckPassword(user *models.User, rawPassword string) bool {
	return user.Password == self.CryptPassword(rawPassword)
}

