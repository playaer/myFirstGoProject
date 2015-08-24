package managers

import (
	"github.com/playaer/myFirstGoProject/models"
	"database/sql"
	"github.com/playaer/myFirstGoProject/utils"
	"crypto/md5"
	"encoding/hex"
	"time"
)

type AuthManager struct {
	db *sql.DB
	currentUser *models.User
}

func (self *AuthManager) SetDb(db *sql.DB) {
	self.db = db
}

func (self *AuthManager) IsAuthenticated() bool {
	return self.currentUser != nil
}

// Find active user by token
func (self *AuthManager) FindActiveByToken(token string) (*models.User, error) {
	db := self.db
	user := new(models.User)
	stmt, err := db.Prepare("SELECT * FROM users WHERE token = ? AND is_active = 1")
	utils.CheckErr(err, nil)
	err = stmt.QueryRow(token).Scan(&user.Id, &user.FullName, &user.Address, &user.Phone, &user.Email, &user.Password, &user.Hash, &user.IsActive, &user.Token)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

// Update user
func (self *AuthManager) Update(user *models.User) {
	db := self.db
	stmt, err := db.Prepare("UPDATE users SET token=? WHERE id = ?")
	utils.CheckErr(err, nil)
	_, err = stmt.Exec(user.Token, user.Id)
	utils.CheckErr(err, nil)
}

// Generate token
func (self *AuthManager) GenerateToken(user *models.User) string {
	hasher := md5.New()
	hasher.Write([]byte(user.Email + time.Now().Format(time.StampNano)))
	token := hex.EncodeToString(hasher.Sum(nil))

	user.Token = token
	self.Update(user)

	return token
}

// Login user. Store current user to local variable
func (self *AuthManager) Auth(user *models.User) {
	self.currentUser = user
}

// Logout user. Clear current user from local variable
func (self *AuthManager) Logout() {
	user := *self.currentUser
	self.currentUser = nil
	user.Token = ""
	self.Update(&user)
}

// Get current authenticated user
func (self *AuthManager) CurrentUser() *models.User {
	return self.currentUser
}



