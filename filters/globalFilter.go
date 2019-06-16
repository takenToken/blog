package filters

import (
	"github.com/astaxie/beego/context"
)

func FilterUser(ctx *context.Context) {
	//v := ctx.Input.Session(controllers.SESSION_ID_USERID)
	//if v == nil && (ctx.Request.RequestURI != "/login" && ctx.Request.RequestURI != "/userLogin") {
	//	ctx.Redirect(302, "/login")
	//}
}
