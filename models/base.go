package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strings"
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

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

//func GetLatestBlog() []*Post {
//	if !Cache.IsExist("latestblog") {
//		var result []*Post
//		query := new(Post).Query().Filter("status", 0).Filter("urltype", 0)
//		count, _ := query.Count()
//		if count > 0 {
//			query.OrderBy("-posttime").Limit(8).All(&result)
//		}
//		Cache.Put("latestblog", result)
//	}
//	v := Cache.Get("latestblog")
//	return v.([]*Post)
//}

//返回带前缀的表名
//func TableName(str string) string {
//	return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
//}
