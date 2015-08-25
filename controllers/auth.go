package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/playaer/myFirstGoProject/di"
	"github.com/playaer/myFirstGoProject/utils"
)

type AuthController struct {
	BaseController
}

// Show login template
func (self *AuthController) Login(r render.Render, di *di.DI, templateVars utils.TemplateVars) {
	authManager := di.AuthManager()
	if authManager.IsAuthenticated() {
		r.HTML(403, "error/403", templateVars)
		return
	}
	r.HTML(200, "auth/login", templateVars)
}

// Show login template
func (self *AuthController) LogOut(w http.ResponseWriter, params martini.Params, r render.Render, di *di.DI, templateVars utils.TemplateVars) {
	authManager := di.AuthManager()
	if !authManager.IsAuthenticated() {
		r.HTML(403, "error/403", templateVars)
		return
	}
	authManager.Logout()
	cookie := http.Cookie{Name: "gousertoken", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)
	r.Redirect("/users/")
}

// Process login action
func (self *AuthController) ProcessLogin(w http.ResponseWriter, req *http.Request, r render.Render, di *di.DI) {
	email := req.FormValue("Email")
	user := di.UserManager().FindActiveByEmail(email)
	if user == nil {
		// not authenticated
		// message: Invalid credentials
		r.Redirect("/auth/")
		return
	}

	password := req.FormValue("Password")
	if di.UserManager().CheckPassword(user, password) {
		// message: Hello, {userName}
		authManager := di.AuthManager()
		token := authManager.GenerateToken(user)
		// set cookie
		cookie := http.Cookie{Name: "gousertoken", Value: token, Path: "/"}
		http.SetCookie(w, &cookie)

		r.Redirect("/users/")
	} else {
		// not authorized
		// message: Invalid credentials
		r.Redirect("/auth/")
	}
}