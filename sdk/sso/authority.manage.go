package sso

import (
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
func CheckAndSetMember(ctx *context.Context) error {
	//有些接口不需要验证
	if skip, err := ctx.Request.SkipJWTExclude(); err != nil || skip {
		return err
	}

	//验证用户是否登录
	var m LoginState
	if err := ctx.Request.GetJWT(&m); err != nil {
		return context.NewError(context.ERR_FORBIDDEN, err)
	}

	//保存登录用户信息
	ctx.Meta.Set("login-state", &m)

	//验证用户是否有页面权限
	tags := ctx.GetContainer().GetTags(ctx.Service)
	ctx.Log.Infof("当前接口配置的tags为: %v", tags)
	if tags == nil || len(tags) == 0 {
		return nil
	}

	menu, err := getUserMenuFromLocal(int(m.UserID))
	if err != nil {
		return err
	}

	for _, tag := range tags {
		if strings.Trim(tag, " ") == "" {
			continue
		}
		if flag := verifyAuthority(menu, tag); !flag {
			return context.NewError(919, "用户没有相关页面权限")
		}
	}
	return nil
}

//verifyAuthority 验证用户页面权限
func verifyAuthority(menu []Menu, tag string) bool {
	for _, val := range menu {
		for _, first := range val.Children {
			for _, second := range first.Children {
				if strings.EqualFold(second.Path, tag) {
					return true
				}
			}
		}
	}
	return false
}
