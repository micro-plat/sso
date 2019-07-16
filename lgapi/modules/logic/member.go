package logic

import (
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/utility"

	"github.com/micro-plat/sso/lgapi/modules/access/member"
	"github.com/micro-plat/sso/lgapi/modules/const/enum"
	"github.com/micro-plat/sso/lgapi/modules/model"
)

//IMemberLogic 用户登录
type IMemberLogic interface {
	CreateLoginUserCode(user_id int64) (code string, err error)
	Login(u string, p string) (*model.LoginState, error)
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

//CreateLoginUserCode 验证用户是否已登录
func (m *MemberLogic) CreateLoginUserCode(user_id int64) (code string, err error) {
	guid := utility.GetGUID()
	if err = m.cache.CreateUserInfoByCode(guid, user_id); err != nil {
		return "", err
	}
	return guid, nil
}

//Login 登录系统
func (m *MemberLogic) Login(u string, p string) (s *model.LoginState, err error) {
	var ls *model.MemberState
	if ls, err = m.db.Query(u, p); err != nil {
		return nil, err
	}

	if strings.ToLower(ls.Password) != strings.ToLower(p) {
		return nil, context.NewError(context.ERR_UNAUTHORIZED, "用户名或密码错误")
	}

	//检查用户是否已锁定
	if ls.Status == enum.UserLock || ls.Status == enum.UserDisable {
		return nil, context.NewError(context.ERR_LOCKED, "用户被锁定或被禁用，暂时无法登录")
	}

	return (*model.LoginState)(ls), err
}
