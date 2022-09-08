package controllers

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Fail("资源不存在", 404)
}
func (c *ErrorController) Error500() {
	c.Fail("服务器内部错误", 500)
}
