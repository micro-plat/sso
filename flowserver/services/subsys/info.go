package subsys

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/flowserver/modules/system"
	"github.com/micro-plat/sso/flowserver/modules/util"
)

//InfoHandler 菜单查询对象
type InfoHandler struct {
	c   component.IContainer
	sys system.ISystem
}

//NewInfoHandler 创建菜单查询对象
func NewInfoHandler(container component.IContainer) (u *InfoHandler) {
	return &InfoHandler{
		c:   container,
		sys: system.NewSystem(container),
	}
}

//Handle 查询指定用户在指定系统的菜单列表
func (u *InfoHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统调用，获取系统信息------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("ident", "timestamp", "sign"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}
	//获取secret
	secret, err := u.getSecret(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	//校验签名
	d := map[string]interface{}{}
	d["ident"] = ctx.Request.GetString("ident")
	d["timestamp"] = ctx.Request.GetString("timestamp")
	ctx.Log.Info("请求请求系统信息数据：", d)
	if ok := util.VerifySign(ctx, d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}
	ctx.Log.Info("2. 执行操作")
	data, err := u.sys.Get(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3. 返回数据")
	return data
}

func (u *InfoHandler) getSecret(ident string) (string, error) {
	if ident == "" {
		return "", context.NewError(context.ERR_NOT_ACCEPTABLE, "ident not exists")
	}
	data, err := u.sys.Get(ident)
	if err != nil {
		return "", err
	}

	return data.GetString("secret"), nil
}
