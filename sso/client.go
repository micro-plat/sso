package sso

import (
	"net/http"
	"strings"

	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/sso/errorcode"
)

//Client sso client
type Client struct {
	cfg *ConfigData
}

//New SSOClient
func New(apiHost, ident, secret string) (*Client, error) {
	cfg := &ConfigData{
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

//GetRolePageTags 获取有权限的Tags
func (client *Client) GetRolePageTags(roleID int, pageURL, tags string) (result types.XMap, err error) {
	tagInput := strings.Split(tags, ",")

	authorityData, err := getRoleTagFromLocal(roleID)
	if err != nil {
		err = errs.NewErrorf(http.StatusForbidden, "获取缓存失败：%d,ident:%s,role:%d", client.cfg.ident, roleID)
		return
	}

	//页面权限
	item, ok := authorityData[pageURL]
	if !ok {
		err = errs.NewErrorf(errorcode.ERR_USER_HASNOPAGEAUTHORITY, "用户没有相应的页面权限")
		return
	}
	pageTags := item.FuncTags
	result = types.XMap{}
	for _, tag := range tagInput {

		if len(pageTags) == 0 {
			result[tag] = true
			continue
		}
		result[tag] = pageTags[tag]
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
