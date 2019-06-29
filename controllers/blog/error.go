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
	this.Layout = "admin/layout.html"
	this.TplName = "admin/showmsg.html"
}

//显示错误提示
func (this *baseController) showmsg(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}

	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout = "admin/layout.html"
	this.TplName = "admin/" + "showmsg.html"
	this.Render()
	this.StopRun()
}
