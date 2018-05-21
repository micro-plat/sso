package member

import "github.com/micro-plat/hydra/context"

//LoginState 用户信息
type LoginState struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	SystemID int    `json:"sys_id"`
	RoleID   int    `json:"role_id"`
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
