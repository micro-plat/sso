package sso

import (
	"fmt"
	"strings"
	"time"

	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
)

type userLogic struct {
	cfg *Config
}

func newUser(cfg *Config) *userLogic {
	return &userLogic{
		cfg: cfg,
	}
}

//GetUserInfoByName 根据用户名获取用户信息
func (u *userLogic) getUserInfoByName(userName string) (info *User, err error) {
	values := net.NewValues()
	values.Set("username", userName)
	values.Set("ident", u.cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("=", "&") + "&key=" + u.cfg.secret
	values.Set("sign", strings.ToUpper(md5.Encrypt(raw)))

	user := &User{}
	result, err := remoteRequest(u.cfg.host, userInfoUrl, values.Join("=", "&"), user)
	if err != nil {
		return nil, err
	}
	return result.(*User), nil
}

//CheckCodeLogin 验证用户登录的code
func (u *userLogic) checkCodeLogin(code string) (res *LoginState, err error) {

	values := net.NewValues()
	values.Set("code", code)
	values.Set("ident", u.cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("=", "&") + "&key=" + u.cfg.secret
	values.Set("sign", strings.ToUpper(md5.Encrypt(raw)))

	lgState := &LoginState{}
	result, err := remoteRequest(u.cfg.host, codeLoginUrl, values.Join("=", "&"), lgState)
	if err != nil {
		return nil, err
	}
	return result.(*LoginState), nil
}

//GetUserMenu 查询用户在某个系统下的菜单数据
func (u *userLogic) getUserMenu(userID int) (*[]*Menu, error) {
	values := net.NewValues()
	values.Set("user_id", types.GetString(userID))
	values.Set("ident", u.cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("=", "&") + "&key=" + u.cfg.secret
	fmt.Println(raw)
	values.Set("sign", strings.ToUpper(md5.Encrypt(raw)))

	menu := &[]*Menu{}
	result, err := remoteRequest(u.cfg.host, userMenuUrl, values.Join("=", "&"), menu)
	if err != nil {
		return nil, err
	}
	return result.(*[]*Menu), nil
}
