package sso

//User 用户信息
type User struct {
	UserName  string `json:"user_name"`
	WxOpID    string `json:"wx_openid"`
	ExtParams string `json:"ext_params"`
	UserID    string `json:"user_id"`
}

//LoginState 用户信息
type LoginState struct {
	UserID    int64  `json:"user_id" m2s:"user_id"`
	UserName  string `json:"user_name" m2s:"user_name"`
	RoleName  string `json:"role_name" m2s:"role_name"`
	SystemID  int    `json:"sys_id" `
	SysIdent  string `json:"ident" `
	RoleID    int    `json:"role_id"`
	Status    int    `json:"status" m2s:"status"`
	IndexURL  string `json:"index_url"`
	ExtParams string `json:"ext_params" m2s:"ext_params"`
}
