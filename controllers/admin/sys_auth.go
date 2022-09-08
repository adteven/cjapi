package admin

import (
	"cjapi/common/jwt"
	"cjapi/common/utils"
	"cjapi/controllers"
	"cjapi/models"
	"cjapi/models/dto"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

type AuthControler struct {
	controllers.BaseController
}

func (c *AuthControler) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Logout", c.Logout)
}

// @Title 登录
// @Description 登录
// @Success 200 {object} controllers.Result
// @router /login [post]
func (c *AuthControler) Login() {
	var authUser *dto.AuthUser

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &authUser)

	logs.Info(authUser)

	if err == nil {
		currentUser, e := models.GetUserByUsername(authUser.Username)

		if e != nil {
			c.Fail("用户不存在", 5002)
		}
		logs.Info("=======currentUser======")
		logs.Info(currentUser)
		if !utils.ComparePwd(currentUser.Password, []byte(authUser.Password)) {
			c.Fail("密码错误", 5003)
		} else {
			loginVO, _ := jwt.GenerateToken(currentUser, time.Hour*24*100)
			c.Ok(loginVO)
		}
	} else {
		c.Fail(err.Error(), 5004)
	}
}

// @Title 退出登录
// @Description 退出登录
// @Success 200 {object} controllers.Result
// @router /logout [delete]
func (c *AuthControler) Logout() {
	err := jwt.RemoveUser(c.Ctx.Input)
	if err != nil {
		c.Fail("退出失败", 5005)
	} else {
		c.Ok("退出成功")
	}
}
