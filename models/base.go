package models

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func GetOptionValue(key string) string {
	if !Cache.IsExist(key) {
		var option Option
		o := orm.NewOrm()
		o.QueryTable(&Option{}).One(option, key)
		Cache.Put(option.Name, option.Value, 1000*time.Hour)
		return option.Value
	} else {
		return Cache.Get(key).(string)
	}
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
func TableName(str string) string {
	return fmt.Sprintf("%s", str)
}

//https://studygolang.com/articles/12348?fr=sidebar
//通过反射复制struct属性
func Copy(src interface{}, dst interface{}) error {

	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() != reflect.Ptr {
		return errors.New("dst isn't a pointer to struct")
	}

	dstElem := dstValue.Elem()
	if dstElem.Kind() != reflect.Struct {
		return errors.New("pointer doesn't point to struct")
	}

	//pointer.Elem()去获取所指向的Value，注意一定要是指针
	srcValue := reflect.ValueOf(src)
	srcElem := srcValue.Elem()

	for i := 0; i < srcElem.Type().NumField(); i++ {

		sf := srcElem.Type().Field(i)
		sv := srcElem.FieldByName(sf.Name)

		if dv := dstElem.FieldByName(sf.Name); dv.IsValid() && dv.CanSet() {
			dv.Set(sv)
		}
	}

	return nil
}

func GetOptions() map[string]string {
	key := "options"
	if !Cache.IsExist(key) {
		var options []*Option
		count, err := orm.NewOrm().QueryTable(&Option{}).All(&options);
		if err != nil {
			logs.Error(err)
		}
		//转换成map
		optionMap := make(map[string]string, count)
		for _, op := range options {
			optionMap[op.Name] = op.Value
		}

		//存入缓存
		optionStr, err := json.Marshal(&optionMap)
		if err != nil {
			logs.Error(err)
		}

		Cache.Put(key, optionStr, 1000*time.Hour)
		return optionMap
	} else {
		var opMap map[string]string
		//类型
		opByte := Cache.Get(key).([]byte)
		if opByte != nil {
			//反序列化
			json.Unmarshal(opByte, &opMap)
		}
		return opMap
	}

}

//最新4篇文章
func GetLastPosts() []*map[string]string {
	post := &Post{}
	var posts []*Post
	post.Query().Limit(4).OrderBy("-posttime").All(&posts)

	ps := make([]*map[string]string, len(posts))
	for index, v := range posts {
		amMap := make(map[string]string, 4)
		amMap["id"] = strconv.FormatInt(v.Id, 10)
		amMap["text"] = v.Content
		ps[index] = &amMap
	}

	return ps
}

//最新7篇文章排名
func GetLastPostRankings() []*map[string]string {
	post := &Post{}
	var posts []*Post
	post.Query().Limit(7).OrderBy("-views").All(&posts)

	ps := make([]*map[string]string, len(posts))
	for index, v := range posts {
		amMap := make(map[string]string, 4)
		amMap["id"] = strconv.FormatInt(v.Id, 10)
		amMap["text"] = v.Content
		ps[index] = &amMap
	}

	return ps
}

//获取最新评论
func GetLastComments() []*Comments {
	comments := &Comments{}
	var commentsList []*Comments
	comments.Query().Limit(7).Filter("is_removed", 0).OrderBy("-submittime").RelatedSel("user").All(&commentsList)

	return commentsList
}
