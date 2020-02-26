package model

import (
	"encoding/json"
)

//MemberState 用户信息
type MemberState struct {
	Password   string `json:"password,omitempty"`
	UserID     int64  `json:"user_id" m2s:"user_id"`
	UserName   string `json:"user_name" m2s:"user_name"`
	RoleName   string `json:"role_name" m2s:"role_name"`
	SystemID   int    `json:"sys_id" m2s:"sys_id"`
	SysIdent   string `json:"ident" m2s:"ident"`
	RoleID     int    `json:"role_id" m2s:"role_id"`
	Status     int    `json:"status" m2s:"status"`
	IndexURL   string `json:"index_url"`
	ExtParams  string `json:"ext_params" m2s:"ext_params"`
	BelongID   int    `json:"belong_id" m2s:"belong_id"`     //所属编号(加油站就是加油站id)
	BelongType int    `json:"belong_type" m2s:"belong_type"` //所属类型(0:加油站, 1:公司)
	IsOwner    int    `json:"is_owner"`                      //是否管理员
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
