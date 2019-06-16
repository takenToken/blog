package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Data["username"] = "liu.qiang"
	this.Data["id"] = "1"
	this.Data["password"] = "123456"
	this.TplName = "user.tpl"
}
