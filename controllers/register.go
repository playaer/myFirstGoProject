package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/playaer/myFirstGoProject/di"
	"github.com/playaer/myFirstGoProject/utils"
)

type RegisterController struct {
	BaseController
}

// Register Action, show register template form
func (self *RegisterController) Register(params martini.Params, r render.Render, di *di.DI, templateVars utils.TemplateVars) {
	authManager := di.AuthManager()
	if authManager.IsAuthenticated() {
		r.HTML(403, "error/403", templateVars)
		return
	}
	r.HTML(200, "register/register", templateVars)
}

// Process register: check form, create new inactive user and send activation email
func (self *RegisterController) ProcessRegister(params martini.Params, req *http.Request, r render.Render, di *di.DI) {
	userManager := di.UserManager()
	user := userManager.NewUser()

	// need validate

	user.Email = req.FormValue("Email")
	rawPassword := req.FormValue("Password")
	user.Password = userManager.CryptPassword(rawPassword)
	user.FullName = req.FormValue("FullName")
	user.Address = req.FormValue("Address")
	user.Phone = req.FormValue("Phone")
	user.IsActive = false
	user.Hash = userManager.GenerateHash(user.Email + user.FullName)

	userManager.Create(user)

	// send email
	mailer := di.Mailer()
	go mailer.Send(mailer.BuildRegistrationMail(user))

	// message: "Activation link was sent to your email."

	r.Redirect("/users/")
}

// Check activation link and activate user
func (self *RegisterController) ProcessActivate(params martini.Params, req *http.Request, r render.Render, di *di.DI, templateVars utils.TemplateVars) {

	userManager := di.UserManager()
	user := userManager.FindInActiveByHash(params["hash"])

	if user == nil {
		r.HTML(404, "error/404", templateVars)
		return
	}

	user.IsActive = true
	user.Hash = ""

	userManager.Update(user)

	// message "Activation complete. How you can enter on site."

	r.Redirect("/auth/")
}
