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
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	oldData, err := json.Marshal(userUpdate.OldData)
	newData, err := json.Marshal(userUpdate.NewData)
	res, err := stmt.Exec(userUpdate.UserId, oldData, newData, userUpdate.UpdatedAt)
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		utils.CheckErr(err, err.Error())
	}

	return lastId
}

// Find all user updates
func (self *UpdateLogManager) FindAll(userId int64) []*models.UserUpdate {

	db := self.db
	rows, err := db.Query("SELECT updated_at, old_data, new_data FROM user_update")
	if err != nil {
		utils.CheckErr(err, err.Error())
	}
	defer rows.Close()

	var oldData, newData []byte
	userUpdates := make([]*models.UserUpdate, 0)
	for rows.Next() {
		userUpdate := new(models.UserUpdate)
		err = rows.Scan(&userUpdate.UpdatedAt, &oldData, &newData)
		json.Unmarshal(oldData, &userUpdate.OldData)
		json.Unmarshal(newData, &userUpdate.NewData)
//		userUpdate.NewData = json.Unmarshal(newData, map[string]string)

//		utils.Fatal(userUpdate)
		if err != nil {
			utils.CheckErr(err, err.Error())
		}
		userUpdates = append(userUpdates, userUpdate)
	}

	if err = rows.Err(); err != nil {
		utils.CheckErr(err, err.Error())
	}


	return userUpdates
}



