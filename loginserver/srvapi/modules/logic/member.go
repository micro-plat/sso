package logic

import (
	"fmt"

	"github.com/Owen-Zhang/base64Captcha"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/sso/errorcode"

	"github.com/micro-plat/sso/loginserver/srvapi/modules/access/member"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/const/config"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/const/enum"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/model"
)

//IMemberLogic 用户登录
type IMemberLogic interface {
	QueryUserInfo(u string, ident string) (info db.QueryRow, err error)
	GetUserInfoByCode(code, ident string) (res *model.LoginState, err error)
	QueryUserSystem(userID int, ident string) (s db.QueryRows, err error)
	QueryAllUserInfo(source string, sourceID string) (s db.QueryRows, err error)
	GetAllUserInfoByUserRole(userID int, ident string) (string, error)
	GetRoleMenus(roleID int, ident string) (types.XMaps, error)
	GenerateVerifyCode(userName string) (string, error)
}

//MemberLogic 用户登录管理
type MemberLogic struct {
	cache member.ICacheMember
	db    member.IDBMember
}

//NewMemberLogic 创建登录对象
func NewMemberLogic() *MemberLogic {
	return &MemberLogic{
		cache: member.NewCacheMember(),
		db:    member.NewDBMember(),
	}
}

//QueryUserSystem 查询用户可用的子系统
func (m *MemberLogic) QueryUserSystem(userID int, ident string) (s db.QueryRows, err error) {
	return m.db.QueryUserSystem(userID, ident)
}

//QueryAllUserInfo 获取所有用户信息
func (m *MemberLogic) QueryAllUserInfo(source string, sourceID string) (s db.QueryRows, err error) {
	return m.db.QueryAllUserInfo(source, sourceID)
}

// QueryUserInfo 返回用户信息
func (m *MemberLogic) QueryUserInfo(u string, ident string) (ls db.QueryRow, err error) {
	if ls, err = m.db.QueryByUserName(u, ident); err != nil {
		return nil, err
	}

	if ls.GetInt("status") == enum.UserLock {
		return nil, errs.NewError(errorcode.ERR_USER_LOCKED, "用户被锁定暂时无法登录")
	}

	if ls.GetInt("status") == enum.UserDisable {
		return nil, errs.NewError(errorcode.ERR_USER_FORBIDDEN, "用户被禁用请联系管理员")
	}
	return ls, err
}

// GetUserInfoByCode 根据Code查询登录的用户信息
func (m *MemberLogic) GetUserInfoByCode(code, ident string) (res *model.LoginState, err error) {

	userStr, err := m.cache.GetUserInfoByCode(code)
	if err != nil || userStr == "" {
		return nil, errs.NewError(errorcode.ERR_LOGIN_ERROR, fmt.Sprintf("没有登录记录,请先登录,err:%v", err))
	}

	userID := types.GetInt(userStr, -1)
	if userID == -1 {
		return nil, errs.NewError(errorcode.ERR_LOGIN_ERROR, "登录出错，请重新登录")
	}

	m.cache.DeleteInfoByCode(code)
	userTemp, err := m.db.QueryByID(userID, ident)
	if err != nil {
		return nil, err
	}

	status := userTemp.Status
	if status == enum.UserLock {
		return nil, errs.NewError(errorcode.ERR_USER_LOCKED, "用户被锁定暂时无法登录")
	}
	if status == enum.UserDisable {
		return nil, errs.NewError(errorcode.ERR_USER_FORBIDDEN, "用户被禁用请联系管理员")
	}

	return (*model.LoginState)(userTemp), nil
}

//GetAllUserInfoByUserRole 获取和当前用户同一个角色的用户ids
func (m *MemberLogic) GetAllUserInfoByUserRole(userID int, ident string) (string, error) {
	return m.db.GetAllUserInfoByUserRole(userID, ident)
}

//GetRoleMenus 获取同一个角色的菜单列表
func (m *MemberLogic) GetRoleMenus(roleID int, ident string) (types.XMaps, error) {
	return m.db.GetRoleMenus(roleID, ident)
}

//GenerateVerifyCode 生成验证码,此处userName可能为电话号码
func (m *MemberLogic) GenerateVerifyCode(userName string) (string, error) {
	var configD = base64Captcha.ConfigDigit{
		Height:     60,
		Width:      150,
		MaxSkew:    0.5,
		DotCount:   50,
		CaptchaLen: 5,
	}
	_, captcaInterfaceInstance, verifyCode := base64Captcha.GenerateCaptcha("", configD)
	base64blob := base64Captcha.CaptchaWriteToBase64Encoding(captcaInterfaceInstance)

	if err := m.cache.SaveLoginVerifyCode(userName, verifyCode, config.VerifyCodeTimeOut); err != nil {
		return "", err
	}
	return base64blob, nil
}
