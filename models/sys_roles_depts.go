package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type SysRolesDepts struct {
	RoleId *SysRole `orm:"column(role_id);rel(fk)"`
	DeptId *SysDept `orm:"column(dept_id);rel(fk)"`
	Id     int64    `orm:"column(id);auto"`
}

func (t *SysRolesDepts) TableName() string {
	return "sys_roles_depts"
}

func init() {
	orm.RegisterModel(new(SysRolesDepts))
}
