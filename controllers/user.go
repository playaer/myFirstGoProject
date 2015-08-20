package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/playaer/myFirstGoProject/di"
	"strconv"
	"net/http"
)

type UserController struct {
	BaseController
}

/**
 * Routes '/', '/users/'
 */
func (u *UserController) List(r render.Render) {


	di := di.New()
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

	di := di.New()
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
	di := di.New()
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
	di := di.New()
	userManager := di.UserManager()
	user := userManager.FindById(params["id"])
	if user == nil {
		r.Error(404)
		return
	}

	// need validate

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		r.Error(500)
	}

	user.Id = id
	user.FullName = req.FormValue("FullName")
	user.Address = req.FormValue("Address")
	user.Phone = req.FormValue("Phone")

	userManager.Update(user)
	r.Redirect("/users/" + params["id"] + "/view/")
}

