package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/playaer/myFirstGoProject/di"
)

type UpdateLogController struct {
	BaseController
}

// Show login template
func (self *UpdateLogController) List(params martini.Params, r render.Render, di *di.DI) {
	authManager := di.AuthManager()
	if !authManager.IsAuthenticated() {
		r.HTML(403, "error/403", nil)
		return
	}
	currentUser := authManager.CurrentUser()

	userId := currentUser.Id

	logs := di.UpdateLogManager().FindAll(userId)
	r.HTML(200, "updateLog/list", logs)
}