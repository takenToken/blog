package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
)

//Cache
var Cache cache.Cache

//初始化
func init() {
	conn := beego.AppConfig.String("redisConfig")
	key := beego.AppConfig.String("redisGlobKey")
	dbNum := beego.AppConfig.String("redisDbNum")
	password := beego.AppConfig.String("redisPassword")
	redisConfig := fmt.Sprintf("{\"key\":\"%s\",\"conn\":\"%s\",\"dbNum\":\"%s\",\"password\":\"%s\"}", key, conn, dbNum, password)
	bm, err := cache.NewCache("redis", redisConfig)
	if err != nil {
		logs.Error(err)
	}
	Cache = bm
}
