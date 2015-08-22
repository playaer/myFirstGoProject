package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

type UserController struct {
	BaseController
}

/**
 * Routes '/', '/users/'
 */
func (u *UserController) List(r render.Render) {
	di := *u.di
	userManager := di.UserManager()
	all, err := userManager.FindAll()
	if err != nil {
		// 500
	}

	r.HTML(200, "user/list", all)
}

/**
 * Route /users/:id/view/
 */
func (u *UserController) View(params martini.Params, r render.Render) {

	di := *u.di
	userManager := di.UserManager()
	id := params["id"]
	user := userManager.FindById(id)
	if user == nil {
		r.Error(404)
	} else {
		r.HTML(200, "user/view", user)
	}
}

/**
 * Show edit template
 * Route /users/edit/profile/
 */
func (u *UserController) Edit(r render.Render) {
	di := *u.di
	authManager := di.AuthManager()
	if !authManager.IsAuthenticated() {
		r.Error(403)
		return
	}
	userManager := di.UserManager()
	id := string(authManager.CurrentUser().Id)
	user := userManager.FindById(id)
	if user == nil {
		r.Error(404)
	} else {
		r.HTML(200, "user/edit", user)
	}
}

/**
 * Save user
 * Route /users/save/profile/
 */
func (u *UserController) Save(req *http.Request, r render.Render) {
	di := *u.di
	authManager := di.AuthManager()
	if !authManager.IsAuthenticated() {
		r.Error(403)
		return
	}
	userManager := di.UserManager()
	id := string(authManager.CurrentUser().Id)
	user := userManager.FindById(id)
	if user == nil {
		r.Error(404)
		return
	}

	newUser := *user

	newUser.FullName = req.FormValue("FullName")
	newUser.Address = req.FormValue("Address")
	newUser.Phone = req.FormValue("Phone")

	userManager.Update(&newUser)

	di.UpdateLogManager().StoreChanges(user, &newUser)

	r.Redirect("/users/" + id + "/view/")
}

