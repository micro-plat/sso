package sso

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/context"
)

//GetMember 获取登录用户信息
func GetMember(ctx *context.Context) *LoginState {
	v, _ := ctx.Meta.Get("login-state")
	if v == nil {
		return nil
	}
	return v.(*LoginState)
}

//CheckAndSetMember 验证jwt同时保存用户登录信息
//isReallyTimeCheckUser 是否每次api调用都去验证用户信息(状态信息等)
func CheckAndSetMember(ctx *context.Context, isReallyTimeCheckUser ...bool) error {
	if skip, err := ctx.Request.SkipJWTExclude(); err != nil || skip {
		return err
	}

	//验证用户是否登录
	var m LoginState
	if err := ctx.Request.GetJWT(&m); err != nil {
		fmt.Println("获取请求jwt失败：%v", err)
		return context.NewError(context.ERR_FORBIDDEN, err)
	}

	if len(isReallyTimeCheckUser) > 0 && isReallyTimeCheckUser[0] {
		if _, err := GetCurrentUserInfo(m.UserName); err != nil {
			return err
		}
	}

	//保存登录用户信息
	ctx.Meta.Set("login-state", &m)

	//验证用户是否有页面权限
	tags := ctx.GetContainer().GetTags(ctx.Service)
	ctx.Log.Infof("当前接口配置的tags为: %v", tags)
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
		return context.NewErrorf(919, "用户没有相应的按钮级权限")
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
