package di

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/playaer/myFirstGoProject/managers"
	"github.com/playaer/myFirstGoProject/utils"
)

var instance DI = &repository{}

type DI interface {
	Db() *sql.DB
	UserManager() *managers.UserManager
}

type repository struct {
	db *sql.DB
	userManager *managers.UserManager
}

func New() DI {
	return instance
}

// Get database instance
func (di *repository) Db() *sql.DB {
	if (di.db == nil) {
		db, err := sql.Open("mysql", "root:@/first_go")
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
	}
	return di.userManager
}
