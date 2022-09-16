package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["cjapi/controllers/admin:AuthControler"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:AuthControler"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:AuthControler"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:AuthControler"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysDeptController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysDeptController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysDeptController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysDeptController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysDeptController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysDeptController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysDeptController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysDeptController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysDeptController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysDeptController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysLogController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysLogController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysLogController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysLogController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysLogController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysLogController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysLogController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysLogController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysLogController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysLogController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"],
        beego.ControllerComments{
            Method: "Nav",
            Router: `/nav`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysMenuController"],
        beego.ControllerComments{
            Method: "GetTree",
            Router: `/tree`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysRoleController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysRoleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysRoleController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysRoleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysRoleController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysRoleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysRoleController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysRoleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysRoleController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysRoleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:SysUserController"],
        beego.ControllerComments{
            Method: "Welcome",
            Router: `/welcome`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"],
        beego.ControllerComments{
            Method: "GetIndex",
            Router: `/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:TbJishuiAreaController"],
        beego.ControllerComments{
            Method: "GetPage",
            Router: `/page`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:XzAreaDetailController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:XzAreaDetailController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:XzAreaDetailController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:XzAreaDetailController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:XzAreaDetailController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:XzAreaDetailController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:XzAreaDetailController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:XzAreaDetailController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/admin:XzAreaDetailController"] = append(beego.GlobalControllerRouter["cjapi/controllers/admin:XzAreaDetailController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers/common:CommonController"] = append(beego.GlobalControllerRouter["cjapi/controllers/common:CommonController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["cjapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["cjapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["cjapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["cjapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["cjapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:UserController"] = append(beego.GlobalControllerRouter["cjapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:UserController"] = append(beego.GlobalControllerRouter["cjapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:UserController"] = append(beego.GlobalControllerRouter["cjapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:UserController"] = append(beego.GlobalControllerRouter["cjapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:UserController"] = append(beego.GlobalControllerRouter["cjapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:UserController"] = append(beego.GlobalControllerRouter["cjapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cjapi/controllers:UserController"] = append(beego.GlobalControllerRouter["cjapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
