package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"reflect"
	"strings"
)

type SysRole struct {
	Id         int64      `json:"id" orm:"column(id);auto" description:"ID"`
	Name       string     `json:"name" orm:"column(name);size(255)" description:"名称"`
	Remark     string     `json:"remark" orm:"column(remark);size(255);null" description:"备注"`
	DataScope  string     `json:"dataScope" orm:"column(data_scope);size(255);null" description:"数据权限"`
	Level      int        `json:"level" orm:"column(level);null" description:"角色级别"`
	Permission string     `json:"permission" orm:"column(permission);size(255);null" description:"功能权限"`
	Users      []*SysUser `orm:"reverse(many)"`
	Menus      []*SysMenu `json:"menus"  orm:"rel(m2m);rel_through(cjapi/models.SysRolesMenus)"`
	Depts      []*SysDept `json:"depts" orm:"rel(m2m);rel_through(cjapi/models.SysRolesDepts)"`
	BaseModel
}

func (t *SysRole) TableName() string {
	return "sys_role"
}

func init() {
	orm.RegisterModel(new(SysRole))
}

// AddSysRole insert a new SysRole into database and returns
// last inserted Id on success.
func AddSysRole(m *SysRole) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetOneRole(id int64) *SysRole {
	o := orm.NewOrm()
	role := SysRole{Id: id}
	o.Read(&role)
	_, _ = o.LoadRelated(&role, "Menus")
	return &role
}

// GetSysRoleById retrieves SysRole by Id. Returns error if
// Id doesn't exist
func GetSysRoleById(id int64) (v *SysRole, err error) {
	o := orm.NewOrm()
	v = &SysRole{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysRole retrieves all SysRole matches certain condition. Returns empty list if
// no records exist
func GetAllSysRole(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysRole))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SysRole
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSysRole updates SysRole by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysRoleById(m *SysRole) (err error) {
	o := orm.NewOrm()
	v := SysRole{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysRole deletes SysRole by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysRole(id int64) (err error) {
	o := orm.NewOrm()
	v := SysRole{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysRole{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
