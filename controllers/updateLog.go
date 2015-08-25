package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/playaer/myFirstGoProject/di"
	"github.com/playaer/myFirstGoProject/utils"
)

type UpdateLogController struct {
	BaseController
}

// Show login template
func (self *UpdateLogController) List(params martini.Params, r render.Render, di *di.DI, templateVars utils.TemplateVars) {
	authManager := di.AuthManager()
	if !authManager.IsAuthenticated() {
		r.HTML(403, "error/403", templateVars)
		return
	}
	currentUser := authManager.CurrentUser()

	userId := currentUser.Id

	logs := di.UpdateLogManager().FindAll(userId)
	templateVars.SetData(logs)
	r.HTML(200, "updateLog/list", templateVars)
}