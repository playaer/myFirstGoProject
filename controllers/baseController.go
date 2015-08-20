package controllers

import (
	"github.com/martini-contrib/render"
	"github.com/playaer/myFirstGoProject/di"
)

type BaseController struct {
	render render.Render
	di *di.DI
}

func (self *BaseController) SetDi(di *di.DI) {
	self.di = di
}


func (self *BaseController) Render(templateName string, data ...map[string]string) {
	if (data == nil) {
		data := make(map[string]string)
		data["a"] = "a"

		self.render.HTML(200, templateName, data)
	} else {
		self.render.HTML(200, templateName, data)
	}
}
