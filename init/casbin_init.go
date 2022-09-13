package init

import (
	"cjapi/runtime"
	"github.com/beego/beego/v2/core/logs"
	beegoormadapter "github.com/casbin/beego-orm-adapter/v3"
	"github.com/casbin/casbin/v2"
)

func initCasbin(aliasName, driverName, dataSourceName string) {
	// Initialize a Beego ORM adapter and use it in a Casbin enforcer:
	a, err := beegoormadapter.NewAdapter(aliasName, driverName, dataSourceName) // Your driver and data source.
	if err != nil {
		logs.Error(err)
	}

	e, err := casbin.NewSyncedEnforcer("conf/rbac_model.conf", a)
	if err != nil {
		logs.Error(err)
	}

	runtime.Runtime.SetCasbin("default", e)
}
