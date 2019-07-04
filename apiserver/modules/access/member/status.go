package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/model"
)

const maxErrorCnt = 5

<<<<<<< HEAD:apiserver/modules/access/member/status.go
/*
=======
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

>>>>>>> 750f5c63baeb3b4a71bc53caecd154a8e0ed6969:flowserver/modules/member/status.go
//LoginState 用户登录状态
type LoginState model.MemberState

//MarshalJSON 修改marshal行为，去掉敏感字段
func (m LoginState) MarshalJSON() ([]byte, error) {
	type mem model.MemberState
	current := mem(m)
	current.Password = ""
	return json.Marshal((*mem)(&current))
}*/

/*
//ReflushCode 刷新登录code
func (m *MemberState) ReflushCode() string {
	m.Code = utility.GetGUID()[0:6]
	return m.Code
}
*/

//Save 保存member信息
func Save(ctx *context.Context, m *model.LoginState) error {
	//不允许同一个账户多处登录
	//container := ctx.GetContainer()
	//v, ok := container.Get("login-code").(ICacheMember)
	//if !ok {
	//	v = NewCacheMember(container)
	//	container.Set("login-code", v)
	//}
	//ms, err := v.Query(m.UserName, m.SystemID)
	//if err != nil {
	//	return context.NewError(403, fmt.Sprintf("登录信息已过期，请重新登录%v", err))
	//}
	//if ms.Code != m.Code {
	//	return context.NewError(403, "用户登录code已过期，请重新登录")
	//}
	////检查用户是否已锁定
	//if ms.Status == UserLock {
	//	return context.NewError(context.ERR_LOCKED, "用户被锁定暂时无法登录")
	//}
	////检查用户是否已禁用
	//if ms.Status == UserDisable {
	//	return context.NewError(context.ERR_FORBIDDEN, "用户被禁用请联系管理员")
	//}
	ctx.Meta.Set("login-state", m)
	return nil
}

//Get 获取member信息
func Get(ctx *context.Context) *model.LoginState {
	v, _ := ctx.Meta.Get("login-state")
	if v == nil {
		return nil
	}
	return v.(*model.LoginState)
}

func Query(ctx *context.Context, container component.IContainer) *model.LoginState {
	m := &model.LoginState{}
	err := ctx.Request.GetJWT(m)
	if m.UserName != "" && err == nil {
		return m
	}
	if err := ctx.Request.Check("code"); err != nil {
		return nil
	}
	codeMemberLib := NewCodeMember(container)
	m, err = codeMemberLib.Query(ctx.Request.GetString("code"))
	if err != nil {
		return nil
	}
	return m
}
