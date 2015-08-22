package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type UpdateLogController struct {
	BaseController
}

// Show login template
func (self *UpdateLogController) List(params martini.Params, r render.Render) {
	di := *self.di
	authManager := di.AuthManager()
	if !authManager.IsAuthenticated() {
		r.Error(403)
		return
	}
	currentUser := authManager.CurrentUser()

	userId := currentUser.Id

	logs := di.UpdateLogManager().FindAll(userId)
	r.HTML(200, "updateLog/list", logs)
}