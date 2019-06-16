package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/session/redis"
)

const SESSION_ID_USERID = "UserID"

type Login struct {
	beego.Controller
}

func (c *Login) Get() {
	c.TplName = "login.tpl"
}

//登录
func (c *Login) Login() {

	userName := c.GetString("username")
	password := c.GetString("password")

	if userName == "" && password == "" {
		c.Redirect("/login", 302)
	}

	ucUser := models.UcUser{}
	err := ucUser.Login(userName, password)
	if (err != nil) {
		c.Ctx.WriteString("Not Found UserName:" + userName)
		return
	}

	logs.Info("session UserID:", userName)
	c.SetSession(SESSION_ID_USERID, userName)

	c.Redirect("/", 302)
}

//退出登录，删除sessionId
func (c *Login) Logout() {
	userId := c.GetSession(SESSION_ID_USERID)
	if userId != nil {
		c.DelSession(SESSION_ID_USERID)
		logs.Info("delete session UserID:", userId)
	}
	c.TplName = "logout.tpl"
}
