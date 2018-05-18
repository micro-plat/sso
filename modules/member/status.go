package member

import "github.com/micro-plat/hydra/context"

//Member 用户信息
type Member struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	SystemID int    `json:"sys_id"`
	RoleID   int    `json:"role_id"`
}

//Save 保存member信息
func Save(ctx *context.Context, m *Member) {
	ctx.Meta.Set("login-state", m)
}

//Get 获取member信息
func Get(ctx *context.Context) *Member {
	v, _ := ctx.Meta.Get("login-state")
	return v.(*Member)
}
