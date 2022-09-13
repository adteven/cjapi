package models

import (
	"cjapi/common/utils"
	"cjapi/models/vo/menu"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"reflect"
	"strings"
)

type SysMenu struct {
	Id            int64     `json:"id" orm:"column(id);auto" description:"ID"`
	IFrame        int8      `json:"iframe" orm:"column(i_frame);null" description:"是否外链"`
	Name          string    `json:"name" orm:"column(name);size(255);null" valid:"Required;" description:"菜单名称"`
	Component     string    `json:"component" orm:"column(component);size(255);null" description:"组件"`
	Pid           int64     `json:"pid" orm:"column(pid)" description:"上级菜单ID"`
	Sort          int       `json:"sort" orm:"column(sort)" description:"排序"`
	Icon          string    `json:"icon" orm:"column(icon);size(255);null" description:"图标"`
	Path          string    `json:"path" orm:"column(path);size(255);null" description:"链接地址"`
	Cache         int8      `json:"cache" orm:"column(cache);null" description:"缓存"`
	Hidden        int8      `json:"hidden" orm:"column(hidden);null" description:"是否隐藏"`
	ComponentName string    `json:"componentName" orm:"column(component_name);size(20);null" description:"组件名称"`
	Permission    string    `json:"permission" orm:"column(permission);size(255);null" description:"权限"`
	Type          int       `json:"type" orm:"column(type);null" description:"类型"`
	Router        string    `json:"router" orm:"column(router);size(255);null" description:"操作的路由"`
	RouterMethod  string    `json:"routerMethod" orm:"column(router_method);size(255);null" description:"路由动作"`
	Children      []SysMenu `json:"children" orm:"-"`
	Label         string    `json:"label" orm:"-" `
	BaseModel
}

func (t *SysMenu) TableName() string {
	return "sys_menu"
}

func init() {
	orm.RegisterModel(new(SysMenu))
}

// AddSysMenu insert a new SysMenu into database and returns
// last inserted Id on success.
func AddSysMenu(m *SysMenu) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysMenuById retrieves SysMenu by Id. Returns error if
// Id doesn't exist
func GetSysMenuById(id int64) (v *SysMenu, err error) {
	o := orm.NewOrm()
	v = &SysMenu{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllMenus(name string) []SysMenu {
	var menus []SysMenu
	o := orm.NewOrm()
	qs := o.QueryTable("sys_menu").Filter("is_del", 0).OrderBy("sort")
	if name != "" {
		qs = qs.Filter("name", name)
	}

	qs.All(&menus)
	return RecursionMenuList(menus, 0)
}

//递归函数
func RecursionMenuList(data []SysMenu, pid int64) []SysMenu {
	var listTree = make([]SysMenu, 0)
	for _, value := range data {
		value.Label = value.Name
		if value.Pid == pid {
			value.Children = RecursionMenuList(data, value.Id)
			listTree = append(listTree, value)
		}
	}
	return listTree
}

// GetAllSysMenu retrieves all SysMenu matches certain condition. Returns empty list if
// no records exist
func GetAllSysMenu(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysMenu))
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

	var l []SysMenu
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

// UpdateSysMenu updates SysMenu by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysMenuById(m *SysMenu) (err error) {
	o := orm.NewOrm()
	v := SysMenu{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DelByMenuIds(ids []int64) (err error) {
	str := utils.ReturnQ(len(ids))
	logs.Info(str)
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE sys_menu SET is_del = ? WHERE id in("+str+")", 1, ids).Exec()
	return
}

// DeleteSysMenu deletes SysMenu by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysMenu(id int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE sys_menu SET is_del = ? WHERE id = ", 1, id).Exec()
	return err
}

//获取权限string
func FindByRouterAndMethod(url string, method string) (permission string) {
	o := orm.NewOrm()
	var menu SysMenu
	err := o.QueryTable(new(SysMenu)).Filter("router", url).Filter("router_method", method).One(&menu)
	if err != nil {
		return ""
	}
	return menu.Permission
}

/**
内部使用根据url和method获取菜单
*/
func FindMenuByRouterAndMethod(url string, method string) SysMenu {
	o := orm.NewOrm()
	var menu SysMenu
	o.QueryTable(new(SysMenu)).Filter("router", url).Filter("router_method", method).One(&menu)
	return menu
}

// 包装menus列表
func BuildMenus(uid int64) []menu.MenuVo {
	o := orm.NewOrm()
	var lists orm.ParamsList
	_, err := o.Raw("SELECT r.* FROM sys_role r, sys_users_roles u "+
		"WHERE r.id = u.role_id AND u.user_id = ?", uid).ValuesFlat(&lists, "id")
	if err != nil {
		logs.Error(err)
	}
	idsStr := utils.Convert(lists)
	logs.Info(idsStr)
	var menus []SysMenu
	_, e := o.Raw("select m.* from sys_menu m LEFT OUTER JOIN sys_roles_menus t on m.id= t.menu_id "+
		"LEFT OUTER JOIN sys_role r on r.id = t.role_id where m.is_del=0  and m.type!=2 and r.id in (?) "+
		"order by m.sort asc",
		idsStr).QueryRows(&menus)

	if e != nil {
		logs.Error(e)
	}

	return buildMenus(buildTree(menus))

}

func buildTree(menus []SysMenu) []SysMenu {
	var trees []SysMenu
	for _, menu := range menus {
		if menu.Pid == 0 {
			trees = append(trees, menu)
		}
	}

	for k, tree := range trees {
		var child []SysMenu
		for _, it := range menus {
			if it.Pid == tree.Id {
				child = append(child, it)
			}
		}
		trees[k].Children = child
	}

	return trees

}

func buildMenus(menus []SysMenu) []menu.MenuVo {
	var list []menu.MenuVo
	for _, menuO := range menus {
		menuList := menuO.Children
		var menuVo = new(menu.MenuVo)

		if menuO.ComponentName != "" {
			menuVo.Name = menuO.ComponentName
		} else {
			menuVo.Name = menuO.Name
		}
		if menuO.Pid == 0 {
			menuVo.Path = "/" + menuO.Path
		} else {
			menuVo.Path = menuO.Path
		}

		if menuO.Hidden == 1 {
			menuVo.Hidden = true
		} else {
			menuVo.Hidden = false
		}

		//判断不是外链
		if menuO.IFrame == 0 {
			if menuO.Pid == 0 {
				if menuO.Component == "" {
					menuVo.Component = "Layout"
				} else {
					menuVo.Component = menuO.Component
				}
			} else if menuO.Component != "" {
				menuVo.Component = menuO.Component
			}
		}

		menuVo.Meta = menu.MenuMetaVo{Title: menuO.Name, Icon: menuO.Icon, NoCache: !utils.IntToBool(menuO.Cache)}

		if len(menuList) > 0 {
			menuVo.AlwaysShow = true
			menuVo.Redirect = "noredirect"
			menuVo.Children = buildMenus(menuList)
		} else if menuO.Pid == 0 {
			var menuVo1 = new(menu.MenuVo)
			menuVo1.Meta = menuVo.Meta
			if menuO.IFrame == 0 {
				menuVo1.Path = "index"
				menuVo1.Name = menuVo.Name
				menuVo1.Component = menuVo.Component
			} else {
				menuVo1.Path = menuO.Path
			}
			menuVo.Name = ""
			menuVo.Meta = menu.MenuMetaVo{}
			menuVo.Component = "Layout"
			var list1 []menu.MenuVo
			list1 = append(list1, *menuVo1)
			menuVo.Children = list1
		}

		list = append(list, *menuVo)

	}

	return list
}
