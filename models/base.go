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

//read option to cache
func GetOptionValue(key string) (string, error) {
	if !Cache.IsExist(key) {
		var option Option
		o := orm.NewOrm()
		err := o.QueryTable(&Option{}).One(option, key)
		if err != nil {
			return "", err
		}

		err = Cache.Put(option.Name, option.Value, 1000*time.Hour)
		if err != nil {
			return "", err
		}

		return option.Value, nil
	}
	return Cache.Get(key).(string), nil
}

//md5码
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//
func RawUrlEncode(str string) string {
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

func GetOptions() (map[string]string) {
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
			return make(map[string]string, 0)
		}

		err = Cache.Put(key, optionStr, 1000*time.Hour)
		if err != nil {
			logs.Error(err)
			return make(map[string]string, 0)
		}
		return optionMap
	}

	var opMap map[string]string
	//类型
	opByte := Cache.Get(key).([]byte)
	if opByte != nil {
		//反序列化
		err := json.Unmarshal(opByte, &opMap)
		if err != nil {
			logs.Error(err)
			return make(map[string]string, 0)
		}
	}
	return opMap

}

//最新4篇文章
func GetLastPosts() []*map[string]string {
	post := &Post{}
	var posts []*Post
	_, err := post.Query().Limit(4).OrderBy("-posttime").All(&posts)
	if err != nil {
		logs.Error(err)
		//err create null
		return make([]*map[string]string, 0)
	}

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
	_, err := post.Query().Limit(7).OrderBy("-views").All(&posts)
	if err != nil {
		logs.Error(err)
		return make([]*map[string]string, 0)
	}

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
	_, err := comments.Query().Limit(7).Filter("is_removed", 0).OrderBy("-submittime").RelatedSel("user").All(&commentsList)
	if err != nil {
		logs.Error(err)
		return nil
	}

	return commentsList
}
