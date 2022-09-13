package models

import (
	"cjapi/models/dto"
	"cjapi/runtime"
	"github.com/beego/beego/v2/client/orm"
)

type SysRolesMenus struct {
	MenuId *SysMenu `orm:"column(sys_menu_id);rel(fk)"`
	RoleId *SysRole `orm:"column(sys_role_id);rel(fk)"`
	Id     int64    `orm:"column(id);auto"`
}

func (t *SysRolesMenus) TableName() string {
	return "sys_roles_menus"
}

func init() {
	orm.RegisterModel(new(SysRolesMenus))
}

func BatchRoleMenuAdd(menu dto.RoleMenu) {
	o := orm.NewOrm()
	o.Raw("delete from sys_roles_menus WHERE role_id = ?", menu.Id).Exec()

	var roleMenus []SysRolesMenus
	var roles = GetOneRole(menu.Id)
	cb := runtime.Runtime.GetCasbinKey("default")
	cb.RemoveFilteredPolicy(0, roles.Permission)
	for _, val := range menu.Menus {

		//var menus = SysMenu{Id: val.Id}
		//var roles = SysRole{Id: menu.Id}
		var menus, _ = GetSysMenuById(val.Id)

		roleMenus = append(roleMenus, SysRolesMenus{MenuId: menus, RoleId: roles})

		cb.AddNamedPolicy("p", roles.Permission, menus.Router, menus.RouterMethod)
	}

	cb.SavePolicy()

	o.InsertMulti(100, roleMenus)
}
