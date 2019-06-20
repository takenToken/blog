package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

//标签表
type Tag struct {
	Id    int64
	Name  string `orm:"size(20);index"`
	Count int64  `orm:"index"`
}

//初始化
func init() {
	orm.RegisterModelWithPrefix("tb_", &Tag{})
}

//查询单个
func (this *Tag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(this); err != nil {
		return err
	}
	return nil
}

//查询
func (this *Tag) Query(fields ...string) orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

//新增
func (this *Tag) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

//修改
func (this *Tag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields ...); err != nil {
		return err
	}
	return nil
}

//删除
func (this *Tag) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(this, fields ...); err != nil {
		return err
	}
	return nil
}

//标签连接
func (this *Tag) Link() string {
	return fmt.Sprintf("<a class=\"category\" href=\"/category/%d\">%s</a>", this.Id, this.Name)
}
