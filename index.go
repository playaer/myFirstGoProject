package main

import (
	app "github.com/playaer/myFirstGoProject/app"
	"github.com/playaer/myFirstGoProject/utils"
)

func main() {
	di := utils.New()
	di.UserManager().SetDb(di.Db())
	app.Run()
}
