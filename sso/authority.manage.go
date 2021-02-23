package sso

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"

	//"github.com/micro-plat/sso/sso/errorcode"
	"net/http"
)

const memberCacheKey = "member_cache_key"

//GetMember 获取登录用户信息
func GetMember(ctx hydra.IContext) *LoginState {
	var s LoginState

	state := ctx.Meta().GetValue(memberCacheKey)
	if state != nil {
		return state.(*LoginState)
	}

	if err := ctx.User().Auth().Bind(&s); err != nil {
		return nil
	}
	ctx.Meta().SetValue(memberCacheKey, &s)
	return &s
}

//CheckAndSetMember 验证jwt同时保存用户登录信息
//isReallyTimeCheckUser 是否每次api调用都去验证用户信息(状态信息等)
func CheckAndSetMember(ctx hydra.IContext) (err error) {
	if ctx.User().Auth().Request() == nil {
		return nil
	}
	//验证用户是否登录
	m := GetMember(ctx)
	if m == nil {
		err = errs.NewError(http.StatusUnauthorized, "获取请求jwt失败")
		return
	}

	pageURL, tag, ok := ctx.Request().Path().GetPageAndTag()
	if !ok {
		err = errs.NewErrorf(http.StatusUnauthorized, "请求地址不可用：%s", pageURL)
		return
	}

	hasRight, err := VerifyAuthority(ctx, pageURL, tag)
	if err != nil {
		return
	}
	if !hasRight {
		err = errs.NewErrorf(http.StatusUnauthorized, "请求未授权,page:%s,tag:%s", pageURL, tag)
		return
	}
	return nil
}

//VerifyAuthority VerifyAuthority
func VerifyAuthority(ctx hydra.IContext, pageURL string, funcTag string) (hasRight bool, err error) {
	//验证用户是否登录
	state := GetMember(ctx)
	if state == nil {
		err = errs.NewError(http.StatusForbidden, "获取请求jwt失败")
		return
	}
	reqPath := ctx.Request().Path().GetRequestPath()
	//3.检查是否需要跳过请求
	if ok, _ := authorityMatch.Match(reqPath); ok {
		hasRight = true
		return
	}

	authorityData, err := getRoleTagFromLocal(state.RoleID)
	if err != nil {
		err = errs.NewErrorf(http.StatusUnauthorized, "获取缓存失败：%d,ident:%s,role:%d", state.UserID, state.SysIdent, state.RoleID, err)
		return
	}
	//页面权限
	item, ok := authorityData[pageURL]
	if !ok {
		err = errs.NewErrorf(http.StatusMethodNotAllowed, "用户没有相应的页面权限")
		return
	}

	//没有配置按钮权限
	if len(item.FuncTags) == 0 {
		hasRight = true
		return
	}
	if _, ok := item.FuncTags[funcTag]; !ok {
		err = errs.NewErrorf(http.StatusMethodNotAllowed, "用户没有相应的按钮权限")
	}

	hasRight = true
	return
}
