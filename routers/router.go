package routers

import (
	"blog/controllers/admin"
	"github.com/astaxie/beego"
)

//初始化
func init() {


	beego.Router("/admin", &admin.IndexController{}, "*:Index")
	beego.Router("/admin/login", &admin.AccountController{}, "*:Login")
	//beego.Router("/admin/register", &admin.AccountController{}, "*:Register")
	//beego.Router("/admin/logout", &admin.AccountController{}, "*:Logout")
	//beego.Router("/admin/account/profile", &admin.AccountController{}, "*:Profile")
}
