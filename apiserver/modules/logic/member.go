package logic

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"

	"github.com/micro-plat/sso/apiserver/modules/access/member"
	"github.com/micro-plat/sso/apiserver/modules/const/enum"
	"github.com/micro-plat/sso/apiserver/modules/model"
)

//IMember 用户登录
type IMemberLogic interface {
	Login(u string, p string, ident string) (*model.LoginState, error)  //在用
	QueryUserInfo(u string, ident string) (info db.QueryRow, err error) //在用
	Query(uid int64) (db.QueryRow, error)
	LoginByOpenID(string, string) (*model.LoginState, error)
	SendCheckMail(from string, password string, host string, port string, to string, link string) error
	QueryAuth(sysID, userID int64) (err error)
}

//Member 用户登录管理
type MemberLogic struct {
	cache member.ICacheMember
	db    member.IDBMember
}

//NewMember 创建登录对象
func NewMemberLogic(c component.IContainer) *MemberLogic {
	return &MemberLogic{
		cache: member.NewCacheMember(c),
		db:    member.NewDBMember(c),
	}
}

//Query 查询用户信息
func (m *MemberLogic) Query(uid int64) (db.QueryRow, error) {
	return m.db.QueryByID(uid)
}

func (m *MemberLogic) QueryAuth(sysID, userID int64) (err error) {
	err = m.cache.QueryAuth(sysID, userID)
	if err != nil {
		data, err := m.db.QueryAuth(sysID, userID)
		if err != nil || data == nil {
			return err
		}
		if err = m.cache.SaveAuth(sysID, userID, data); err != nil {
			return err
		}
	}
	return nil
}

//SendCheckMail 发送确认邮件
func (m *MemberLogic) SendCheckMail(from string, password string, host string, port string, to string, link string) error {
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = "用户账户确认"
	e.HTML = []byte(fmt.Sprintf("<h1>欢迎!</h1><br/><h1>感谢您在本系统注册，请复制以下链接到微信打开以完成帐号绑定!</h1><br/><h1 style='color:bule'>%s</h1>", link))
	return e.Send(host+":"+port, smtp.PlainAuth("", from, password, host))
}

//LoginByOpenID 使用open_id进行登录
func (m *MemberLogic) LoginByOpenID(openid string, ident string) (s *model.LoginState, err error) {
	row, err := m.db.QueryByOpenID(openid)
	if err != nil {
		return nil, err
	}
	u := row.GetString("user_name")
	p := row.GetString("password")
	return m.Login(u, p, ident)
}

//Login 登录系统
func (m *MemberLogic) Login(u string, p string, ident string) (s *model.LoginState, err error) {
	//从缓存中获取用户信息，不存在时从数据库中获取
	// ls, err := m.cache.Query(u, p, sys)
	// if ls == nil || err != nil {
	var ls *model.MemberState
	if ls, err = m.db.Query(u, p, ident); err != nil {
		return nil, err
	}
	// }
	//保存用户数据到缓存
	if err = m.cache.Save(ls); err != nil {
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
	//设置登录成功
	err = m.cache.SetLoginSuccess(u)
	return (*model.LoginState)(ls), err
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
