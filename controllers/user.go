package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	models "github.com/playaer/myFirstGoProject/models"
	"github.com/playaer/myFirstGoProject/utils"
)

type UserController struct {
	BaseController
}

func getUsers() map[string]models.User {
	users := map[string]models.User{
//		"1": models.User{"1", "First User", "Address", "227-84-61",},
//		"2": models.User{"2", "Second User", "Address", "227-84-61",},
//		"3": models.User{"3", "Third User", "Address", "227-84-61",},
//		"4": models.User{"4", "Forth User", "Address", "227-84-61",},
	}

	return users
}

/**
 * Routes '/', '/users/'
 */
func (u *UserController) List(r render.Render) {


	di := utils.New()
	userManager := di.UserManager()
	all, _ := userManager.FindAll()
	if (len(all) == 0) {
		all = []*models.User{
//			models.User{1, "First User", "Address", "227-84-61",},
//			models.User{2, "Second User", "Address", "227-84-61",},
//			models.User{3, "Third User", "Address", "227-84-61",},
//			models.User{4, "Forth User", "Address", "227-84-61",},
		}
	}
	r.HTML(200, "user/list", all)
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

