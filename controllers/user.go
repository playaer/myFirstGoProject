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
 * Route /users/:id/edit/
 */
func (u *UserController) Edit(params martini.Params, r render.Render) {
	di := *u.di
	userManager := di.UserManager()
	id := params["id"]
	user := userManager.FindById(id)
	if user == nil {
		r.Error(404)
	} else {
		r.HTML(200, "user/edit", user)
	}
}

/**
 * Save user
 * Route /users/:id/save/
 */
func (u *UserController) Save(params martini.Params, req *http.Request, r render.Render) {
	di := *u.di
	userManager := di.UserManager()
	user := userManager.FindById(params["id"])
	if user == nil {
		r.Error(404)
		return
	}

	newUser := *user

	di.UpdateLogManager().StoreChanges(user, &newUser)

	newUser.FullName = req.FormValue("FullName")
	newUser.Address = req.FormValue("Address")
	newUser.Phone = req.FormValue("Phone")

	userManager.Update(&newUser)
	r.Redirect("/users/" + params["id"] + "/view/")
}

