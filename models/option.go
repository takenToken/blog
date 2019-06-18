package models

import (
	"github.com/astaxie/beego/orm"
)

//
type Option struct {
	Id    int
	Name  string
	Value string
}

func init() {
	orm.RegisterModelWithPrefix("tb_", &Option{})
}

//新增
func (this *Option) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}

	Cache.ClearAll()
	return nil
}

//修改
func (this *Option) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(this, fields ...)
	if err != nil {
		return err
	}

	Cache.ClearAll()
	return nil
}

//查询
func (this *Option) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}
