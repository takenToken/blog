package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego/orm"
	"regexp"
	"strings"
	"time"
)

type Post struct {
	Id         int64
	User       *User       `orm:"rel(fk);index"`
	Title      string      `orm:"size(100);index"`
	Color      string      `orm:"size(7);index"`
	Urlname    string      `orm:"size(100);index"`
	Urltype    int8        `orm:"index"`
	Content    string      `orm:"type(text)"`
	Tags       string      `orm:"size(100);index"`
	Posttime   time.Time   `orm:"auto_now_add;type(datetime);index"`
	Views      int64       `orm:"index"`
	Status     int8        `orm:"index"`
	Updated    time.Time   `orm:"auto_now_add;type(datetime);index"`
	Istop      int8        `orm:"index"`
	Cover      string      `orm:"size(70);default(/static/upload/default/blog-default-0.png)"`
	Comments   []*Comments `orm:"reverse(many)"`
}

//初始化
func init() {
	orm.RegisterModel(&Post{})
}

func (this *Post) TableName() string{
	return TableName("tb_post")
}

//查询单个
func (this *Post) Read(fields ...string) error {
	if err := orm.NewOrm().Read(this); err != nil {
		return err
	}
	return nil
}

//查询
func (this *Post) Query(fields ...string) orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

//新增
func (this *Post) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

//修改
func (this *Post) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields ...); err != nil {
		return err
	}
	return nil
}

//删除
func (this *Post) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(this, fields ...); err != nil {
		return err
	}
	return nil
}

//带颜色的标题
func (this *Post) ColorTitle() string {
	if this.Color != "" {
		return fmt.Sprintf("<div style=\"color:%s\">%s</div>", this.Color, this.Title)
	} else {
		return this.Title
	}
}

//内容URL
func (this *Post) Link() string {
	if this.Urlname != "" {
		if this.Urltype == 1 {
			return fmt.Sprintf("/%s", RawUrlEncode(this.Urlname))
		}
		return fmt.Sprintf("/article/%s", RawUrlEncode(this.Urlname))
	}
	return fmt.Sprintf("/article/%d", this.Id)
}

//带链接的标签
func (m *Post) TagsLink() string {
	if m.Tags == "" {
		return ""
	}
	var buf bytes.Buffer
	arr := strings.Split(strings.Trim(m.Tags, ","), ",")
	for k, v := range arr {
		if k > 0 {
			buf.WriteString(", ")
		}
		tag := Tag{Name: v}
		if tag.Read("Name") != nil {
			return ""
		}
		buf.WriteString(tag.Link())
	}
	return buf.String()
}

//摘要
func (m *Post) Excerpt() string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	rep := re.ReplaceAllStringFunc(m.Content, strings.ToLower)
	//将hr标签转换为特殊标记
	re, _ = regexp.Compile("------------")
	rep = re.ReplaceAllString(rep, "_markdown_hr_")
	//去除所有尖括号内的HTML代码
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	rep = re.ReplaceAllString(rep, "")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{1,}")
	rep = re.ReplaceAllString(rep, "")
	//如果断定截取的断点可能会存在中文字符，则需要转为rune后再截取，否则可能会截成乱码
	data := []rune(rep)
	if i := strings.Index(rep, "_markdown_hr_"); i != -1 {
		return rep[:i] + "..."
	} else if i = -1; len(data) > 58 {
		return string(data[:62]) + "..."
	}
	return rep
}

func (m *Post) Del_Excerpt() string {
	if i := strings.Index(m.Content, "------------"); i != -1 {
		x := len("------------")
		return m.Content[i+x:]
	}
	return m.Content
}

func (post *Post) GetPreAndNext() (pre, next *Post) {
	pre = new(Post)
	next = new(Post)
	err := new(Post).Query().Filter("id__lt", post.Id).Filter("status", 0).Filter("urltype", 0).OrderBy("-id").Limit(1).One(pre)
	if err == orm.ErrNoRows {
		pre = nil
	}
	err = new(Post).Query().Filter("id__gt", post.Id).Filter("status", 0).Filter("urltype", 0).Limit(1).One(next)
	if err == orm.ErrNoRows {
		next = nil
	}
	return
}

