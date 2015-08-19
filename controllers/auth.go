package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type AuthController struct {
	BaseController
}

func (self *AuthController) Login(params martini.Params, r render.Render) {
	r.HTML(200, "auth/login", "")
}

func (self *AuthController) ProcessLogin(params martini.Params, r render.Render) {
	r.HTML(200, "auth/login", "")
}