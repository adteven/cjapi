// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"cjapi/controllers/admin"
	"cjapi/filter"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	af := filter.AuthFilter{}
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&admin.AuthControler{},
			),
		),
		beego.NSNamespace("/admin",
			beego.NSBefore(af.CheckPermssion),
			beego.NSNamespace("/user",
				beego.NSInclude(
					&admin.SysUserController{},
				),
			),
			beego.NSNamespace("/menu",
				beego.NSInclude(
					&admin.SysMenuController{},
				),
			),
		),
	)
	beego.AddNamespace(ns)
}
