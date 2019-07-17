package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"github.com/jordan-wright/email"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"

	"github.com/micro-plat/sso/mgrapi/modules/access/member"
	"github.com/micro-plat/sso/mgrapi/modules/const/enum"
	"github.com/micro-plat/sso/mgrapi/modules/model"
	"github.com/micro-plat/sso/mgrapi/modules/util"
)

//IMember 用户登录
type IMemberLogic interface {
	Login(u string, p string, ident string) (*model.LoginState, error) //在用
	LoginNew(code string) (*model.LoginState, error)
	QueryUserInfo(u string, ident string) (info db.QueryRow, err error) //在用
	Query(uid int64) (db.QueryRow, error)
	CacheQuery(u string, ident string) (ls *model.MemberState, err error)
	LoginByOpenID(string, string) (*model.LoginState, error)
	SendCheckMail(from string, password string, host string, port string, to string, link string) error
	QueryAuth(sysID, userID int64) (err error)
	QueryRoleByNameAndIdent(name, password, ident string) (s *model.MemberState, err error)
	SaveLoginStateToCache(s *model.MemberState) error
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
	return &MemberLogic{
		c:     c,
		cache: member.NewCacheMember(c),
		db:    member.NewDBMember(c),
		http:  &http.Client{},
	}
}

//Query 查询用户信息
func (m *MemberLogic) Query(uid int64) (db.QueryRow, error) {
	return m.db.QueryByID(uid)
}

// CacheQuery xx
func (m *MemberLogic) CacheQuery(userName string, ident string) (ls *model.MemberState, err error) {
	return m.cache.Query(userName, ident)
}

// SaveLoginStateToCache xx
func (m *MemberLogic) SaveLoginStateToCache(s *model.MemberState) error {
	return m.cache.Save(s)
}

//QueryRoleByNameAndIdent xx
func (m *MemberLogic) QueryRoleByNameAndIdent(name, password, ident string) (s *model.MemberState, err error) {
	return m.db.Query(name, password, ident)
}

// QueryAuth xx
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
		return nil, context.NewError(context.ERR_LENGTH_REQUIRED, "用户被禁用请联系管理员")
	}
	//检查密码是否有效，无效时累加登录失败次数
	if strings.ToLower(ls.Password) != strings.ToLower(p) {
		v, _ := m.cache.SetLoginFail(u)
		return nil, context.NewError(context.ERR_PRECONDITION_FAILED, fmt.Sprintf("用户名或密码错误:%d", v))
	}
	//设置登录成功
	err = m.cache.SetLoginSuccess(u)
	return (*model.LoginState)(ls), err
}

//LoginNew 跳转登录
func (m *MemberLogic) LoginNew(code string) (*model.LoginState, error) {
	config := model.GetConf(m.c)

	pars := map[string]interface{}{
		"ident":     config.Ident,
		"timestamp": time.Now().Unix(),
		"code":      code,
		"sysid":     0,
	}
	_, sign := util.MakeSign(pars, config.Secret)
	pars["sign"] = sign

	return m.remoteLogin(config, pars)
}

func (m *MemberLogic) remoteLogin(config *model.Conf, pars map[string]interface{}) (*model.LoginState, error) {

	url := fmt.Sprintf(
		"%s?ident=%s&timestamp=%d&code=%s&sign=%s&sysid=%d",
		config.GetUserInfoCode(),
		pars["ident"],
		pars["timestamp"],
		pars["code"],
		pars["sign"],
		pars["sysid"])

	resp, err := m.http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求:%v失败(%v)", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("登录失败,HttpStatus:%d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取远程数据失败 %s %v", url, err)
	}

	var state model.LoginState
	err = json.Unmarshal(body, &state)
	if err != nil {
		return nil, fmt.Errorf("解析返回结果失败 %s：%v(%s)", url, err, string(body))
	}

	return &state, nil
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
