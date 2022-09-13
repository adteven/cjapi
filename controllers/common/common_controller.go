package common

import (
	"cjapi/controllers"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type CommonController struct {
	controllers.BaseController
}

func (c *CommonController) URLMapping() {
	c.Mapping("Upload", c.Upload)
}

// @Title 上传图像
// @Description 上传图像
// @Success 200 {object} controllers.Result
// @router /upload [post]
func (c *CommonController) Upload() {
	logs.Info("======file start======")
	f, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	defer f.Close()
	var path = "static/upload/" + h.Filename
	e := c.SaveToFile("file", path) // 保存位置在 static/upload, 没有文件夹要先创建
	logs.Error(e)
	if e != nil {
		c.Fail(e.Error(), 5009)
	}
	apiUrl, _ := beego.AppConfig.String("api_url")
	imgUrl := apiUrl + "/" + path

	c.Ok(imgUrl)
}
