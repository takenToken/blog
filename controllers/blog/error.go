package blog

type ErrorController struct {
	baseController
}

//404异常
func (this *ErrorController) Error404() {
	this.Data["content"] = "page not found"
	//this.TplName = "double/404.html"
	this.display("404")
}

//服务器异常
func (this *ErrorController) Error500() {
	this.Data["content"] = "server error"
	this.TplName = "admin/showmsg.html"
}

