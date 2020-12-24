package sso

import (
	"net/http"
	"strings"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

//GetMember 获取登录用户信息
func GetMember(ctx hydra.IContext) *LoginState {
	var s LoginState
	if err := ctx.User().Auth().Bind(&s); err != nil {
		return nil
	}
	return &s
}

//CheckAndSetMember 验证jwt同时保存用户登录信息
//isReallyTimeCheckUser 是否每次api调用都去验证用户信息(状态信息等)
func CheckAndSetMember(ctx hydra.IContext, isReallyTimeCheckUser ...bool) error {
	if ctx.User().Auth().Request() == nil {
		return nil
	}

	//验证用户是否登录
	m := GetMember(ctx)
	if m == nil {
		return errs.NewError(http.StatusForbidden, "获取请求jwt失败")
	}

	if len(isReallyTimeCheckUser) > 0 && isReallyTimeCheckUser[0] {
		if _, err := GetCurrentUserInfo(m.UserName); err != nil {
			return err
		}
	}

	//验证用户是否有页面权限
	router, err := ctx.Request().Path().GetRouter()
	tags := router.Pages
	ctx.Log().Infof("当前接口配置的tags为: %v", tags)
	if tags == nil || len(tags) == 0 {
		return nil
	}

	userHasTags, err := getUserTagFromLocal(int(m.UserID))
	if err != nil {
		return err
	}
	configTag := strings.TrimSpace(tags[0])
	if configTag == "" || configTag == "*" {
		return nil
	}
	if flag := verifyAuthority(userHasTags, configTag); !flag {
		return errs.NewErrorf(919, "用户没有相应的按钮级权限")
	}
	return nil
}

//verifyAuthority 验证用户是否包含此tag
func verifyAuthority(userHasTags []Menu, tag string) bool {
	for _, temp := range userHasTags {
		if strings.EqualFold(strings.TrimSpace(temp.Path), tag) {
			return true
		}
	}
	return false
}
