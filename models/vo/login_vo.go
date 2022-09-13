package vo

import "cjapi/models"

type LoginVo struct {
	Token       string            `json:"token"`
	User        *models.SysUser   `json:"user"`
	ExpireAt    int64             `json:"expireAt"`
	Permissions []string          `json:"permissions"`
	Roles       []*models.SysRole `json:"roles"`
}
