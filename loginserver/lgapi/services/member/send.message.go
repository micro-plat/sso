// +build sms

package member

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/common/module/model"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/logic"
	"gitlab.100bm.cn/micro-plat/vcs/vcs"
)

//SendMessageHandler 发送短信验证码
type SendMessageHandler struct {
	c   component.IContainer
	mem logic.IMemberLogic
}

//NewSendCodeHandler 发送短信验证码
func NewSendCodeHandler(container component.IContainer) (u *SendMessageHandler) {
	return &SendMessageHandler{
		c:   container,
		mem: logic.NewMemberLogic(container),
	}
}

//Handle 发送短信验证码
func (u *SendMessageHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------发送短息验证码--------")

	ctx.Log.Info("1.验证参数")
	if err := ctx.Request.Check("username"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.验证用户是否存在")
	userInfo, err := u.mem.ValidUserAndGetUserInfo(ctx.Request.GetString("username"))
	if err != nil {
		return err
	}
	conf := model.GetConf(u.c)

	//获取短信发送配置
	smsConf, err := u.c.GetSubConf("app")
	if err != nil {
		return
	}

	params := &vcs.SendRequest{
		ReqID:      strconv.FormatInt(context.NewUUID(u.c).Get(), 10),
		Ident:      ctx.Request.GetString("ident"),
		PhoneNo:    userInfo.GetString("mobile"),
		TemplateID: smsConf.GetString("sms_temp_id"),
		Keywords:   fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)),
	}
	fmt.Println(params, "params:%+v")
	fmt.Println(ctx.Request.GetString("ident"), "ident")
	ctx.Log.Info("3.发送短信")
	vcs.SetConfig(vcs.WithSmsSendUrl(conf.WithSmsSendURL))
	if _, err := vcs.SendSmsCode(ctx, params); err != nil {
		return err
	}
	return "success"

}
