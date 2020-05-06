package sso

import (
	"fmt"
	"net/url"
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
	raw := values.Join("", "") + u.cfg.secret
	values.Set("sign", md5.Encrypt(raw))

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
	raw := values.Join("", "") + u.cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	lgState := &LoginState{}
	result, err := remoteRequest(u.cfg.host, codeLoginUrl, values.Join("=", "&"), lgState)
	if err != nil {
		return nil, err
	}
	return result.(*LoginState), nil
}

//GetUserMenu 查询用户在某个系统下的菜单数据
func (u *userLogic) getUserMenu(userID int) ([]Menu, error) {
	values := net.NewValues()
	values.Set("user_id", types.GetString(userID))
	values.Set("ident", u.cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("", "") + u.cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	var other []Menu
	_, err := remoteRequest(u.cfg.host, userMenuUrl, values.Join("=", "&"), &other)
	if err != nil {
		return nil, err
	}
	return other, nil
}

//getUserSystems 返回用户可用的子系统列表(有权限,除当前系统外)
func (u *userLogic) getUserOtherSystems(userID int) (*[]*System, error) {
	values := net.NewValues()
	values.Set("user_id", types.GetString(userID))
	values.Set("ident", u.cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("", "") + u.cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	sysList := &[]*System{}
	result, err := remoteRequest(u.cfg.host, userSysUrl, values.Join("=", "&"), sysList)
	if err != nil {
		return nil, err
	}
	return result.(*[]*System), nil
}

//GetAllUser 返回所有正常用户
func (u *userLogic) GetAllUser(source string, sourceID string) (*[]*User, error) {
	values := net.NewValues()
	values.Set("ident", u.cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))
	values.Set("source", source)
	values.Set("source_id", sourceID)

	values = values.Sort()
	raw := values.Join("", "") + u.cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	userList := &[]*User{}
	result, err := remoteRequest(u.cfg.host, userAllUrl, values.Join("=", "&"), userList)
	if err != nil {
		return nil, err
	}
	return result.(*[]*User), nil
}

//ForgetPwd 忘记密码并修改密码
func (u *userLogic) ForgetPwd(source, sourceID, password string) error {
	newpwd, _ := url.QueryUnescape(password)
	values := net.NewValues()
	values.Set("ident", u.cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))
	values.Set("source", source)
	values.Set("source_id", sourceID)
	values.Set("password", newpwd)

	values = values.Sort()
	raw := values.Join("", "") + u.cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	// fmt.Println("newpwd:", newpwd)
	// fmt.Println("ecode:", values.Encode())

	var res map[string]interface{}
	_, err := remoteRequest(u.cfg.host, forgetPassword, values.Encode(), res)
	if err != nil {
		return err
	}
	return nil
}

//GetUserTags 获取用户有权限的tags
func (u *userLogic) GetUserTags(userID int) ([]Menu, error) {
	return getUserTagFromLocal(userID)
}

//getRoleUsers 获取与某个角色关联的所有用户
func (u *userLogic) getRoleUsers(userID int64) (userIds string, err error) {
	cfg := u.cfg
	values := net.NewValues()
	values.Set("ident", cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))
	values.Set("user_id", types.GetString(userID))

	values = values.Sort()
	raw := values.Join("", "") + cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	result := make(map[string]string)
	_, err = remoteRequest(cfg.host, roleAllUser, values.Encode(), &result)
	if err != nil {
		return "", err
	}
	return result["data"], nil
}

//AddUser 增加用户
func (u *userLogic) AddUser(userName, mobile, fullName, targetIdent, source, sourceSecrect string, sourceID string) error {
	cfg := u.cfg
	values := net.NewValues()
	values.Set("ident", targetIdent)
	values.Set("timestamp", types.GetString(time.Now().Unix()))
	values.Set("mobile", mobile)
	values.Set("user_name", userName)
	values.Set("full_name", fullName)
	values.Set("target_ident", targetIdent)
	values.Set("source", source)
	values.Set("source_id", sourceID)

	values = values.Sort()
	raw := values.Join("", "") + sourceSecrect
	values.Set("sign", md5.Encrypt(raw))

	result := make(map[string]string)
	_, err := remoteRequest(cfg.host, addUser, values.Encode(), &result)
	if err != nil {
		return err
	}
	return nil
}

//Login 用户名密码登录
func (u *userLogic) Login(userName, password string) (LoginState, error) {
	values := net.NewValues()
	values.Set("user_name", userName)
	values.Set("password", password)
	values.Set("ident", u.cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("", "") + u.cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	lgState := LoginState{}
	fmt.Println(values.Encode(), "login-values")
	_, err := remoteRequest(u.cfg.host, passwordLogin, values.Encode(), &lgState)
	if err != nil {
		return LoginState{}, err
	}
	return lgState, nil
}

//ChangePwd 修改密码
func (u *userLogic) ChangePwd(userID int64, expassword, newpassword string) error {

	oldpwd, _ := url.QueryUnescape(expassword)
	newpwd, _ := url.QueryUnescape(newpassword)
	values := net.NewValues()
	values.Set("user_id", types.GetString(userID))
	values.Set("expassword", oldpwd)
	values.Set("newpassword", newpwd)
	values.Set("ident", u.cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("", "") + u.cfg.secret
	values.Set("sign", md5.Encrypt(raw))
	// fmt.Println("request:", values.Encode())

	var res map[string]interface{}
	_, err := remoteRequest(u.cfg.host, changePassword, values.Encode(), &res)
	if err != nil {
		return err
	}
	return nil
}
