package main

import (
	app "github.com/playaer/myFirstGoProject/app"
	"github.com/playaer/myFirstGoProject/di"
)

func main() {
	di := di.New()
	di.UserManager().SetDb(di.Db())
	app.Run(&di)
}
