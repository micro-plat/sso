package user

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/user"
	"github.com/micro-plat/sso/modules/app"
	"github.com/micro-plat/wechat/mp/oauth2"
	"github.com/micro-plat/lib4go/net/http"
)

type UserBindHandler struct {
	container component.IContainer
	userLib   user.IUser
}

func NewUserBindHandler(container component.IContainer) (u *UserBindHandler) {
	return &UserBindHandler{
		container: container,
		userLib:   user.NewUser(container),
	}
}

func (u *UserBindHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------绑定用户邮箱--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("email","code");err != nil {
		return err
	}
	email := ctx.Request.GetString("email")
	code := ctx.Request.GetString("code")

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
