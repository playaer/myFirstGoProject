package app

import (
	"github.com/go-martini/martini"
	controllers "github.com/playaer/myFirstGoProject/controllers"
	"github.com/martini-contrib/render"
)

func Run() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
		Directory: "public/templates",
	}))


	userController := new(controllers.UserController)

	m.Get("/", userController.List)
	m.Get("/users/", userController.List)
	m.Get("/users/:id/view/", userController.View)
	m.Get("/users/:id/edit/", userController.Edit)
	m.Post("/users/:id/save/", userController.Save)

	m.Run()
}
