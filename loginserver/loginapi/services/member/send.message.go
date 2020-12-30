// +build sms

package member

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/lib4dev/vcs"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/common/module/model"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/logic"
)

//SendMessageHandler 发送短信验证码
type SendMessageHandler struct {
	mem logic.IMemberLogic
}

//NewSendCodeHandler 发送短信验证码
func NewSendCodeHandler() (u *SendMessageHandler) {
	return &SendMessageHandler{
		mem: logic.NewMemberLogic(),
	}
}

//Handle 发送短信验证码
func (u *SendMessageHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------发送短息验证码--------")

	ctx.Log().Info("1.验证参数")
	if err := ctx.Request().Check("username"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.验证用户是否存在")
	userInfo, err := u.mem.ValidUserAndGetUserInfo(ctx.Request().GetString("username"))
	if err != nil {
		return err
	}
	conf := model.GetConf()

	//获取短信发送配置
	smsConf, err := ctx.APPConf().GetServerConf().GetSubConf("app")
	if err != nil {
		return
	}

	params := &vcs.SendRequest{
		ReqID:      strconv.FormatInt(int64(components.Def.UUID()), 10),
		Ident:      ctx.Request().GetString("ident"),
		PhoneNo:    userInfo.GetString("mobile"),
		TemplateID: smsConf.GetString("sms_temp_id"),
		Keywords:   fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)),
	}
	fmt.Println(params, "params:%+v")
	fmt.Println(ctx.Request().GetString("ident"), "ident")
	ctx.Log().Info("3.发送短信")
	vcs.SetConfig(vcs.WithSmsSendUrl(conf.WithSmsSendURL))
	if _, err := vcs.SendSmsCode(ctx, params); err != nil {
		return err
	}
	return "success"

}
