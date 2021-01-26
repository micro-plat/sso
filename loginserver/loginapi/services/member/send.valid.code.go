package member

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/const/enum"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/logic"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/model"
	"github.com/micro-plat/sso/sso/errorcode"
)

//SendValidCodeHandler 发送验证码
type SendValidCodeHandler struct {
	mem       logic.IMemberLogic
	validcode *logic.ValidCodeLogic
}

//NewSendCodeHandler 发送短信验证码
func NewSendCodeHandler() (u *SendValidCodeHandler) {
	return &SendValidCodeHandler{
		mem:       logic.NewMemberLogic(),
		validcode: logic.NewValidCodeLogic(),
	}
}

//Handle 发送短信验证码
func (u *SendValidCodeHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------发送验证码--------")

	ctx.Log().Info("1.验证参数")
	if err := ctx.Request().Check("username"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.验证用户是否存在：", ctx.Request().GetString("username"))
	userInfo, err := u.mem.GetUserInfo(ctx.Request().GetString("username"))
	if err != nil {
		err = errs.NewError(errorcode.ERR_USER_NOTEXISTS, err)
		return err
	}
	ctx.Log().Info("3.发送验证码")
	randd := rand.New(rand.NewSource(time.Now().UnixNano()))
	validCode := fmt.Sprintf("%06v", randd.Int31n(1000000))

	conf := model.GetLoginConf()
	ident := ctx.Request().GetString("ident")
	switch conf.ValidCodeType {
	case enum.ValidCodeTypeSMS:
		err = u.validcode.SendSmsCode(userInfo, ident, validCode)
	case enum.ValidCodeTypeWechat:
		err = u.validcode.SendWechatCode(userInfo, ident, validCode)
	default:
		err = errs.NewError(errorcode.ERR_VALID_CODE_TYPE_ERROR, fmt.Errorf("无效的ValidCodeType:%s", conf.ValidCodeType))
	}
	if err != nil {
		return err
	}
	ctx.Log().Info("4.发送成功")
	return "success"

}
