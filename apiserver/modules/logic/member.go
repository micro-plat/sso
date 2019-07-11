package logic

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"

	"github.com/micro-plat/sso/apiserver/modules/access/member"
	"github.com/micro-plat/sso/apiserver/modules/access/operate"
	"github.com/micro-plat/sso/apiserver/modules/const/enum"
	"github.com/micro-plat/sso/apiserver/modules/model"
	"github.com/micro-plat/sso/apiserver/modules/model/user"
)

//IMember 用户登录
type IMemberLogic interface {
	Login(u string, p string, ident string) (*model.LoginState, error)
	QueryUserInfo(u string, ident string) (info db.QueryRow, err error)
	GetUserInfoByKey(key string) (res *user.UserKeyResp, err error)
}

//MemberLogic 用户登录管理
type MemberLogic struct {
	cache  member.ICacheMember
	db     member.IDBMember
	dboper operate.IDbOperate
	dbcode member.ICodeMember
}

//NewMemberLogic 创建登录对象
func NewMemberLogic(c component.IContainer) *MemberLogic {
	return &MemberLogic{
		cache:  member.NewCacheMember(c),
		db:     member.NewDBMember(c),
		dboper: operate.NewDbOperate(c),
		dbcode: member.NewCodeMember(c),
	}
}

//Login 登录系统
func (m *MemberLogic) Login(u string, p string, ident string) (s *model.LoginState, err error) {
	var ls *model.MemberState
	if ls, err = m.db.Query(u, p, ident); err != nil {
		return nil, err
	}

	//检查用户是否已锁定
	if ls.Status == enum.UserLock {
		return nil, context.NewError(context.ERR_LOCKED, "用户被锁定暂时无法登录(423)")
	}
	//检查用户是否已禁用
	if ls.Status == enum.UserDisable {
		return nil, context.NewError(context.ERR_LENGTH_REQUIRED, "用户被禁用请联系管理员(411)")
	}
	//检查密码是否有效，无效时累加登录失败次数
	if strings.ToLower(ls.Password) != strings.ToLower(p) {
		v, _ := m.cache.SetLoginFail(u)
		return nil, context.NewError(context.ERR_PRECONDITION_FAILED, fmt.Sprintf("用户名或密码错误(412):%d", v))
	}

	//保存用户数据到缓存
	if err = m.cache.Save(ls); err != nil {
		return nil, err
	}

	//设置登录成功
	err = m.cache.SetLoginSuccess(u)
	member := (*model.LoginState)(ls)

	//保存用户信息
	_, err = m.dbcode.Save(member)
	if err != nil {
		return nil, err
	}

	//记录登录行为
	if err := m.dboper.LoginOperate(member); err != nil {
		return nil, err
	}

	return member, err
}

// QueryUserInfo 返回用户信息
func (m *MemberLogic) QueryUserInfo(u string, ident string) (ls db.QueryRow, err error) {

	if ls, err = m.db.QueryByUserName(u, ident); err != nil {
		return nil, err
	}
	//检查用户是否已锁定
	if ls.GetInt("status") == enum.UserLock {
		return nil, context.NewError(context.ERR_LOCKED, "用户被锁定暂时无法登录")
	}
	//检查用户是否已禁用
	if ls.GetInt("status") == enum.UserDisable {
		return nil, context.NewError(context.ERR_FORBIDDEN, "用户被禁用请联系管理员")
	}
	return ls, err
}

// GetUserInfoByKey 根据key查询登录的用户信息
func (m *MemberLogic) GetUserInfoByKey(key string) (res *user.UserKeyResp, err error) {
	userStr, err := m.cache.GetUserInfoByKey(key)

	if err != nil || userStr == "" {
		return nil, context.NewError(context.ERR_FORBIDDEN, fmt.Sprintf("没有登录记录,请先登录,err:%s", err))
	}

	userID := types.GetInt(userStr, -1)
	if userID == -1 {
		return nil, context.NewError(context.ERR_FORBIDDEN, "登录出错，请重新登录")
	}

	userTemp, err := m.db.QueryByID(userID)
	if err != nil {
		return nil, err
	}

	status := userTemp.GetInt("status")
	if status == enum.UserLock || status == enum.UserDisable {
		return nil, context.NewError(context.ERR_LOCKED, "用户被锁定或者被禁用")
	}

	return &user.UserKeyResp{
		UserId:   userTemp.GetInt("user_id"),
		UserName: userTemp.GetString("user_name"),
	}, nil
}
