package sso

//User 用户信息
type User struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	FullNme   string `json:"full_name"`
	WxOpID    string `json:"wx_openid"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	ExtParams string `json:"ext_params"`
	Status    string `json:"status"`
}

//LoginState 用户信息
type LoginState struct {
	UserID        int64  `json:"user_id" m2s:"user_id"`
	UserName      string `json:"user_name" m2s:"user_name"`
	FullName      string `json:"full_name" m2s:"full_name"`
	RoleName      string `json:"role_name" m2s:"role_name"`
	SystemID      int    `json:"sys_id" `
	SysIdent      string `json:"ident" `
	RoleID        int    `json:"role_id"`
	Status        int    `json:"status" m2s:"status"`
	IndexURL      string `json:"index_url"`
	ExtParams     string `json:"ext_params" m2s:"ext_params"`
	Source        string `json:"source" m2s:"source"`                   //来源
	SourceID      string `json:"source_id" m2s:"source_id"`             //来源id
	LastLoginTime string `json:"last_login_time" m2s:"last_login_time"` //上次登录时间
}
