package controllers

import (
	"github.com/astaxie/beego"
)

type SubPreparer interface {
	SubPrepare()
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	if app, ok := this.AppController.(SubPreparer); ok {
		app.SubPrepare()
	}
}