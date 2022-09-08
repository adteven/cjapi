package filter

import (
	"cjapi/common"
	"cjapi/common/jwt"
	"cjapi/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"regexp"
	"strings"
)

const bearerLength = len("Bearer ")

type AuthFilter struct {
}

func (builder *AuthFilter) CheckPermssion(ctx *context.Context) {
	url := ctx.Input.URL()
	ignoreUrls, _ := beego.AppConfig.String("ignore_urls")
	if strings.Contains(ignoreUrls, url) || strings.Contains(url, "/swagger") || strings.Contains(url, "/static") {
		return
	} else {
		mytoken := ctx.Input.Header("Authorization")
		if len(mytoken) < bearerLength {
			ctx.Output.Status = 401
			ctx.Output.JSON(controllers.ErrMsg("header Authorization has not Bearer token", 40001),
				true, true)
			return
		}
		token := strings.TrimSpace(mytoken[bearerLength:])
		usr, err := jwt.ValidateToken(token)
		if err != nil {
			ctx.Output.Status = 401
			ctx.Output.JSON(controllers.ErrMsg(err.Error(), 40001),
				true, true)
			return
		}
		//z这里就是管理员呢的登录验证

		ctx.Input.SetData(common.ContextKeyUserObj, usr)
		//url排除
		if urlExclude(url) {
			return
		}
	}
}

//url排除
func urlExclude(url string) bool {
	//公共路由直接放行
	reg := regexp.MustCompile(`[0-9]+`)
	newUrl := reg.ReplaceAllString(url, "*")
	var ignoreUrls = "/admin/menu/build,/admin/user/center,/admin/user/updatePass,/admin/auth/info," +
		"/admin/auth/logout,/admin/materialgroup/*,/admin/material/*,/shop/product/isFormatAttr/*"
	if strings.Contains(ignoreUrls, newUrl) {
		return true
	}

	return false
}
