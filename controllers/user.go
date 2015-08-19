package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	models "github.com/playaer/myFirstGoProject/models"
)

type UserController struct {
	BaseController
}



func getUsers() map[string]interface{} {
	users := make(map[string]interface{})

	users["0"] = models.User{"0", "First User", "Address", "227-84-61"}
	users["1"] = models.User{"1", "Second User", "Address", "227-84-61"}
	users["2"] = models.User{"2", "Third User", "Address", "227-84-61"}
	users["3"] = models.User{"3", "Forth User", "Address", "227-84-61"}

	return users
}

/**
 * Routes '/', '/users/'
 */
func (u *UserController) List(r render.Render) {

	r.HTML(200, "user/list", getUsers())
}

/**
 * Route /users/:id/view/
 */
func (u *UserController) View(params martini.Params, r render.Render) {

	id := params["id"]
	users := getUsers()
	r.HTML(200, "user/view", users[id])
}

/**
 * Show edit template
 * Route /users/:id/edit/
 */
func (u *UserController) Edit(params martini.Params, r render.Render) {
	id := params["id"]
	users := getUsers()
	r.HTML(200, "user/edit", users[id])
}

/**
 * Save user
 * Route /users/:id/save/
 */
func (u *UserController) Save(params martini.Params, r render.Render) {
	id := params["id"]

	r.Redirect("/users/" + id + "/view/")
}

