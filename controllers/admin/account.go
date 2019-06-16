package admin

//login结构
type AccountController struct {
	baseController
}

func (this *AccountController) Login(){
	if this.GetString("dosubmit") == "yes"{

	}

	this.TplName = "admin/account_login.html"
}