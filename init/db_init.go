package init

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func dbInit() {
	runmode, _ := web.AppConfig.String("runmode")
	isDev := (runmode == "dev")
	registDatabase()
	if isDev {
		orm.Debug = isDev
	}
}

func registDatabase() {
	//初始化数据库
	dbUser, _ := web.AppConfig.String("mysqluser")
	dbPass, _ := web.AppConfig.String("mysqlpass")
	dbName, _ := web.AppConfig.String("mysqldb")
	dbHost, _ := web.AppConfig.String("mysqlhost")
	dbPort, _ := web.AppConfig.String("mysqlport")
	maxIdleConn, _ := web.AppConfig.Int("db_max_idle_conn")
	maxOpenConn, _ := web.AppConfig.Int("db_max_open_open")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql",
	//	dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=true&loc=Asia%2FShanghai",
	//)
	//初始化casbin
	initCasbin("default", "mysql",
		dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	orm.MaxIdleConnections(maxIdleConn)
	orm.MaxOpenConnections(maxOpenConn)
	orm.DefaultTimeLoc = time.UTC

}
