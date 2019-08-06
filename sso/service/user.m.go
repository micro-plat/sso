package service

import "encoding/json"

//User 用户信息
type User struct {
	UserName  string `json:"user_name"`
	WxOpID    string `json:"wx_openid"`
	ExtParams string `json:"ext_params"`
	UserID    string `json:"user_id"`
}

//MemberState 用户信息
type MemberState struct {
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

//LoginState 用户登录状态
type LoginState MemberState

//MarshalJSON 修改marshal行为，去掉敏感字段
func (m LoginState) MarshalJSON() ([]byte, error) {
	type mem MemberState
	current := mem(m)
	return json.Marshal((*mem)(&current))
}
