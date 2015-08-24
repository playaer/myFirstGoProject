package managers

import (
	"github.com/playaer/myFirstGoProject/models"
	"database/sql"
	"github.com/playaer/myFirstGoProject/utils"
	"time"
	"encoding/json"
)

type UpdateLogManager struct {
	db *sql.DB
}

func (self *UpdateLogManager) SetDb(db *sql.DB) {
	self.db = db
}

func (self *UpdateLogManager) StoreChanges(oldUser *models.User, newUser *models.User) {
	log := new(models.UserUpdate)
	log.UserId = newUser.Id
	log.OldData = self.prepareData(oldUser)
	log.NewData = self.prepareData(newUser)
	log.UpdatedAt = time.Now()

	self.saveToDb(log)
}

func (self *UpdateLogManager) prepareData(user *models.User) map[string]string {
	data := map[string]string{
		"FullName": user.FullName,
		"Address": user.Address,
		"Phone": user.Phone,
	}

	return data
}

// save new user update record
func (self *UpdateLogManager) saveToDb(userUpdate *models.UserUpdate) int64 {
	db := self.db

	stmt, err := db.Prepare("INSERT INTO user_update(user_id, old_data, new_data, updated_at) VALUES(?, ?, ?, ?)")
	utils.CheckErr(err, nil)
	oldData, err := json.Marshal(userUpdate.OldData)
	newData, err := json.Marshal(userUpdate.NewData)
	res, err := stmt.Exec(userUpdate.UserId, oldData, newData, userUpdate.UpdatedAt)
	utils.CheckErr(err, nil)
	lastId, err := res.LastInsertId()
	utils.CheckErr(err, nil)

	return lastId
}

// Find all user updates
func (self *UpdateLogManager) FindAll(userId int64) []*models.UserUpdate {

	db := self.db
	rows, err := db.Query("SELECT updated_at, old_data, new_data FROM user_update WHERE user_id = ? ORDER BY id DESC", userId)
	utils.CheckErr(err, nil)
	defer rows.Close()

	var oldData, newData []byte
	userUpdates := make([]*models.UserUpdate, 0)
	for rows.Next() {
		userUpdate := new(models.UserUpdate)
		err = rows.Scan(&userUpdate.UpdatedAt, &oldData, &newData)
		json.Unmarshal(oldData, &userUpdate.OldData)
		json.Unmarshal(newData, &userUpdate.NewData)

		utils.CheckErr(err, nil)
		userUpdates = append(userUpdates, userUpdate)
	}

	if err = rows.Err(); err != nil {
		utils.CheckErr(err, nil)
	}


	return userUpdates
}



