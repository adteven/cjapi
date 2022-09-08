package controllers

import (
	"cjapi/models/dto"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) GetParams() dto.BasePage {
	var (
		page   int
		size   int
		blurry string
	)
	c.Ctx.Input.Bind(&page, "page")
	c.Ctx.Input.Bind(&size, "size")
	c.Ctx.Input.Bind(&blurry, "blurry")

	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}
	return dto.BasePage{Page: page, Size: size, Blurry: blurry}
}

type any = interface{}

type Result struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"`
}

func (c *BaseController) Valid(data any) {
	valid := validation.Validation{}
	b, _ := valid.Valid(data)
	if !b {
		for _, err := range valid.Errors {
			c.Fail(err.Message, 5001)
		}
	}
}

func (c *BaseController) Ok(data any) {
	c.Data["json"] = SuccessData(data)
	c.ServeJSON()
}

func (c *BaseController) Fail(msg string, status int) {
	c.Data["json"] = ErrMsg(msg, status)
	c.ServeJSON()
}

func ErrMsg(msg string, status ...int) Result {
	var r Result
	if len(status) > 0 {
		r.Status = status[0]
	} else {
		r.Status = 500000
	}
	r.Msg = msg
	r.Data = nil

	return r
}

func ErrData(msg error, status ...int) Result {
	var r Result
	if len(status) > 0 {
		r.Status = status[0]
	} else {
		r.Status = 500000
	}
	r.Msg = msg.Error()
	r.Data = nil

	return r
}

func SuccessData(data any) Result {
	var r Result

	r.Status = 200
	r.Msg = "ok"
	r.Data = data

	return r
}
