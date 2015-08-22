package di

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/playaer/myFirstGoProject/managers"
	"github.com/playaer/myFirstGoProject/utils"
	"github.com/playaer/myFirstGoProject/config"
)

var instance DI = &repository{}

type DI interface {
	Db() *sql.DB
	UserManager() *managers.UserManager
	UpdateLogManager() *managers.UpdateLogManager
	AuthManager() *managers.AuthManager
	Config() *config.Config
	Mailer() *utils.Mailer
}

type repository struct {
	db *sql.DB
	userManager *managers.UserManager
	updateLogManager *managers.UpdateLogManager
	authManager *managers.AuthManager
	config *config.Config
	mailer *utils.Mailer
}

func New() DI {
	return instance
}

// Get database instance
func (di *repository) Db() *sql.DB {
	if (di.db == nil) {
		config := instance.Config()
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
func (di *repository) UserManager() *managers.UserManager {
	if (di.userManager == nil) {
		di.userManager = new(managers.UserManager)
		db := instance.Db()
		di.userManager.SetDb(db)
	}
	return di.userManager
}

// Get updateLogManager
func (di *repository) UpdateLogManager() *managers.UpdateLogManager {
	if (di.updateLogManager == nil) {
		di.updateLogManager = new(managers.UpdateLogManager)
		db := instance.Db()
		di.updateLogManager.SetDb(db)
	}
	return di.updateLogManager
}

// Get auth manager
func (di *repository) AuthManager() *managers.AuthManager {
	if (di.authManager == nil) {
		di.authManager = new(managers.AuthManager)
		db := instance.Db()
		di.authManager.SetDb(db)
	}
	return di.authManager
}

// Get app config
func (di *repository) Config() *config.Config {
	if (di.config == nil) {
		di.config = config.New()
	}
	return di.config
}

// Get app config
func (di *repository) Mailer() *utils.Mailer {
	if (di.mailer == nil) {
		mailer := new(utils.Mailer)
		di.mailer = mailer.New(instance.Config())
	}
	return di.mailer
}
