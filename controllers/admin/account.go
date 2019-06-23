package admin

import (
	"blog/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"os"
	"regexp"
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

			if user.Password != models.Md5([]byte(password+"="+user.Salt)) {
				this.Data["errmsg"] = "帐号或密码错误"
			} else if user.Active == 0 {
				this.Data["errmsg"] = "该帐号未激活"
			} else {
				user.Logincount += 1
				user.Lastlogin = this.getTime()
				user.Update()

				authKey := models.Md5([]byte("|" + user.Password))
				if remember == "yes" {
					this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authKey, 7*86400)
				} else {
					this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authKey)
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
	lastavator := user.Avator
	if this.Ctx.Request.Method == "POST" {
		errmsg := make(map[string]string)
		password := strings.TrimSpace(this.GetString("password"))
		newpassword := strings.TrimSpace(this.GetString("newpassword"))
		newpassword2 := strings.TrimSpace(this.GetString("newpassword2"))
		avator := strings.TrimSpace(this.GetString("avator"))
		nickname := strings.TrimSpace(this.GetString("nickname"))
		if avator == "" {
			avator = "/static/upload/default/user-default-60x60.png"
		}
		if avator != lastavator {
			models.Cache.Delete("newcomments")
			if user.Upcount > 0 {
				user.Upcount--
			}
			if !this.Isdefaultsrc(lastavator) {
				os.Remove("." + lastavator)
			}
		}
		updated := false
		valid := validation.Validation{}
		if v := valid.Required(nickname, "nickname"); !v.Ok {
			errmsg["nickname"] = "请输入昵称"
		} else if v := valid.MaxSize(nickname, 15, "nickname"); !v.Ok {
			errmsg["nickname"] = "昵称长度不能大于15个字符"
		}
		var user1 models.User
		err := user1.Query().Filter("nickname", nickname).One(&user1)
		if err == nil && user1.Id != user.Id {
			errmsg["nickname"] = fmt.Sprintf("昵称:%s 已被使用", nickname)
		}
		if nickname != user.Nickname && len(errmsg) == 0 {
			user.Nickname = nickname
			user.Update("nickname")
			updated = true
		}
		if avator != user.Avator && len(errmsg) == 0 {
			user.Avator = avator
			user.Update("avator", "upcount")
			updated = true
		}
		if newpassword != "" {
			if password == "" || models.Md5([]byte(password+"="+user.Salt)) != user.Password {
				errmsg["password"] = "当前密码错误"
			} else if len(newpassword) < 6 {
				errmsg["newpassword"] = "密码长度不能少于6个字符"
			} else if newpassword != newpassword2 {
				errmsg["newpassword2"] = "两次输入的密码不一致"
			}
			if len(errmsg) == 0 {
				user.Password = models.Md5([]byte(newpassword + "=" + user.Salt))
				user.Update("password")
				updated = true
			}
		}
		this.Data["updated"] = updated
		this.Data["errmsg"] = errmsg
	}
	this.Data["user"] = user
	this.display()
}

func checkUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z]([a-zA-Z0-9-_]{4,14})+$", username); !ok {
		return false
	}
	return true
}

func checkPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9[:punct:]]{4,19}$", password); !ok {
		return false
	}
	return true
}
