package service

import (
	"strings"
	"time"

	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/sso/model"
)

//GetUserInfoByName 根据用户名获取用户信息
func GetUserInfoByName(conf *model.Config, userName string) (info *User, err error) {
	values := net.NewValues()
	values.Set("username", userName)
	values.Set("ident", conf.Ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("=", "&") + "&key=" + conf.Secret
	values.Set("sign", strings.ToUpper(md5.Encrypt(raw)))

	user := &User{}
	result, err := remoteRequest(conf.Host, model.UserInfoUrl, values.Join("=", "&"), user)
	if err != nil {
		return nil, err
	}
	return result.(*User), nil
}

//CheckCodeLogin 验证用户登录的code
func CheckCodeLogin(conf *model.Config, code string) (res *LoginState, err error) {

	values := net.NewValues()
	values.Set("code", code)
	values.Set("ident", conf.Ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("=", "&") + "&key=" + conf.Secret
	values.Set("sign", strings.ToUpper(md5.Encrypt(raw)))

	lgState := &LoginState{}
	result, err := remoteRequest(conf.Host, model.CodeLoginUrl, values.Join("=", "&"), lgState)
	if err != nil {
		return nil, err
	}
	return result.(*LoginState), nil
}
