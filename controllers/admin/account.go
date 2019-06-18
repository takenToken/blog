package admin

import (
	"blog/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strconv"
	"strings"
)

//login结构
type AccountController struct {
	baseController
}

func (this *AccountController) Login() {
	if this.GetString("dosubmit") == "yes" {
		userName := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := this.GetString("remember")

		if userName != "" && password != "" {
			var user models.User
			user.Username = userName
			err := user.Read("username")
			if err != nil {
				this.Data["errmsg"] = "帐号或密码错误"
				this.TplName = "admin/account_login.html"
				return
			}
			fmt.Println(models.Md5([]byte("123456=abtfg")))
			if user.Password != models.Md5([]byte(password+"="+user.Salt)) {
				this.Data["errmsg"] = "帐号或密码错误"
			} else if user.Active == 0 {
				this.Data["errmsg"] = "该帐号未激活"
			} else {
				user.Logincount += 1
				user.Lastlogin = this.getTime()
				user.Update()

				authkey := models.Md5([]byte(this.getClientIp() + "|" + user.Password))
				if remember == "yes" {
					this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey, 7*86400)
				} else {
					this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey)
				}
				this.Redirect(this.Ctx.Request.Referer(), 302)
			}

		} else if userName == "" {
			this.Data["errmsg"] = "请输入账号"
		} else if password == "" {
			this.Data["errmsg"] = "请输入密码"
		}
	}

	this.TplName = "admin/account_login.html"
}

//登出
func (this *AccountController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.Redirect("/admin/login", 302)
}

//注册
func (this *AccountController) Register() {
	errmsg := make(map[string]string)

	userName := strings.TrimSpace(this.GetString("username"))
	password := strings.TrimSpace(this.GetString("password"))
	password2 := strings.TrimSpace(this.GetString("password2"))
	nickname := strings.TrimSpace(this.GetString("nickname"))
	email := strings.TrimSpace(this.GetString("email"))

	valid := validation.Validation{}
	if v := valid.Required(userName, "username"); !v.Ok {
		errmsg["username"] = "请输入用户名!"
		this.display()
		return
	}

	if password != password2 {
		errmsg["password"] = "两次密码不一致!"
		this.display()
		return
	}

	salt := this.RandStr(5)
	newPassword := models.Md5([]byte(password + "=" + salt))

	user := &models.User{
		Username:  userName,
		Password:  newPassword,
		Salt:      salt,
		Nickname:  nickname,
		Email:     email,
		Lastlogin: this.getTime()}

	if err := user.Insert(); err != nil {
		this.showmsg(err.Error())
	}

	this.Redirect("/admin/user/list", 302)
}

//配置文件
func (this *AccountController) Profile() {
	user := models.User{Id: this.userid}
	if err := user.Read(); err != nil {
		this.showmsg(err.Error())
	}

	this.Data["user"] = user
	this.display()
}
