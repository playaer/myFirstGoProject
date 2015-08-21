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
	Config() *config.Config
}

type repository struct {
	db *sql.DB
	userManager *managers.UserManager
	config *config.Config
}

func New() DI {
	return instance
}

// Get database instance
func (di *repository) Db() *sql.DB {
	if (di.db == nil) {
		config := instance.Config()
		db, err := sql.Open(config.DbDriver, config.DbUser + ":" + config.DbPass + "@/" + config.DbName)
		utils.CheckErr(err, "sql.Open failed")

		if err = db.Ping(); err != nil {
			utils.CheckErr(err, "sql.Ping failed")
		}
		di.db = db
	}
	return di.db
}

// Get database instance
func (di *repository) UserManager() *managers.UserManager {
	if (di.userManager == nil) {
		di.userManager = new(managers.UserManager)
		db := instance.Db()
		di.userManager.SetDb(db)
	}
	return di.userManager
}

// Get app config
func (di *repository) Config() *config.Config {
	if (di.config == nil) {
		di.config = config.New()
	}
	return di.config
}
