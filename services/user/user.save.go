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
	"github.com/micro-plat/sso/modules/user"
)

type UserSaveHandler struct {
	container component.IContainer
	userLib   user.IUser
	member    member.IMember
}

func NewUserSaveHandler(container component.IContainer) (u *UserSaveHandler) {
	return &UserSaveHandler{
		container: container,
		userLib:   user.NewUser(container),
		member:    member.NewMember(container),
	}
}

func (u *UserSaveHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑/添加用户信息--------")
	ctx.Log.Info("1.参数校验")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}

	var inputData user.UserEditInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.权限验证")
	// 修改 会验证权限
	if inputData.IsAdd == 0 {
		if err := u.member.QueryAuth(int64(l.SystemID), inputData.UserID); err != nil {
			return err
		}
	}
	ctx.Log.Info("3.执行操作")
	if err := u.userLib.Save(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	//新添加用户要进行邮箱检验
	if inputData.IsAdd == 1 {
		guid := utility.GetGUID()
		conf := app.GetConf(u.container)
		resUri := url.QueryEscape(fmt.Sprintf(conf.GetBindUrl(), guid))
		ctx.Log.Infof("发送验证邮件到:%s,guid：%v", inputData.Email, guid)
		link := fmt.Sprintf(enum.WxApiCode, resUri)
		if err := u.member.SendCheckMail(enum.From, enum.Password, enum.Host, enum.Port, inputData.Email, link); err != nil {
			return err
		}
		if err := u.userLib.SetEmail(guid, inputData.Email); err != nil {
			return err
		}
	}
	ctx.Log.Info("4.返回结果")
	return "success"
}
