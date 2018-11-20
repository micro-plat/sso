package user

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/net/http"
	"github.com/micro-plat/lib4go/utility"
	"github.com/micro-plat/sso/modules/app"
	"github.com/micro-plat/sso/modules/const/enum"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/operate"
	"github.com/micro-plat/sso/modules/system"
	"github.com/micro-plat/sso/modules/user"
	"github.com/micro-plat/wechat/mp/oauth2"
)

//UserHandler is
type UserHandler struct {
	container component.IContainer
	userLib   user.IUser
	sys       system.ISystem
	op        operate.IOperate
	member    member.IMember
}

//NewUserHandler is
func NewUserHandler(container component.IContainer) (u *UserHandler) {
	return &UserHandler{
		container: container,
		userLib:   user.NewUser(container),
		sys:       system.NewSystem(container),
		op:        operate.NewOperate(container),
		member:    member.NewMember(container),
	}
}

//GetHandle 获取用户所拥有的系统列表
func (u *UserHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------查询用户所有系统列表---------")
	userID := member.Get(ctx).UserID
	datas, err := u.sys.GetAll(userID)
	if err != nil {
		return err
	}
	ctx.Log.Info("返回数据")
	return datas
}

//PostHandle 查询用户信息数据
func (u *UserHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询用户信息数据--------")
	ctx.Log.Info("1.参数校验")
	var inputData user.QueryUserInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	rows, count, err := u.userLib.Query(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return map[string]interface{}{
		"count": count,
		"list":  rows,
	}
}

//PutHandle 修改用户状态
func (u *UserHandler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------修改用户状态--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("user_id", "status"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.userLib.ChangeStatus(ctx.Request.GetInt("user_id"), ctx.Request.GetInt("status")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.记录行为")
	if err := u.op.UserOperate(
		member.Query(ctx, u.container),
		"修改用户状态",
		"user_id",
		ctx.Request.GetInt("user_id"),
		"status",
		ctx.Request.GetInt("status"),
	); err != nil {
		return err
	}
	ctx.Log.Info("4.返回结果")
	return "success"
}

//DeleteHandle 删除用户
func (u *UserHandler) DeleteHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------删除用户--------")
	ctx.Log.Info("1.参数校验")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	if err := ctx.Request.Check("user_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.权限验证")
	if err := u.member.QueryAuth(int64(l.SystemID), ctx.Request.GetInt64("user_id")); err != nil {
		return err
	}
	ctx.Log.Info("3.执行操作")
	if err := u.userLib.Delete(ctx.Request.GetInt("user_id")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("4.记录行为")
	if err := u.op.UserOperate(
		member.Query(ctx, u.container),
		"删除用户",
		"user_id",
		ctx.Request.GetInt("user_id"),
	); err != nil {
		return err
	}
	ctx.Log.Info("5.返回结果")
	return "success"
}

//GetAllHandle 查询当前系统下用户列表
func (u *UserHandler) GetAllHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询当前系统下用户列表--------")
	ctx.Log.Info("1.参数校验")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	if err := ctx.Request.Check("pi", "ps"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	rows, count, err := u.userLib.GetAll(l.SystemID, ctx.Request.GetInt("pi"), ctx.Request.GetInt("ps"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return map[string]interface{}{
		"count": count,
		"list":  rows,
	}
}

//InfoHandle 查询用户信息
func (u *UserHandler) InfoHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询用户信息--------")
	ctx.Log.Info("1.参数校验")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	var uid int64
	err := ctx.Request.Check("user_id")
	if err != nil {
		uid = member.Get(ctx).UserID
	} else {
		uid = ctx.Request.GetInt64("user_id")
	}
	ctx.Log.Info("2.验证权限")
	if err = u.member.QueryAuth(int64(l.SystemID), uid); err != nil {
		return context.NewError(context.ERR_FORBIDDEN, err)
	}

	ctx.Log.Info("3.执行操作")
	data, err := u.userLib.Get(int(uid))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("4.返回结果")
	return map[string]interface{}{
		"userinfo": data,
	}
}

//EditHandle 编辑个人基本资料
func (u *UserHandler) EditHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑个人基本资料--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("user_name", "mobile", "email"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	if err := u.userLib.Edit(ctx.Request.GetString("user_name"), ctx.Request.GetString("mobile"), ctx.Request.GetString("email")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.记录行为")
	if err := u.op.UserOperate(
		member.Query(ctx, u.container),
		"编辑个人资料",
		"user_name",
		ctx.Request.GetString("user_name"),
		"mobile",
		ctx.Request.GetInt("mobile"),
		"email",
		ctx.Request.GetString("email"),
	); err != nil {
		return err
	}
	ctx.Log.Info("4.返回结果")
	return "success"
}

//EditDetailHandle 编辑用户详细资料（包括系统数据）
func (u *UserHandler) EditDetailHandle(ctx *context.Context) (r interface{}) {
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

//SaveHandle 添加用户信息
func (u *UserHandler) SaveHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加用户信息--------")
	ctx.Log.Info("1.参数校验")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	ctx.Log.Info("2.参数校验")
	var inputData user.UserInputNew
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("3.执行操作")
	if err := u.userLib.Add(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	// 判断是否需要发送邮件
	b, err := u.userLib.IsSendEmail(&inputData)
	// 发邮件
	if err == nil && b == true {
		guid := utility.GetGUID()
		conf := app.GetConf(u.container)
		resURI := url.QueryEscape(fmt.Sprintf(conf.GetBindUrl(), guid))
		ctx.Log.Infof("发送验证邮件到:%s,guid：%v", inputData.Email, guid)
		link := fmt.Sprintf(enum.WxApiCode, resURI)
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

//BindHandle 绑定用户邮箱
func (u *UserHandler) BindHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------绑定用户邮箱--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("guid", "code"); err != nil {
		return err
	}
	guid := ctx.Request.GetString("guid")
	code := ctx.Request.GetString("code")

	//判断邮箱是否过期
	email, err := u.userLib.GetEmail(guid)
	if err != nil || email == "" {
		return fmt.Errorf("链接已经过期，请联系管理员.错误：%v,邮箱：%v", err, email)
	}
	ctx.Log.Info("2. 根据code查询用户openid")

	conf := app.GetConf(u.container)
	endpoint := oauth2.NewEndpoint(conf.AppID, conf.Secret)
	url := endpoint.ExchangeTokenURL(code)
	client := http.NewHTTPClient()
	content, status, err := client.Get(url)
	if err != nil || status != 200 {
		return fmt.Errorf("远程请求失败:%s(%v)%d", url, err, status)
	}
	userInfo := make(db.QueryRow)
	if err = json.Unmarshal([]byte(content), &userInfo); err != nil {
		return fmt.Errorf("返回串无法解析:(%v)%s", err, content)
	}
	ctx.Log.Info("3. 根据openid进行用户绑定")
	openID := userInfo.GetString("openid")
	if err := u.userLib.Bind(email, openID); err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
