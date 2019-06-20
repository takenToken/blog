package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

// 字段别名 Name `orm:"column(user_name)"`
//用户表模型
type User struct {
	Id         int64
	Username   string    `json:"username" orm:"unique;size(15);index"`
	Password   string    `orm:"size(32)"`
	Nickname   string    `json:"nickname" orm:"size(15);index"`
	Salt       string    `orm:"size(20)"`
	Email      string    `json:"email" orm:"size(50);index"`
	Lastlogin  time.Time `json:"lastlogin" orm:"auto_now_add;type(datetime);index"`
	Logincount int64     `orm:"index"`
	Lastip     string    `json:"lastip" orm:"size(32);index"`
	Authkey    string    `orm:"size(10)"`
	Active     int8
	Permission string `orm:"size(100);index"`
	Avator     string `json:"avator" orm:"size(150);default(/static/upload/default/user-default-60x60.png)"`
	Upcount    int64
	//Post       []*Post       `orm:"reverse(many)"`
	//Comments   []*Comments   `orm:"reverse(many)"`
}

func init() {
	// register model
	orm.RegisterModelWithPrefix("tb_", new(User))
}

//新增
func (u *User) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

//fields更新哪里字段
func (u *User) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(u, fields...)
	if err != nil {
		return err
	}
	return nil
}

//删除
func (u *User) Delete(fields ...string) (int64, error) {
	count, err := orm.NewOrm().Delete(u, fields ...)
	return count, err
}

//fields代表哪些字段作为查询条件
func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (u *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(u)
}
