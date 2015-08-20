package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

type AuthController struct {
	BaseController
}

// Show login template
func (self *AuthController) Login(params martini.Params, r render.Render) {
	r.HTML(200, "auth/login", "")
}

// Process login action
func (self *AuthController) ProcessLogin(params martini.Params, req *http.Request, r render.Render) {
	di := *self.di
	email := req.FormValue("Email")
	user := di.UserManager().FindByEmail(email)
	if user == nil {
		// not authenticated
		r.HTML(200, "auth/login", "")
		return
	}

	password := req.FormValue("Password")
	if di.UserManager().CheckPassword(user, password) {
		r.Redirect("/users/")
	} else {
		// not authorized
		r.HTML(200, "auth/login", "")
	}
}