package vo

import "time"

type PostVO struct {
	Id       int64
	User_id  int64
	UserName string
	Title    string
	Color    string
	Urlname  string
	Urltype  int8
	Content  string
	Tags     string
	Posttime time.Time
	Views    int64
	Status   int8
	Updated  time.Time
	Istop    int8
	Cover    string
}
