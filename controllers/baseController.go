package controllers

import (
	"github.com/playaer/myFirstGoProject/di"
)

type BaseController struct {
	di *di.DI
}

func (self *BaseController) SetDi(di *di.DI) {
	self.di = di
}
