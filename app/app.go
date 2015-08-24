package app

import (
	"github.com/go-martini/martini"
	controllers "github.com/playaer/myFirstGoProject/controllers"
	"github.com/martini-contrib/render"
	"github.com/playaer/myFirstGoProject/di"
	"html/template"
	"time"
	"net/http"
	"github.com/playaer/myFirstGoProject/managers"
	"github.com/playaer/myFirstGoProject/utils"
)

func Run() {
	m := martini.Classic()

	m.Use(func(c martini.Context) {
		d := di.New()
		c.Map(d)
		c.Map(d.AuthManager())
	})

	m.Use(func(req *http.Request, manager *managers.AuthManager) {
		tokenCookie, err := req.Cookie("gousertoken")
		if err != nil {
			return
		}
		user, err := manager.FindActiveByToken(tokenCookie.Value)
		if err != nil || user == nil {
			return
		}
		manager.Auth(user)
		utils.Debug("user authenticated")
	})

	m.Use(render.Renderer(render.Options{
		Layout: "layout",
		Directory: "templates",
		Funcs: []template.FuncMap{
			{
				"formatTime": func(t *time.Time) string {
					return t.Format(time.Stamp)
				},
				"isAuthenticated": func() bool {
					return false//manager.IsAuthenticated()
				},
			},
		},
	}))

	m.Use(func(c martini.Context, di *di.DI){
		c.Next()
		di.Db().Close()
		utils.Debug("DB connection closed")
	})

	userController := new(controllers.UserController)
	m.Get("/", userController.List)
	m.Get("/users/", userController.List)
	m.Get("/users/:id/view/", userController.View)
	m.Get("/users/edit/profile/", userController.Edit)
	m.Post("/users/save/profile/", userController.Save)

	registerController := new(controllers.RegisterController)
	m.Get("/register/", registerController.Register)
	m.Post("/register/processRegister/", registerController.ProcessRegister)
	m.Get("/register/activate/:hash/", registerController.ProcessActivate)

	authController := new(controllers.AuthController)
	m.Get("/auth/", authController.Login)
	m.Get("/auth/logout/", authController.LogOut)
	m.Post("/auth/processLogin/", authController.ProcessLogin)

	updatesController := new(controllers.UpdateLogController)
	m.Get("/updates/", updatesController.List)

	m.Run()
}
