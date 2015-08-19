package app

import (
	"github.com/go-martini/martini"
	controllers "github.com/playaer/myFirstGoProject/controllers"
)

var m *ClassicMartini
m := martini.Classic()

func Run() {
	configureRoutes()
	m.Run()
}

func configureRoutes() {
	m.Get("/", controllers.User.list)
	m.Get("/users/", controllers.User.View)
	m.Get("/users/:id/view/", controllers.User.View)
	m.Get("/users/:id/edit/", controllers.User.Edit)
	m.Post("/users/:id/save/", controllers.User.Save)
}