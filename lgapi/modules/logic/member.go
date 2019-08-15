package logic

import (
	"net/http"
	"strings"
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/utility"

	"github.com/micro-plat/sso/lgapi/modules/access/member"
	"github.com/micro-plat/sso/lgapi/modules/const/enum"
	"github.com/micro-plat/sso/lgapi/modules/model"
)

//IMemberLogic 用户登录
type IMemberLogic interface {
	CreateLoginUserCode(userID int64) (code string, err error)
	Login(u, p, ident string) (*model.LoginState, error)
	ChangePwd(userID int, expassword string, newpassword string) (err error)
	CheckHasRoles(userID int64, ident string) error
}

//MemberLogic 用户登录管理
type MemberLogic struct {
	c     component.IContainer
	cache member.ICacheMember
	db    member.IDBMember
	http  *http.Client
}

//NewMemberLogic 创建登录对象
func NewMemberLogic(c component.IContainer) *MemberLogic {
	config := model.GetConf(c)
	return &MemberLogic{
		c:     c,
		cache: member.NewCacheMember(c),
		db:    member.NewDBMember(c),
		http:  &http.Client{Timeout: time.Duration(config.SendCodeTimeOut) * time.Second},
	}
}

//CreateLoginUserCode 验证用户是否已登录
func (m *MemberLogic) CreateLoginUserCode(userID int64) (code string, err error) {
	guid := utility.GetGUID()
	if err = m.cache.CreateUserInfoByCode(guid, userID); err != nil {
		return "", err
	}
	return guid, nil
}

//Login 登录系统
func (m *MemberLogic) Login(u, p, ident string) (s *model.LoginState, err error) {
	var ls *model.MemberState
	if ls, err = m.db.Query(u, p, ident); err != nil {
		return nil, err
	}

	if strings.ToLower(ls.Password) != strings.ToLower(p) {
		return nil, context.NewError(context.ERR_BAD_REQUEST, "用户名或密码错误")
	}

	//检查用户是否已锁定
	if ls.Status == enum.UserLock || ls.Status == enum.UserDisable {
		return nil, context.NewError(context.ERR_BAD_REQUEST, "用户被锁定或被禁用，暂时无法登录")
	}

	return (*model.LoginState)(ls), err
}

// ChangePwd 修改密码
func (m *MemberLogic) ChangePwd(userID int, expassword string, newpassword string) (err error) {
	return m.db.ChangePwd(userID, expassword, newpassword)
}

//CheckHasRoles jiancha daqian yong hu jiaoshe
func (m *MemberLogic) CheckHasRoles(userID int64, ident string) error {
	user, err := m.db.QueryByID(userID)
	if err != nil {
		return err
	}

	status := user.GetInt("status")
	if status == enum.UserLock || status == enum.UserDisable {
		return context.NewError(context.ERR_LOCKED, "用户被锁定或被禁用，暂时无法登录")
	}

	return m.db.CheckUserHasAuth(ident, userID)
}
