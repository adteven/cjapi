package admin

import (
	"cjapi/controllers"
	"cjapi/models"
	"cjapi/models/dto"
	"cjapi/models/vo"
	"encoding/json"
	"errors"
	"strings"
)

// TbJishuiAreaController operations for TbJishuiArea
type TbJishuiAreaController struct {
	controllers.BaseController
}

// URLMapping ...
func (c *TbJishuiAreaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetPage", c.GetPage)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TbJishuiArea
// @Param	body		body 	models.TbJishuiArea	true		"body for TbJishuiArea content"
// @Success 201 {int} models.TbJishuiArea
// @Failure 403 body is empty
// @router / [post]
func (c *TbJishuiAreaController) Post() {
	var v dto.JiShuiAreaDto
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	c.Valid(v)
	_, e := models.AddJiShuiArea(&v)
	if e != nil {
		c.Fail(e.Error(), 502)
	}
	c.Ok("操作成功")
}

// @Title 积水点列表
// @Description 积水点分页列表
// @Success 200 {object} app.Response
// @router /page [get]
func (c *TbJishuiAreaController) GetPage() {
	total, list := models.GetAllArea(c.GetParams())
	c.Ok(vo.ResultList{Data: list, Total: total})
}

// GetOne ...
// @Title Get One
// @Description get TbJishuiArea by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TbJishuiArea
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TbJishuiAreaController) GetOne() {
	id, _ := c.GetInt64("id")
	v, err := models.GetJiShuiAreaById(id)
	if err != nil {
		c.Fail(err.Error(), 403)
		return
	}
	c.Ok(v)
}

// GetAll ...
// @Title Get All
// @Description get TbJishuiArea
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TbJishuiArea
// @Failure 403
// @router / [get]
func (c *TbJishuiAreaController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllJiShuiArea(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TbJishuiArea
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TbJishuiArea	true		"body for TbJishuiArea content"
// @Success 200 {object} models.TbJishuiArea
// @Failure 403 :id is not int
// @router / [put]
func (c *TbJishuiAreaController) Put() {
	var v dto.JiShuiAreaDto
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	c.Valid(v)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateJiShuiAreaById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the TbJishuiArea
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TbJishuiAreaController) Delete() {
	id, _ := c.GetInt64(":id")
	if err := models.DeleteJiShuiArea(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
