package vo

type LoginVo struct {
	Token       string      `json:"token"`
	User        interface{} `json:"user"`
	ExpireAt    int64       `json:"expireAt"`
	Permissions []string    `json:"permissions"`
	Roles       interface{} `json:"roles"`
}
