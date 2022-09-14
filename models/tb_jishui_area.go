package models

import (
	"cjapi/models/dto"
	"cjapi/models/vo"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/jinzhu/copier"
	"reflect"
	"strconv"
	"strings"
)

type JiShuiArea struct {
	Area           string `orm:"column(area);size(128)" description:"属于哪个区"`
	Id             int64  `orm:"column(id);auto" description:"id"`
	JsNumber       string `orm:"column(js_number);size(100)" description:"水位"`
	Lat            string `orm:"column(lat);size(128)" description:"维度"`
	Lng            string `orm:"column(lng);size(128)" description:"经度"`
	Name           string `orm:"column(name);size(255)" description:"产品名称"`
	Road           string `orm:"column(road);size(255)" description:"道路名称"`
	ProjectName    string `orm:"column(project_name);size(100)" description:"工程名称"`
	ProjectContent string `orm:"column(project_content);size(512)" description:"工程内容"`
	ProjectUnit    string `orm:"column(project_unit);size(100)" description:"责任单位"`
	ProjectPeriod  string `orm:"column(project_period);size(100)" description:"施工年限"`
	Pictures       string `orm:"column(pictures);size(512)" description:"积水点图片"`
	Process        string `orm:"column(process);size(255)" description:"处理进度"`
	BaseModel
}

func (t *JiShuiArea) TableName() string {
	return "tb_jishui_area"
}

func init() {
	orm.RegisterModel(new(JiShuiArea))
}

func GetAllArea(base dto.BasePage, query ...interface{}) (int, []vo.JiShuiAreaVo) {
	var (
		tableName = "tb_jishui_area"
		data      []JiShuiArea
		condition = ""
	)
	if base.Blurry != "" {
		condition = " and name= '" + base.Blurry + "'"
	}
	if len(query) > 0 {
		groupId := query[0].(int64)
		if groupId > 0 {
			condition += " and group_id=" + strconv.FormatInt(groupId, 10)
		}
	}
	total, _, rs := GetPagesInfo(tableName, base.Page, base.Size, condition)
	rs.QueryRows(&data)
	var voData []vo.JiShuiAreaVo
	for _, q := range data {
		var vo vo.JiShuiAreaVo
		copier.Copy(&vo, &q)
		vo.Pictures = strings.Split(q.Pictures, ",")
		voData = append(voData, vo)
	}
	return total, voData
}

// AddTbJiShuiArea insert a new TbJiShuiArea into database and returns
// last inserted Id on success.
func AddJiShuiArea(dto *dto.JiShuiAreaDto) (id int64, err error) {
	var m JiShuiArea
	copier.Copy(&m, dto)
	m.Pictures = strings.Join(dto.Pictures, ",")
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// GetTbJiShuiAreaById retrieves TbJiShuiArea by Id. Returns error if
// Id doesn't exist
func GetJiShuiAreaById(id int64) (v *JiShuiArea, err error) {
	o := orm.NewOrm()
	v = &JiShuiArea{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTbJiShuiArea retrieves all TbJiShuiArea matches certain condition. Returns empty list if
// no records exist
func GetAllJiShuiArea(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(JiShuiArea))
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

	var l []JiShuiArea
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

// UpdateTbJiShuiArea updates TbJiShuiArea by Id and returns error if
// the record to be updated doesn't exist
func UpdateJiShuiAreaById(dto *dto.JiShuiAreaDto) (err error) {
	o := orm.NewOrm()
	v := JiShuiArea{Id: dto.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		var m JiShuiArea
		copier.Copy(&m, dto)
		m.Pictures = strings.Join(dto.Pictures, ",")
		if num, err = o.Update(&m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTbJiShuiArea deletes TbJiShuiArea by Id and returns error if
// the record to be deleted doesn't exist
func DeleteJiShuiArea(id int64) (err error) {
	o := orm.NewOrm()
	v := JiShuiArea{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&JiShuiArea{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
