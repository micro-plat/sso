package user

import (
	"fmt"
	"net/url"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/utility"
	"github.com/micro-plat/sso/modules/app"
	"github.com/micro-plat/sso/modules/const/enum"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/operate"
	"github.com/micro-plat/sso/modules/user"
)

type UserEditHandler struct {
	container component.IContainer
	userLib   user.IUser
	member    member.IMember
	op        operate.IOperate
}

func NewUserEditHandler(container component.IContainer) (u *UserEditHandler) {
	return &UserEditHandler{
		container: container,
		userLib:   user.NewUser(container),
		member:    member.NewMember(container),
		op:        operate.NewOperate(container),
	}
}

//GetHandle 编辑个人基本资料
func (u *UserEditHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑个人基本资料--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("username", "tel", "email"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	if err := u.userLib.Edit(ctx.Request.GetString("username"), ctx.Request.GetString("tel"), ctx.Request.GetString("email")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.记录行为")
	if err := u.op.UserOperate(
		member.Query(ctx, u.container),
		"编辑个人资料",
		"username",
		ctx.Request.GetInt("username"),
		"tel",
		ctx.Request.GetInt("tel"),
		"email",
		ctx.Request.GetString("email"),
	); err != nil {
		return err
	}
	ctx.Log.Info("4.返回结果")
	return "success"
}

//PostHandle 编辑用户详细资料（包括系统数据）
func (u *UserEditHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------编辑用户详细资料（包括系统数据）--------")
	ctx.Log.Info("1.参数校验")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}

	var input user.UserInputNew
	if err := ctx.Request.Bind(&input); err != nil && input.UserID == 0 {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("3.执行操作")
	if err := u.userLib.Save(&input); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	// 判断是否需要发送邮件
	b, err := u.userLib.IsSendEmail(&input)
	// 发邮件
	if err == nil && b == true {
		guid := utility.GetGUID()
		conf := app.GetConf(u.container)
		resURI := url.QueryEscape(fmt.Sprintf(conf.GetBindUrl(), guid))
		ctx.Log.Infof("发送验证邮件到:%s,guid：%v", input.Email, guid)
		link := fmt.Sprintf(enum.WxApiCode, resURI)
		if err := u.member.SendCheckMail(enum.From, enum.Password, enum.Host, enum.Port, input.Email, link); err != nil {
			return err
		}
		if err := u.userLib.SetEmail(guid, input.Email); err != nil {
			return err
		}
	}

	ctx.Log.Info("4.返回结果")
	return "success"
}
