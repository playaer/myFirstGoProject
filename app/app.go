package app

import (
	"github.com/go-martini/martini"
	controllers "github.com/playaer/myFirstGoProject/controllers"
	"github.com/martini-contrib/render"
	"github.com/playaer/myFirstGoProject/di"
	"html/template"
	"time"
)

func Run(di *di.DI) {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
		Directory: "templates",
		Funcs: []template.FuncMap{
			{
				"formatTime": func(t *time.Time) string {
					return t.Format(time.Stamp)
				},
			},
		},
	}))


	userController := new(controllers.UserController)
	userController.SetDi(di)
	m.Get("/", userController.List)
	m.Get("/users/", userController.List)
	m.Get("/users/:id/view/", userController.View)
	m.Get("/users/:id/edit/", userController.Edit)
	m.Post("/users/:id/save/", userController.Save)

	registerController := new(controllers.RegisterController)
	registerController.SetDi(di)
	m.Get("/register/", registerController.Register)
	m.Post("/register/processRegister/", registerController.ProcessRegister)
	m.Get("/register/activate/:hash/", registerController.ProcessActivate)

	authController := new(controllers.AuthController)
	authController.SetDi(di)
	m.Get("/auth/", authController.Login)
	m.Post("/auth/processLogin/", authController.ProcessLogin)

	updatesController := new(controllers.UpdateLogController)
	updatesController.SetDi(di)
	m.Get("/updates/", updatesController.List)

	m.Run()
}
