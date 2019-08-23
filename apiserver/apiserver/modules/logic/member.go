package logic

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"

	"github.com/micro-plat/sso/apiserver/apiserver/modules/access/member"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/const/enum"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

//IMemberLogic 用户登录
type IMemberLogic interface {
	QueryUserInfo(u string, ident string) (info db.QueryRow, err error)
	GetUserInfoByCode(code, ident string) (res *model.LoginState, err error)
	QueryUserSystem(userID int, ident string) (s db.QueryRows, err error)
}

//MemberLogic 用户登录管理
type MemberLogic struct {
	cache member.ICacheMember
	db    member.IDBMember
}

//NewMemberLogic 创建登录对象
func NewMemberLogic(c component.IContainer) *MemberLogic {
	return &MemberLogic{
		cache: member.NewCacheMember(c),
		db:    member.NewDBMember(c),
	}
}

//QueryUserSystem 查询用户可用的子系统
func (m *MemberLogic) QueryUserSystem(userID int, ident string) (s db.QueryRows, err error) {
	return m.db.QueryUserSystem(userID, ident)
}

// QueryUserInfo 返回用户信息
func (m *MemberLogic) QueryUserInfo(u string, ident string) (ls db.QueryRow, err error) {
	if ls, err = m.db.QueryByUserName(u, ident); err != nil {
		return nil, err
	}

	if ls.GetInt("status") == enum.UserLock {
		return nil, context.NewError(model.ERR_USER_LOCKED, "用户被锁定暂时无法登录")
	}

	if ls.GetInt("status") == enum.UserDisable {
		return nil, context.NewError(model.ERR_USER_FORBIDDEN, "用户被禁用请联系管理员")
	}
	return ls, err
}

// GetUserInfoByCode 根据Code查询登录的用户信息
func (m *MemberLogic) GetUserInfoByCode(code, ident string) (res *model.LoginState, err error) {

	userStr, err := m.cache.GetUserInfoByCode(code)
	if err != nil || userStr == "" {
		return nil, context.NewError(model.ERR_LOGIN_ERROR, fmt.Sprintf("没有登录记录,请先登录,err:%s", err))
	}

	userID := types.GetInt(userStr, -1)
	if userID == -1 {
		return nil, context.NewError(model.ERR_LOGIN_ERROR, "登录出错，请重新登录")
	}

	m.cache.DeleteInfoByCode(code)
	userTemp, err := m.db.QueryByID(userID, ident)
	if err != nil {
		return nil, err
	}

	status := userTemp.Status
	if status == enum.UserLock {
		return nil, context.NewError(model.ERR_USER_LOCKED, "用户被锁定暂时无法登录")
	}
	if status == enum.UserDisable {
		return nil, context.NewError(model.ERR_USER_FORBIDDEN, "用户被禁用请联系管理员")
	}

	return (*model.LoginState)(userTemp), nil
}
