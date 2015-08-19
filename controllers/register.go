package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type RegisterController struct {
	BaseController
}

func (self *RegisterController) Register(params martini.Params, r render.Render) {
	r.HTML(200, "register/register", "")
}

func (self *RegisterController) ProcessRegister(params martini.Params, r render.Render) {
	r.HTML(200, "register/register", "")
}
