package user

import (
	"github.com/micro-plat/sso/modules/app"
	"net/url"
	"fmt"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/user"
	"github.com/micro-plat/sso/modules/const/enum"
	"github.com/micro-plat/lib4go/utility"
)

type UserSaveHandler struct {
	container component.IContainer
	userLib   user.IUser
	member member.IMember
}

func NewUserSaveHandler(container component.IContainer) (u *UserSaveHandler) {
	return &UserSaveHandler{
		container: container,
		userLib:   user.NewUser(container),
		member: member.NewMember(container),
	}
}

func (u *UserSaveHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑/添加用户信息--------")
	ctx.Log.Info("1.参数校验")
	var inputData user.UserEditInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.userLib.Save(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	
	//新添加用户要进行邮箱检验
	if inputData.IsAdd == 1 {
		guid := utility.GetGUID()
		resUri := url.QueryEscape(fmt.Sprintf(app.GetBindUrl(u.container),guid))
		ctx.Log.Infof("发送验证邮件到:%s,guid：%v",inputData.Email,guid)
		link := fmt.Sprintf(enum.WxApiCode,resUri)
		if err := u.member.SendCheckMail(enum.From,enum.Password,enum.Host,enum.Port,inputData.Email,link); err != nil {
			return err
		}
		if err := u.userLib.SetEmail(guid,inputData.Email); err != nil {
			return err
		}
	}
	ctx.Log.Info("3.返回结果")
	return "success"
}
