package member

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/sso/flowserver/modules/app"
	"github.com/micro-plat/sso/flowserver/modules/system"

	"github.com/micro-plat/wechat/mp/message/template"
)

type IWxcode interface {
	Check(un string, code string) error
	GetWXCode() string
	Send(un string, sysIdent string, appid string, secret string, serverAddr string, code string) error
}

type Wxcode struct {
	c         component.IContainer
	db        IDBMember
	sys       system.ISystem
	cacheTime int
}

const (
	wxCodeCacheFormat = "sso:login:wx-valid-code:{@userName}"
)

func NewWxcode(c component.IContainer) *Wxcode {
	return &Wxcode{
		c:         c,
		db:        NewDBMember(c),
		sys:       system.NewSystem(c),
		cacheTime: 3600 * 24,
	}
}

//GetWXCode 发送微信验证码
func (l *Wxcode) GetWXCode() string {
	rand.Seed(time.Now().UnixNano())
	var num string
	for i := 0; i < 4; i++ {
		x := rand.Intn(10)
		num = fmt.Sprintf("%s%d", num, x)
	}
	return num
}

//Check 验证微信验证码
func (l *Wxcode) Check(un string, code string) error {
	key := transform.Translate(wxCodeCacheFormat, "userName", un)
	cache := l.c.GetRegularCache()
	ccode, err := cache.Get(key)
	if err != nil {
		return err
	}
	defer cache.Delete(key)
	if ccode != code {
		return context.NewError(901, fmt.Errorf("微信验证码错误"))
	}
	return nil
}

//Send 发送微信验证码
func (l *Wxcode) Send(un string, sysIdent string, appid string, secret string, serverAddr string, code string) error {
	row, err := l.db.GetUserInfo(un)
	if err != nil || row.GetString("wx_openid") == "" {
		return context.NewError(406, err)
	}

	sys, err := l.sys.Get(sysIdent)
	if err != nil {
		return context.NewError(406, err)
	}

	ctx := app.GetWeChatContext(l.c)
	if _, err := template.Send(ctx, &template.TemplateMessage2{
		ToUser:     row.GetString("wx_openid"),
		TemplateId: "_DL41WrU7r6uNYyjD45c5B11ECkOAhwdDG8qqQxbvGs",
		Data: map[string]interface{}{
			"first": map[string]string{
				"value": fmt.Sprintf("您正在登录[%s]", sys.GetString("name")),
			},
			"keyword1": map[string]string{
				"value": code,
				"color": "#43CD80",
			},
			"keyword2": map[string]string{
				"value": "5分钟",
			},
			"keyword3": map[string]string{
				"value": time.Now().Format("2006/01/02 15:04:05"),
			},
			"remark": map[string]string{
				"value": "若非本人操作请注意账户安全",
			},
		},
	}); err != nil {
		return fmt.Errorf("发送验证码失败:%v(%s)%v", err, row.GetString("wx_openid"), serverAddr)
	}
	key := transform.Translate(wxCodeCacheFormat, "userName", un)
	cache := l.c.GetRegularCache()
	if err := cache.Set(key, code, l.cacheTime); err != nil {
		return fmt.Errorf("保存到缓存失败:%v", err)
	}
	return nil
}
