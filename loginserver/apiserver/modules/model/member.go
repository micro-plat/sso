package model

import (
	"encoding/json"

	"github.com/micro-plat/hydra/components"
)

//MemberState 用户信息
type MemberState struct {
	Password       string `json:"-"`
	UserID         int64  `json:"user_id" m2s:"user_id"`
	FullName       string `json:"full_name" m2s:"full_name"`
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
	LastLoginTime  string `json:"last_login_time"`
	LoginTimeout   int    `json:"login_timeout" m2s:"login_timeout"`
	ExtParams      string `json:"ext_params" m2s:"ext_params"`
}

//ReflushCode 刷新登录code
func (m *MemberState) ReflushCode() string {
	m.Code = components.Def.UUID().ToString()[0:6]
	return m.Code
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
