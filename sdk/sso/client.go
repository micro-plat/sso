package sso

import (
	"strings"

	"github.com/micro-plat/lib4go/types"
)

//Client sso client
type Client struct {
	cfg *Config
}

//New SSOClient
func New(apiHost, ident, secret string) (*Client, error) {
	cfg := &Config{
		host:   apiHost,
		ident:  ident,
		secret: secret,
	}
	if err := cfg.Valid(); err != nil {
		return nil, err
	}
	return &Client{cfg: cfg}, nil
}

//CheckCodeLogin 验证回传code并获取登录用户信息
func (client *Client) CheckCodeLogin(code string) (res *LoginState, err error) {
	u := newUser(client.cfg)
	return u.checkCodeLogin(code)
}

//GetUserInfoByName 根据用户名获取用户信息
func (client *Client) GetUserInfoByName(userName string) (info *User, err error) {
	u := newUser(client.cfg)
	return u.getUserInfoByName(userName)
}

//GetUserMenu 获取用户菜单信息
func (client *Client) GetUserMenu(userID int) ([]Menu, error) {
	u := newUser(client.cfg)
	return u.getUserMenu(userID)
}

//GetSystemInfo 获取系统信息
func (client *Client) GetSystemInfo() (data *System, err error) {
	s := NewSystem(client.cfg)
	return s.getSystemInfo()
}

//GetUserOtherSystems 获取用户可用的其他子系统
func (client *Client) GetUserOtherSystems(userID int) (*[]*System, error) {
	s := newUser(client.cfg)
	return s.getUserOtherSystems(userID)
}

//GetAllUser 获取所有用户信息
func (client *Client) GetAllUser(source string, sourceID string) (*[]*User, error) {
	s := newUser(client.cfg)
	return s.GetAllUser(source, sourceID)
}

//ForgetPwd 忘记密码并修改密码
func (client *Client) ForgetPwd(source, sourceID, possword string) error {
	s := newUser(client.cfg)
	return s.ForgetPwd(source, sourceID, possword)
}

//GetUserDisplayTags 获取用户有权限的Tags
func (client *Client) GetUserDisplayTags(UserID int, tags string) (result []types.XMap, err error) {
	tagInput := strings.Split(tags, ",")

	s := newUser(client.cfg)
	userHasTags, err := s.GetUserTags(UserID)
	if err != nil {
		return nil, err
	}

	for _, tag := range tagInput {
		detail := types.XMap{"tag": tag, "display": false}
		for _, temp := range userHasTags {
			if strings.EqualFold(strings.TrimSpace(tag), strings.TrimSpace(temp.Path)) {
				detail["display"] = true
				break
			}
		}
		result = append(result, detail)
	}
	return
}

//getUserDataPermission 获取用户的 数据权限管理　数据
func (client *Client) getUserDataPermission(userID int64, tableName string, opt ...PermissionOption) (r string, err error) {
	return newDataPermission(client.cfg).getUserDataPermission(userID, tableName, opt...)
}

//AddUser 增加用户
func (client *Client) AddUser(userName, mobile, fullName, targetIdent, source, sourceSecrect string, sourceID string, roleID int) error {
	return newUser(client.cfg).AddUser(userName, mobile, fullName, targetIdent, source, sourceSecrect, sourceID, roleID)
}

//Login 用户密码登录
func (client *Client) Login(userName, password string) (LoginState, error) {
	return newUser(client.cfg).Login(userName, password)
}

//ChangePwd 修改密码
func (client *Client) ChangePwd(userID int64, expassword, newpassword string) error {
	return newUser(client.cfg).ChangePwd(userID, expassword, newpassword)
}
