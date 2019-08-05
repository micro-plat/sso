package model

import "encoding/json"

//MemberState 用户信息
type MemberState struct {
	Password       string `json:"password,omitempty"`
	UserID         int64  `json:"user_id" m2s:"user_id"`
	UserName       string `json:"user_name" m2s:"user_name"`
	RoleName       string `json:"role_name" m2s:"role_name"`
	SystemID       int    `json:"sys_id" `
	SysIdent       string `json:"ident" `
	RoleID         int    `json:"role_id"`
	Status         int    `json:"status" m2s:"status"`
	IndexURL       string `json:"index_url"`
	LoginURL       string `json:"login_url"`
	Code           string `json:"code"`
	ProfilePercent int    `json:"profile_percent"`
	LoginTimeout   int    `json:"login_timeout" m2s:"login_timeout"`
	ExtParams      string `json:"ext_params" m2s:"ext_params"`
}

//LoginState 用户登录状态
type LoginState MemberState

//MarshalJSON 修改marshal行为，去掉敏感字段
func (m LoginState) MarshalJSON() ([]byte, error) {
	type mem MemberState
	current := mem(m)
	current.Password = ""
	return json.Marshal((*mem)(&current))
}
