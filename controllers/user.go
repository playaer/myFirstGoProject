package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/playaer/myFirstGoProject/utils"
	"github.com/playaer/myFirstGoProject/di"
	"strconv"
)

type UserController struct {
	BaseController
}

/**
 * Routes '/', '/users/'
 */
func (u *UserController) List(r render.Render, di *di.DI) {
	userManager := di.UserManager()
	all := userManager.FindAll()
	r.HTML(200, "user/list", all)
}

/**
 * Route /users/:id/view/
 */
func (u *UserController) View(params martini.Params, r render.Render, di *di.DI) {
	userManager := di.UserManager()
	id := params["id"]
	user := userManager.FindById(id)
	if user == nil {
		r.HTML(404, "error/404", nil)
	} else {
		r.HTML(200, "user/view", user)
	}
}

/**
 * Show edit template
 * Route /users/edit/profile/
 */
func (u *UserController) Edit(r render.Render, di *di.DI) {
	authManager := di.AuthManager()
	utils.Debug(authManager)
	if !authManager.IsAuthenticated() {
		r.HTML(403, "error/403", nil)
		return
	}
	currentUser := authManager.CurrentUser()

	r.HTML(200, "user/edit", currentUser)
}

/**
 * Save user
 * Route /users/save/profile/
 */
func (u *UserController) Save(req *http.Request, r render.Render, di *di.DI) {
	authManager := di.AuthManager()
	if !authManager.IsAuthenticated() {
		r.HTML(403, "error/403", nil)
		return
	}
	userManager := di.UserManager()
	currentUser := authManager.CurrentUser()

	// clone user
	newUser := *currentUser
	newUser.FullName = req.FormValue("FullName")
	newUser.Address = req.FormValue("Address")
	newUser.Phone = req.FormValue("Phone")

	userManager.Update(&newUser)

	// store user updates to log
	di.UpdateLogManager().StoreChanges(currentUser, &newUser)

	// redirect
	strId := strconv.FormatInt(currentUser.Id, 10)
	r.Redirect("/users/" + strId + "/view/")
}

