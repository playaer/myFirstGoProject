package controllers
import "net/http"

type User struct {
	BaseController
}

/**
 * Route /users/
 */
func (u *User) list() string {


	return "users list"
}

/**
 * Route /users/:id/view/
 */
func (u *User) view(id int) string {


	return "user view"
}

/**
 * Show edit template
 * Route /users/:id/edit/
 */
func (u *User) edit(id int) string {


	return ""
}

/**
 * Save user
 * Route /users/:id/save/
 */
func (u *User) save(id int, request *http.Request) string {


	return ""
}

