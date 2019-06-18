package admin

import (
	"blog/models"
	"github.com/astaxie/beego/validation"
	"strings"
)

type UserController struct {
	baseController
}

//分页查询
func (this *UserController) List() {
	var page int64
	var pageSize int64 = 10
	var list []*models.User
	var user models.User
	if page, _ = this.GetInt64("page"); page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	count, _ := user.Query().Count()
	if count > 0 {
		user.Query().OrderBy("-id").Limit(pageSize, offset).All(&list)
	}

	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(page, count, pageSize, "/admin/user/list?page=%d").ToString()
	this.display()
}

//新增
func (this *UserController) Add() {
	this.display()
}

//编辑
func (this *UserController) Edit() {
	id, _ := this.GetInt64("id")
	user := models.User{Id: id}
	if err := user.Read(); err != nil {
		this.showmsg("用户不存在")
	}

	errmsg := make(map[string]string)

	if this.Ctx.Request.Method == "POST" {
		password := strings.TrimSpace(this.GetString("password"))
		password2 := strings.TrimSpace(this.GetString("password2"))
		email := strings.TrimSpace(this.GetString("email"))
		avator := strings.TrimSpace(this.GetString("avator"))
		permissionlist := strings.TrimSpace(
			this.GetString("permission1") + "|" +
				this.GetString("permission2") + "|" +
				this.GetString("permission3") + "|" +
				this.GetString("permission4") + "|" +
				this.GetString("permission5") + "|" +
				this.GetString("permission6") + "|" +
				this.GetString("permission7"))
		active, _ := this.GetInt64("active")
		valid := validation.Validation{}

		if password != "" {
			if v := valid.Required(password2, "password2"); !v.Ok {
				errmsg["password2"] = "请再次输入密码"
			} else if password != password2 {
				errmsg["password2"] = "两次输入的密码不一致"
			} else {
				//密码加随机因子salt
				user.Password = models.Md5([]byte(password + "=" + user.Salt))
			}
		}
		if v := valid.Required(email, "email"); !v.Ok {
			errmsg["email"] = "请输入email地址"
		} else if v := valid.Email(email, "email"); !v.Ok {
			errmsg["email"] = "Email无效"
		} else {
			user.Email = email
		}

		if active > 0 {
			user.Active = 1
		} else {
			user.Active = 0
		}

		if len(errmsg) == 0 {
			if user.Id != 1 {
				user.Permission = permissionlist
			}
			user.Avator = avator
			user.Update()
			this.Redirect("/admin/user/list", 302)
		}

	}

	this.Data["errmsg"] = errmsg
	this.Data["user"] = user
	this.display()
}

//删除
func (this *UserController) Delete() {
	this.display()
}
