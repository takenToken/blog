package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

//404异常
func (this *ErrorController) Error404() {
	this.Data["content"] = "page not found"
	this.TplName = "404.tpl"
}

//服务器异常
func (this *ErrorController) Error501() {
	this.Data["content"] = "server error"
	this.TplName = "501.tpl"
}

//db异常
func (this *ErrorController) ErrorDb() {
	this.Data["content"] = "database is now down"
	this.TplName = "dberror.tpl"
}
