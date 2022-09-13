package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type SysUsersRoles struct {
	Id     int64    `orm:"column(id);auto"`
	UserId *SysUser `orm:"column(sys_user_id);rel(fk)" description:"用户ID"`
	RoleId *SysRole `orm:"column(sys_role_id);rel(fk)" description:"角色ID"`
}

func (t *SysUsersRoles) TableName() string {
	return "sys_users_roles"
}

func init() {
	orm.RegisterModel(new(SysUsersRoles))
}

// AddSysUsersRoles insert a new SysUsersRoles into database and returns
// last inserted Id on success.
func AddSysUsersRoles(m *SysUsersRoles) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
