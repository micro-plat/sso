package member

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

const (
	UserNormal int = iota
	UserLock
	UserDisable
)

//IMember 用户登录
type IMember interface {
	Login(u string, p string, sys int) (*LoginState, error)
	Query(uid int64) (db.QueryRow, error)
	LoginByOpenID(string, int) (*LoginState, error)
}

//Member 用户登录管理
type Member struct {
	cache ICacheMember
	db    IDBMember
}

//NewMember 创建登录对象
func NewMember(c component.IContainer) *Member {
	return &Member{
		cache: NewCacheMember(c),
		db:    NewDBMember(c),
	}
}

//Query 查询用户信息
func (m *Member) Query(uid int64) (db.QueryRow, error) {
	return m.db.QueryByID(uid)
}

//SendCheckMail 发送确认邮件
func (m *Member) SendCheckMail(from string, password string, host string, port string, to string, link string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("系统管理员<%s>", from)
	e.To = []string{to}
	e.Subject = "用户账户确认"
	e.HTML = []byte(fmt.Sprintf("<h1>%s!</h1>", link))
	return e.Send(host+":"+port, smtp.PlainAuth("", from, password, host))
}

//LoginByOpenID 使用open_id进行登录
func (m *Member) LoginByOpenID(openid string, sys int) (s *LoginState, err error) {
	row, err := m.db.QueryByOpenID(openid)
	if err != nil {
		return nil, err
	}
	u := row.GetString("user_name")
	p := row.GetString("password")
	return m.Login(u, p, sys)
}

//Login 登录系统
func (m *Member) Login(u string, p string, sys int) (s *LoginState, err error) {
	//从缓存中获取用户信息，不存在时从数据库中获取
	// ls, err := m.cache.Query(u, p, sys)
	// if ls == nil || err != nil {
	var ls *MemberState
	if ls, err = m.db.Query(u, p, sys); err != nil {
		return nil, err
	}
	// }
	//保存用户数据到缓存
	if err = m.cache.Save(ls); err != nil {
		return nil, err
	}
	//检查用户是否已锁定
	if ls.Status == UserLock {
		return nil, context.NewError(context.ERR_LOCKED, "用户被锁定暂时无法登录")
	}
	//检查用户是否已禁用
	if ls.Status == UserDisable {
		return nil, context.NewError(context.ERR_FORBIDDEN, "用户被禁用请联系管理员")
	}
	//检查密码是否有效，无效时累加登录失败次数
	if strings.ToLower(ls.Password) != strings.ToLower(p) {
		v, _ := m.cache.SetLoginFail(u)
		return nil, context.NewError(context.ERR_FORBIDDEN, fmt.Sprintf("用户名或密码错误:%d", v))
	}
	//设置登录成功
	err = m.cache.SetLoginSuccess(u)
	return (*LoginState)(ls), err
}
