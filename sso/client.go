package sso

import (
	"github.com/micro-plat/sso/sso/model"
	"github.com/micro-plat/sso/sso/service"
)

//SSOClient sso client
type SSOClient struct {
	cfg *model.Config
}

//New SSOClient
func New(apiHost, ident, secret string) (*SSOClient, error) {
	cfg := &model.Config{
		Host:   apiHost,
		Ident:  ident,
		Secret: secret,
	}
	if err := cfg.Valid(); err != nil {
		return nil, err
	}
	return &SSOClient{cfg: cfg}, nil
}

//CheckCodeLogin 验证回传code并获取登录用户信息
func (client *SSOClient) CheckCodeLogin(code string) (res *service.LoginState, err error) {
	return service.CheckCodeLogin(client.cfg, code)
}

//GetUserInfoByName 根据用户名获取用户信息
func (client *SSOClient) GetUserInfoByName(userName string) (info *service.User, err error) {
	return service.GetUserInfoByName(client.cfg, userName)
}

//GetUserMenu 获取用户菜单信息
func (client *SSOClient) GetUserMenu(userID int) (*[]*service.Menu, error) {
	return service.GetUserMenu(client.cfg, userID)
}

//GetSystemInfo 获取系统信息
func (client *SSOClient) GetSystemInfo() (data *service.System, err error) {
	return service.GetSystemInfo(client.cfg)
}
