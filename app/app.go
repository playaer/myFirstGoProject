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
		Directory: "templates",
	}))


	userController := new(controllers.UserController)
	m.Get("/", userController.List)
	m.Get("/users/", userController.List)
	m.Get("/users/:id/view/", userController.View)
	m.Get("/users/:id/edit/", userController.Edit)
	m.Post("/users/:id/save/", userController.Save)

	registerController := new(controllers.RegisterController)
	m.Get("/register/", registerController.Register)
	m.Post("/register/processRegister/", registerController.ProcessRegister)

	authController := new(controllers.AuthController)
	m.Get("/auth/", authController.Login)
	m.Post("/auth/processLogin/", authController.ProcessLogin)

	m.Run()
}
