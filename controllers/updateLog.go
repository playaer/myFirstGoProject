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
	var userId int64 = 16

	logs := di.UpdateLogManager().FindAll(userId)
	r.HTML(200, "updateLog/list", logs)
}