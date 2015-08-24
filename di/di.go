package di

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/playaer/myFirstGoProject/managers"
	"github.com/playaer/myFirstGoProject/utils"
	"github.com/playaer/myFirstGoProject/config"
)

//var di DI = DI{}

//type DI interface {
//	Db() *sql.DB
//	UserManager() *managers.UserManager
//	UpdateLogManager() *managers.UpdateLogManager
//	AuthManager() *managers.AuthManager
//	Config() *config.Config
//	Mailer() *utils.Mailer
//}

type DI struct {
	db *sql.DB
	userManager *managers.UserManager
	updateLogManager *managers.UpdateLogManager
	authManager *managers.AuthManager
	config *config.Config
	mailer *utils.Mailer
}

func New() *DI {
	utils.Debug("di init!")

	return new(DI)
}

// Get database di
func (di *DI) Db() *sql.DB {
	if (di.db == nil) {
		config := di.Config()
		db, err := sql.Open(config.DbDriver, config.DbUser + ":" + config.DbPass + "@/" + config.DbName + config.DbParams)
		utils.CheckErr(err, "sql.Open failed")

		if err = db.Ping(); err != nil {
			utils.CheckErr(err, "sql.Ping failed")
		}
		di.db = db
	}
	return di.db
}

// Get user manager
func (di *DI) UserManager() *managers.UserManager {
	if (di.userManager == nil) {
		di.userManager = new(managers.UserManager)
		db := di.Db()
		di.userManager.SetDb(db)
	}
	return di.userManager
}

// Get updateLogManager
func (di *DI) UpdateLogManager() *managers.UpdateLogManager {
	if (di.updateLogManager == nil) {
		di.updateLogManager = new(managers.UpdateLogManager)
		db := di.Db()
		di.updateLogManager.SetDb(db)
	}
	return di.updateLogManager
}

// Get auth manager
func (di *DI) AuthManager() *managers.AuthManager {
	if (di.authManager == nil) {
		utils.Debug("auth manager init!")
		di.authManager = new(managers.AuthManager)
		db := di.Db()
		di.authManager.SetDb(db)
	}
	utils.Debug("auth manager get!")

	return di.authManager
}

// Get app config
func (di *DI) Config() *config.Config {
	if (di.config == nil) {
		di.config = config.New()
	}
	return di.config
}

// Get app config
func (di *DI) Mailer() *utils.Mailer {
	if (di.mailer == nil) {
		mailer := new(utils.Mailer)
		di.mailer = mailer.New(di.Config())
	}
	return di.mailer
}
