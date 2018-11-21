package subsys

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/sso/flowserver/modules/member"
	"github.com/micro-plat/sso/flowserver/modules/operate"
	"github.com/micro-plat/sso/flowserver/modules/system"
	"github.com/micro-plat/sso/flowserver/modules/util"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c      component.IContainer
	m      member.IMember
	code   member.ICodeMember
	wxcode member.IWxcode
	sys    system.ISystem
	op     operate.IOperate
	member member.IDBMember
	cache  member.ICacheMember
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c:      container,
		m:      member.NewMember(container),
		code:   member.NewCodeMember(container),
		wxcode: member.NewWxcode(container),
		sys:    system.NewSystem(container),
		op:     operate.NewOperate(container),
		member: member.NewDBMember(container),
		cache:  member.NewCacheMember(container),
	}
}

//Handle 子系统远程登录
func (u *LoginHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统用户远程登录---------")

	ctx.Log.Info("1.检查参数")
	if err := ctx.Request.Check("username", "password", "ident", "timestamp", "sign"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	//获取secret
	secret, err := u.getSecret(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	//校验签名
	d := map[string]interface{}{}
	d["username"] = ctx.Request.GetString("username")
	d["password"] = ctx.Request.GetString("password")
	d["ident"] = ctx.Request.GetString("ident")
	d["timestamp"] = ctx.Request.GetString("timestamp")
	ctx.Log.Info("请求原数据", d)
	if ok := util.VerifySign(ctx, d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}

	ctx.Log.Info("2.执行操作")
	//处理用户登录
	member, err := u.m.Login(ctx.Request.GetString("username"),
		md5.Encrypt(ctx.Request.GetString("password")),
		ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	//保存用户信息
	_, err = u.code.Save(member)
	if err != nil {
		return err
	}

	//记录登录行为
	if err := u.op.LoginOperate(member); err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return member

}

func (u *LoginHandler) getSecret(ident string) (string, error) {
	if ident == "" {
		return "", context.NewError(context.ERR_NOT_ACCEPTABLE, "ident not exists")
	}
	data, err := u.sys.Get(ident)
	if err != nil {
		return "", err
	}
	return data.GetString("secret"), nil
}
