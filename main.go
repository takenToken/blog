package main

import (
	"blog/controllers"
	_ "blog/init"
	_ "blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
)

func initLogger() {
	var consoleconfig=`{"level":7}`
	logs.SetLogger(logs.AdapterConsole,consoleconfig)
	var fileconfig = `{"filename":"log/log.log","level":7}`
	logs.SetLogger(logs.AdapterFile,fileconfig)
	//var mailconfig = `{"username":"beegotest@gmail.com",
	//					"password":"xxxxxxxx",
	//					"host":"smtp.gmail.com:587",
	//					"sendTos":["xiemengjun@gmail.com"]}`
	//logs.SetLogger(logs.AdapterMail,mailconfig)
	logs.Async(1e3)
}

//添加自定义模板函数
func hasPermission(permissionlist map[string]int, value int) (out bool) {
	for _, id := range permissionlist {
		if value == id {
			return true
		}
	}
	return false
}

func hasPermissionstr(permissionlist string, value int) (out bool) {
	for _, id := range strings.Split(permissionlist, "|") {
		if id == strconv.Itoa(value) {
			return true
		}
	}
	return false
}

func main() {
	//异常处理
	beego.ErrorController(&controllers.ErrorController{})
	//过滤器
	//beego.InsertFilter("/*", beego.BeforeRouter, filters.FilterUser)
	beego.Run()
}
