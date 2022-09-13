package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type SysDept struct {
	Id         int       `orm:"column(id);auto" description:"ID"`
	Name       string    `orm:"column(name);size(255)" description:"名称"`
	Pid        int64     `orm:"column(pid)" description:"上级部门"`
	Enabled    int8      `orm:"column(enabled)" description:"状态"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" description:"创建日期"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null"`
	IsDel      int8      `orm:"column(is_del);null"`
}

func (t *SysDept) TableName() string {
	return "sys_dept"
}

func init() {
	orm.RegisterModel(new(SysDept))
}

// AddSysDept insert a new SysDept into database and returns
// last inserted Id on success.
func AddSysDept(m *SysDept) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysDeptById retrieves SysDept by Id. Returns error if
// Id doesn't exist
func GetSysDeptById(id int) (v *SysDept, err error) {
	o := orm.NewOrm()
	v = &SysDept{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysDept retrieves all SysDept matches certain condition. Returns empty list if
// no records exist
func GetAllSysDept(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysDept))
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

	var l []SysDept
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

// UpdateSysDept updates SysDept by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysDeptById(m *SysDept) (err error) {
	o := orm.NewOrm()
	v := SysDept{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysDept deletes SysDept by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysDept(id int) (err error) {
	o := orm.NewOrm()
	v := SysDept{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysDept{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
