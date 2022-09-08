package vo

import "cjapi/models"

type LoginVo struct {
	Token       string                        `json:"token"`
	User        *models.SysUser               `json:"user"`
	ExpireAt    int64                         `json:"expireAt"`
	Permissions []map[interface{}]interface{} `json:"permissions"`
	Roles       []map[interface{}]interface{} `json:"roles"`
}
