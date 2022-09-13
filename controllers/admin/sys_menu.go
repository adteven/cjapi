package admin

import (
	"cjapi/common/jwt"
	"cjapi/controllers"
	"cjapi/models"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
)

// SysMenuController operations for SysMenu
type SysMenuController struct {
	controllers.BaseController
}

// URLMapping ...
func (c *SysMenuController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create SysMenu
// @Param	body		body 	models.SysMenu	true		"body for SysMenu content"
// @Success 201 {int} models.SysMenu
// @Failure 403 body is empty
// @router / [post]
func (c *SysMenuController) Post() {
	var v models.SysMenu
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSysMenu(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.Ok(nil)
}

// GetOne ...
// @Title Get One
// @Description get SysMenu by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SysMenu
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SysMenuController) GetOne() {
	id, _ := c.GetInt64(":id", 0)
	v, err := models.GetSysMenuById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get SysMenu
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SysMenu
// @Failure 403
// @router / [get]
func (c *SysMenuController) GetAll() {
	name := c.GetString("blurry")
	menus := models.GetAllMenus(name)
	c.Ok(menus)
}

// @Title 构建菜单
// @Description 菜单构建 获取进入后的菜单树形结构
// @Success 200 {object} controllers.Result
// @router /nav [get]
func (c *SysMenuController) Nav() {
	uid, _ := jwt.GetAdminUserId(c.Ctx.Input)
	logs.Info(uid)
	menus := models.BuildMenus(uid)
	c.Ok(menus)
}

// @Title 获取菜单树
// @Dsecriptioin 获取菜单的属性结构
// @Success 200 {object} controllers.Result
// @router /tree [get]
func (c *SysMenuController) GetTree() {
	menus := models.GetAllMenus("")
	c.Ok(menus)
}

// Put ...
// @Title Put
// @Description update the SysMenu
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SysMenu	true		"body for SysMenu content"
// @Success 200 {object} models.SysMenu
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SysMenuController) Put() {
	id, _ := c.GetInt64(":id", 0)
	v := models.SysMenu{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSysMenuById(&v); err == nil {
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
// @Description delete the SysMenu
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SysMenuController) Delete() {
	id, _ := c.GetInt64(":id", 0)
	if err := models.DeleteSysMenu(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
