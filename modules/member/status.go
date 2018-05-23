package member

import "github.com/micro-plat/hydra/context"

const maxErrorCnt = 5

//LoginState 用户信息
type LoginState struct {
	UserID   int64  `json:"user_id" m2s:"user_id"`
	UserName string `json:"user_name" m2s:"user_name"`
	SystemID int    `json:"sys_id"`
	RoleID   int    `json:"role_id"`
	Status   int    `json:"status"`
	IndexURL string `json:"index_url"`
}

//Save 保存member信息
func Save(ctx *context.Context, m *LoginState) {
	ctx.Meta.Set("login-state", m)
}

//Get 获取member信息
func Get(ctx *context.Context) *LoginState {
	v, _ := ctx.Meta.Get("login-state")
	return v.(*LoginState)
}
