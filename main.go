package main

import (
	"cjapi/controllers"
	_ "cjapi/init"
	_ "cjapi/routers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
)

func main() {

	//InsertFilter是提供一个过滤函数

	//var FilterCors = func(ctx *context.Context) {
	//	logs.Info(ctx.Input.Header("Origin"))
	//	cors.Allow(&cors.Options{
	//		AllowAllOrigins:  true,
	//		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-Token", "X-Requested-With", "Access-Control-Allow-Credentials"},
	//		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//		AllowCredentials: true,
	//	})
	//}

	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowAllOrigins:  true,
	//	AllowOrigins:     []string{"http://localhost:8084"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-Token", "X-Requested-With", "Access-Control-Allow-Credentials"},
	//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	AllowCredentials: true,
	//}))
	//beego.InsertFilter("/*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"*"},
	//	AllowHeaders:     []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "Access-Control-Allow-Origin"},
	//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	AllowCredentials: true,
	//}))

	var corsFunc = func(ctx *context.Context) {
		var success = []byte("SUPPORT OPTIONS")
		origin := ctx.Input.Header("Origin")
		ctx.Output.Header("Access-Control-Allow-Methods", "OPTIONS,DELETE,POST,GET,PUT,PATCH")
		ctx.Output.Header("Access-Control-Max-Age", "3600")
		ctx.Output.Header("Access-Control-Allow-Headers", "Authorization,X-Custom-Header,accept,Content-Type,Access-Token")
		ctx.Output.Header("Access-Control-Allow-Credentials", "true")
		ctx.Output.Header("Access-Control-Allow-Origin", origin)
		if ctx.Input.Method() == http.MethodOptions {
			// options请求，返回200
			ctx.Output.SetStatus(http.StatusOK)
			_ = ctx.Output.Body(success)
		}
	}
	beego.InsertFilter("/*", beego.BeforeStatic, corsFunc)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
