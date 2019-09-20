package main

import (
	"github.com/micro-plat/hydra/context"
	ssosdk "github.com/micro-plat/sso/sdk/sso"
)

//bind 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) handling() {
	//每个请求执行前执行
	r.Handling(func(ctx *context.Context) (rt interface{}) {
		//验证jwt并缓存登录用户信息
		if err := ssosdk.CheckAndSetMember(ctx); err != nil {
			return err
		}
		return nil
	})
}
