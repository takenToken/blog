package init

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

func init() {
	//username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8

	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPassword := beego.AppConfig.String("mysqlpassword")
	mysqlHost := beego.AppConfig.String("mysqlhost")
	mysqlDb := beego.AppConfig.String("mysqldb")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqlUser, mysqlPassword, mysqlHost, "3306", mysqlDb)

	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 连接池
	maxIdle := 30
	maxConn := 30
	// set default database
	orm.RegisterDataBase("default", "mysql", dataSource, maxIdle, maxConn)

	//
	orm.Debug = true
}
