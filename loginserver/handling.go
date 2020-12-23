package main

//bind 检查应用程序配置文件，并根据配置初始化服务
func handling() {
	//每个请求执行前执行
	// app.Handling(func(ctx context.IContext) (rt interface{}) {

	// 	//跳过jwt排除的请求
	// 	if skip, err := ctx.Request.SkipJWTExclude(); err != nil || skip {
	// 		return err
	// 	}

	// 	//保存登录信息
	// 	var m model.LoginState
	// 	if err := ctx.Request.GetJWT(&m); err != nil {
	// 		return context.NewError(context.ERR_FORBIDDEN, err)
	// 	}

	// 	if err := member.Save(ctx, &m); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })
}
