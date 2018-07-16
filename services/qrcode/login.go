package qrcode

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/lib4go/db"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/net/http"
	"github.com/micro-plat/sso/modules/app"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/wechat/mp/oauth2"
	"github.com/micro-plat/wechat/mp/qrcode"
)

type LoginHandler struct {
	c    component.IContainer
	m    member.IMember
	code member.ICodeMember
}

func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c:    container,
		m:    member.NewMember(container),
		code: member.NewCodeMember(container),
	}
}

func (u *LoginHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.获取二维码登录地址")
	conf := app.GetConf(u.c)
	url := conf.QRLoginCheckURL
	uuid := ctx.Request.GetUUID()
	sysid := ctx.Request.GetString("sysid", "0")
	rt := fmt.Sprintf("%s?uid=%s&sysid=%s", url, uuid, sysid)
	rurl := oauth2.AuthCodeURL(conf.AppID, rt, "snsapi_base", "")
	ctx.Log.Info("2.实际登录地址:", rt)
	wectx := app.GetWeChatContext(u.c)
	surl, err := qrcode.ShortURL(wectx, rurl)
	if err != nil {
		return fmt.Errorf("生成短链接失败:%v", err)
	}
	return map[string]interface{}{
		"url": surl,
	}
}

//PostHandle 使用微信code查询用户openid,并登录，推送到ws端code
func (u *LoginHandler) PostHandle(ctx *context.Context) (r interface{}) {
	if err := ctx.Request.Check("uid", "code", "sysid"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("1. 根据code查询用户openid")
	sysid := ctx.Request.GetInt("sysid", 0)
	code := ctx.Request.GetString("code")
	conf := app.GetConf(u.c)
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
	ctx.Log.Info("2. 根据openid登录")
	openid := userInfo.GetString("openid")
	member, err := u.m.LoginByOpenID(openid, sysid)
	if err != nil {
		return fmt.Errorf("登录失败:(%v)%s", err, openid)
	}
	redirectURL := ctx.Request.GetString("redirect_uri")
	if redirectURL == "" {
		redirectURL = member.IndexURL
	}
	loginCode, err := u.code.Save(member)
	if err != nil {
		return fmt.Errorf("保存用户登录code失败:%v", err)
	}

	//设置jwt数据
	ctx.Log.Info("3. 通知登录端code")
	ctx.Response.SetJWT(member)
	context.WSExchange.Notify(ctx.Request.GetString("uid"), 200, "/qrcode/login/success", map[string]interface{}{
		"code":  loginCode,
		"sysid": sysid,
	})
	return "success"
}

//PutHandle 提交用户登录结果信息
func (u *LoginHandler) PutHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1. 用户已扫码")
	if err := ctx.Request.Check("uid"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2. 通知用户使用手机扫码")
	context.WSExchange.Notify(ctx.Request.GetString("uid"), 200, "/qrcode/login/check", "success")
	return "success"
}
