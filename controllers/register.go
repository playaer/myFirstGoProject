package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

type RegisterController struct {
	BaseController
}

func (self *RegisterController) Register(params martini.Params, r render.Render) {
	r.HTML(200, "register/register", nil)
}

func (self *RegisterController) ProcessRegister(params martini.Params, req *http.Request, r render.Render) {

	di := *self.di
	userManager := di.UserManager()
	user := userManager.NewUser()

	// need validate

	user.Email = req.FormValue("Email")
	rawPassword := req.FormValue("Password")
	user.Password = userManager.CryptPassword(rawPassword)
	user.FullName = req.FormValue("FullName")
	user.Address = req.FormValue("Address")
	user.Phone = req.FormValue("Phone")

	userManager.Create(user)
	r.Redirect("/auth/")
}
