package logic

import (
	"strings"

	"github.com/Owen-Zhang/base64Captcha"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/access/login"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/const/enum"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model/config"
)

type ILoginLogic interface {
	CheckUserIsLocked(mobile string) error
	Login(mobile, password, ident string) (*model.LoginState, error)
	QueryUserMenu(uid int64, ident string) ([]map[string]interface{}, error)
	GetSystemInfo(ident string) (s db.QueryRow, err error)
	GenerateVerifyCode(mobile string) (string, error)
	CheckVerifyCode(mobile, verifyCode string) error
	ChangePwd(userID int64, expassword, newpassword string) error
}

type LoginLogic struct {
	c     component.IContainer
	db    login.IDbLogin
	cache login.ICacheLogin
}

func NewLoginLogic(c component.IContainer) *LoginLogic {
	return &LoginLogic{
		c:     c,
		db:    login.NewDbLogin(c),
		cache: login.NewCacheLogin(c),
	}
}

//CheckUserIsLocked CheckUserIsLocked
func (l *LoginLogic) CheckUserIsLocked(mobile string) error {
	failCount := config.UserLoginFailCount
	count, err := l.cache.GetLoginFailCnt(mobile)
	if err != nil {
		return err
	}

	//用户是否被锁定
	if count < failCount {
		return nil
	}

	//解锁时间是否过期
	if exists := l.cache.ExistsUnLockTime(mobile); exists {
		return context.NewError(model.ERR_USER_LOCKED, "用户被锁定,请联系管理员")
	} else {
		if err := l.unLockUser(mobile); err != nil {
			return err
		}
	}
	return nil
}

//unLockUser 解锁用户
func (l *LoginLogic) unLockUser(mobile string) error {
	l.cache.SetLoginSuccess(mobile)
	return l.db.UnLock(mobile)
}

//Login Login
func (l *LoginLogic) Login(mobile, password, ident string) (*model.LoginState, error) {
	m, err := l.db.Query(mobile, password, ident)
	if err != nil {
		return nil, err
	}

	if err := l.checkUserInfo(mobile, password, m); err != nil {
		return nil, err
	}

	return (*model.LoginState)(m), nil
}

//CheckUserInfo 检查用户
func (l *LoginLogic) checkUserInfo(mobile, password string, state *model.MemberState) (err error) {
	if state.Status == enum.UserDisable {
		return context.NewError(model.ERR_USER_FORBIDDEN, "用户被禁用，请联系管理员")
	}
	if state.Status == enum.UserLock {
		return context.NewError(model.ERR_USER_LOCKED, "用户被锁定，请联系管理员")
	}

	if strings.ToLower(state.Password) == strings.ToLower(password) {
		l.cache.SetLoginSuccess(mobile)
		return nil
	}

	count, err := l.cache.SetLoginFail(mobile)
	if err != nil {
		return err
	}

	err = context.NewError(model.ERR_USER_PWDWRONG, "用户名或密码错误")
	if count < config.UserLoginFailCount {
		return err
	}

	//更新用户状态
	if err := l.db.UpdateUserStatus(state.UserID, enum.UserLock); err != nil {
		return err
	}
	//设置解锁过期时间
	if err := l.cache.SetUnLockTime(mobile, config.UserLockTime); err != nil {
		return err
	}

	return err
}

//QueryUserMenu 查询用户菜单
func (l *LoginLogic) QueryUserMenu(uid int64, ident string) ([]map[string]interface{}, error) {
	return l.db.QueryUserMenu(uid, ident)
}

//GetSystemInfo 获取系统信息
func (l *LoginLogic) GetSystemInfo(ident string) (s db.QueryRow, err error) {
	return l.db.GetSystemInfo(ident)
}

//GenerateVerifyCode 生成登录图片验证码
func (l *LoginLogic) GenerateVerifyCode(mobile string) (string, error) {
	var configD = base64Captcha.ConfigDigit{
		Height:     60,
		Width:      150,
		MaxSkew:    0.5,
		DotCount:   50,
		CaptchaLen: 5,
	}
	_, captcaInterfaceInstance, verifyCode := base64Captcha.GenerateCaptcha("", configD)
	base64blob := base64Captcha.CaptchaWriteToBase64Encoding(captcaInterfaceInstance)

	if err := l.cache.SaveLoginVerifyCode(mobile, verifyCode, config.VerifyCodeTimeOut); err != nil {
		return "", err
	}
	return base64blob, nil
}

//CheckVerifyCode 验证用户输入验证码
func (l *LoginLogic) CheckVerifyCode(mobile, verifyCode string) error {
	vc, err := l.cache.GetLoginVerifyCode(mobile)
	if err != nil {
		return err
	}
	if !strings.EqualFold(verifyCode, vc) {
		return context.NewError(model.ERR_VALIDATECODE_WRONG, "验证码错误或者过期")
	}
	return nil
}

//ChangePwd 修改密码
func (l *LoginLogic) ChangePwd(userID int64, expassword, newpassword string) error {
	return l.db.ChangePwd(userID, expassword, newpassword)
}
