package sso

import (
	"github.com/micro-plat/hydra/context"
)

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

//SaveSSOClient  保存sso client
func saveSSOClient(m *Client) {
	ssoClient = m
}

//GetSSOClient  获取sso client
func getSSOClient() *Client {
	return ssoClient
}

//CheckAndSetMember 验证jwt同时保存用户登录信息
func CheckAndSetMember(ctx *context.Context) error {
	if skip, err := ctx.Request.SkipJWTExclude(); err != nil || skip {
		return err
	}

	var m LoginState
	if err := ctx.Request.GetJWT(&m); err != nil {
		return context.NewError(context.ERR_FORBIDDEN, err)
	}
	ctx.Meta.Set("login-state", &m)

	return nil
}

//GetMember 获取member信息
func GetMember(ctx *context.Context) *LoginState {
	v, _ := ctx.Meta.Get("login-state")
	if v == nil {
		return nil
	}
	return v.(*LoginState)
}
