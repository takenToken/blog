package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// 字段别名 Name `orm:"column(user_name)"`
type UcUser struct {
	Id       int    `json:"id"`
	UserCode string `json:"userCode"`
	UserName string `json:"userName"`
	Email    string `json:"Email"`
	Password string
	Salt     string
	Status   string
	MobileNo string
}

func init() {
	// register model
	orm.RegisterModel(new(UcUser))
}

//新增
func (u *UcUser) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}



func (u *UcUser) Login(userCode, password string) (err error) {
	o := orm.NewOrm()
	//r := o.Raw("select user_name,password from uc_user where user_name=?", username)
	user := &UcUser{UserCode: userCode}
	rsErr := o.Read(&user, "userCode")
	if rsErr == orm.ErrNoRows {
		fmt.Println("查询不到")
		return rsErr
	} else if rsErr == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return rsErr
	} else {
		fmt.Println(user.Id, user.UserName)
	}

	if user.Password == password {
		return nil
	}

	return nil
}
