package main

import (
	app "github.com/playaer/myFirstGoProject/app"
	"github.com/playaer/myFirstGoProject/utils"
)

func main() {
	app.Run()

	defer func() {
		if r := recover(); r != nil {
			utils.Debug(r)
		}
	}()
}
