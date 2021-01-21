package model

import (
	"encoding/json"

	"github.com/micro-plat/hydra/components"
)

//MemberState 用户信息
type MemberState struct {
	Password       string `json:"password,omitempty"`
	UserID         int64  `json:"user_id" m2s:"user_id"`
	UserName       string `json:"user_name" m2s:"user_name"`
	FullName       string `json:"full_name" m2s:"full_name"`
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
	Source         string `json:"source" m2s:"source"`                   //来源
	SourceID       string `json:"source_id" m2s:"source_id"`             //来源id
	LastLoginTime  string `json:"last_login_time" m2s:"last_login_time"` //上次登录时间
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

//ReflushCode 刷新登录code
func (m *MemberState) ReflushCode() string {
	m.Code = components.Def.UUID().ToString()[0:6]
	return m.Code
}

//LoginReq　登录用到的数据
type LoginReq struct {
	UserName  string //用户名
	Password  string //密码
	Ident     string //系统标识
	ValidCode string //验证码
}
