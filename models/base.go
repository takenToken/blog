package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

func GetOptionValue(key string) string {
	var option Option

	if !Cache.IsExist(key) {
		o := orm.NewOrm()
		o.QueryTable(&Option{}).One(option, key)

		Cache.Put(option.Name, option.Value, 1000*time.Hour)
	}

	return option.Value
}

//md5码
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
