package logic

import (
	"fmt"
	"time"
	"strings"
	"encoding/json"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/utility"
	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/net/http"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/access/member"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/access/system"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/const/enum"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/model"
)

//IMemberLogic 用户登录
type IMemberLogic interface {
	CreateLoginUserCode(userID int64) (code string, err error)
	CheckUserIsLocked(userName string, failCount int) error
	Login(userName, password, ident string, failCount, userLockTime int) (*model.LoginState, error)
	ChangePwd(userID int, expassword string, newpassword string) (err error)
	CheckHasRoles(userID int64, ident string) error
	GenerateCodeAndSysInfo(ident string, userID int64) (map[string]string, error)
	CheckUerInfo(userID int64, sign, timestamp string) error
	GenerateWxStateCode(userID int64) (string, error)
	ValidStateAndGetOpenID(stateCode, wxCode string) (map[string]string,error)
	SaveUserOpenID(data map[string]string) error
}

//MemberLogic 用户登录管理
type MemberLogic struct {
	c     component.IContainer
	cache member.ICacheMember
	db    member.IDBMember
	sysDB system.IDbSystem
}

//NewMemberLogic 创建登录对象
func NewMemberLogic(c component.IContainer) *MemberLogic {
	return &MemberLogic{
		c:     c,
		cache: member.NewCacheMember(c),
		db:    member.NewDBMember(c),
		sysDB: system.NewDbSystem(c),
	}
}

//CheckUserIsLocked 检查用户是否被锁定
func (m *MemberLogic) CheckUserIsLocked(userName string, failCount int) error {
	count, err := m.cache.GetLoginFailCnt(userName)
	if err != nil {
		return err
	}

	//用户是否被锁定
	if count < failCount {
		return nil
	}

	//解锁时间是否过期
	if exists := m.cache.ExistsUnLockTime(userName); exists {
		return context.NewError(model.ERR_USER_LOCKED, "用户被锁定,请联系管理员")
	} else {
		if err := m.unLockUser(userName); err != nil {
			return err
		}
	}
	return nil
}

// ChangePwd 修改密码
func (m *MemberLogic) ChangePwd(userID int, expassword string, newpassword string) (err error) {
	return m.db.ChangePwd(userID, expassword, newpassword)
}

//Login 登录系统
func (m *MemberLogic) Login(userName, password, ident string, failCount, userLockTime int) (s *model.LoginState, err error) {
	if !strings.EqualFold(ident, "") {
		if err := m.CheckSystemStatus(ident); err != nil {
			return nil, err
		}
	}

	var ls *model.MemberState
	if ls, err = m.db.Query(userName, password, ident); err != nil {
		return nil, err
	}

	if err = m.checkUserInfo(userName, password, ls, failCount, userLockTime); err != nil {
		return nil, err
	}

	return (*model.LoginState)(ls), err
}

//GenerateCodeAndSysInfo 生成登录后的Code
func (m *MemberLogic) GenerateCodeAndSysInfo(ident string, userID int64) (map[string]string, error) {
	if strings.EqualFold(ident, "") {
		return map[string]string{}, nil
	}

	code, err := m.CreateLoginUserCode(userID)
	if err != nil {
		return nil, context.NewError(context.ERR_BAD_REQUEST, err)
	}

	sysInfo, err := m.sysDB.QuerySysInfoByIdent(ident)
	if err != nil {
		return nil, err
	}

	return map[string]string{"code": code, "callback": sysInfo.GetString("index_url")}, nil
}

//CheckSystemStatus 检查系统的状态
func (m *MemberLogic) CheckSystemStatus(ident string) error {
	data, err := m.sysDB.QuerySysInfoByIdent(ident)
	if err != nil {
		return err
	}
	if data.GetInt("enable") == enum.SystemDisable {
		return context.NewError(model.ERR_SYS_LOCKED, "系统被禁用,不能登录")
	}
	return nil
}

//CheckUserInfo 检查用户
func (m *MemberLogic) checkUserInfo(userName, password string, state *model.MemberState, failCount, userLockTime int) (err error) {
	if state.Status == enum.UserDisable {
		return context.NewError(model.ERR_USER_FORBIDDEN, "用户被禁用，请联系管理员")
	}
	if state.Status == enum.UserLock {
		return context.NewError(model.ERR_USER_LOCKED, "用户被锁定，请联系管理员")
	}

	if strings.ToLower(state.Password) == strings.ToLower(password) {
		m.cache.SetLoginSuccess(userName)
		return nil
	}

	count, err := m.cache.SetLoginFail(userName)
	if err != nil {
		return err
	}

	err = context.NewError(model.ERR_USER_PWDWRONG, "用户名或密码错误")
	if count < failCount {
		return err
	}

	//更新用户状态
	if err := m.db.UpdateUserStatus(state.UserID, enum.UserLock); err != nil {
		return err
	}
	//设置解锁过期时间
	if err := m.cache.SetUnLockTime(userName, userLockTime); err != nil {
		return err
	}

	return err
}

//CreateLoginUserCode 生成用户登录的标识保存到缓存中(code)
func (m *MemberLogic) CreateLoginUserCode(userID int64) (code string, err error) {
	guid := utility.GetGUID()
	if err = m.cache.CreateUserInfoByCode(guid, userID); err != nil {
		return "", err
	}
	return guid, nil
}

//CheckHasRoles 检查用户是否有相应的角色
func (m *MemberLogic) CheckHasRoles(userID int64, ident string) error {
	if err := m.CheckSystemStatus(ident); err != nil {
		return err
	}

	user, err := m.db.QueryByID(userID)
	if err != nil {
		return err
	}

	status := user.GetInt("status")
	if status == enum.UserLock {
		return context.NewError(model.ERR_USER_LOCKED, "用户被锁定,暂时无法登录")
	}
	if status == enum.UserDisable {
		return context.NewError(model.ERR_USER_FORBIDDEN, "用户被禁用，暂时无法登录")
	}

	return m.db.CheckUserHasAuth(ident, userID)
}

//unLockUser 解锁用户
func (m *MemberLogic) unLockUser(userName string) error {
	m.cache.SetLoginSuccess(userName)
	return m.db.UnLock(userName)
}

//CheckUerInfo 验证要绑定的用户信息
func (m *MemberLogic) CheckUerInfo(userID int64, sign, timestamp string) error {
	//验证签名
	values := net.NewValues()
	values.Set("user_id", string(userID))
	values.Set("timestamp", timestamp)

	values = values.Sort()
	raw := values.Join("", "") + model.WxBindSecrect
	if !strings.EqualFold(sign, md5.Encrypt(raw)) {
		return context.NewError(model.ERR_BIND_INFOWRONG, "绑定信息错误,请重新去用户系统扫码")
	}

	//查询用户
	data, err := m.db.QueryByID(userID)
	if err != nil {
		return err
	}

	//验证用户状态
	status := data.GetInt("status")
	if status == enum.UserLock {
		return context.NewError(model.ERR_USER_LOCKED, "用户被禁用")
	}
	if status == enum.UserDisable {
		return context.NewError(model.ERR_USER_LOCKED, "用户被锁定")
	}

	if data.GetString("wx_openid") != "" {
		return context.NewError(model.ERR_USER_EXISTSWX, "用户已绑定微信")
	}

	return nil
}

//GenerateWxStateCode 生成微信stateCode凭证
func (m *MemberLogic) GenerateWxStateCode(userID int64) (string, error) {
	stateCode := utility.GetGUID()
	if err := m.cache.SaveWxStateCode(stateCode, string(userID)); err != nil {
		return "", err
	}
	return stateCode, nil
}

//ValidStateAndGetOpenID 验证state信息并获取openid
func (m *MemberLogic) ValidStateAndGetOpenID(stateCode, wxCode string) (map[string]string, error) {
	userID, err := m.cache.GetWxStateCodeUserId(stateCode)
	if err != nil {
		return nil,err
	}
	if userID == "" {
		return nil, context.NewError(model.ERR_BIND_TIMEOUT, "绑定超时")
	}

	config := model.GetConf(m.c)
	url := config.WxTokenURL + "?appid=" + config.WxAppID + "&secret=" + config.WxSecret + "&code=" + wxCode + "&grant_type=authorization_code"
	openID, err := m.GetWxUserOpID(url)
	if err != nil {
		return nil, err
	}
	return map[string]string {
		"openid": openID,
		"userid": userID,
	}, nil
}

//SaveUserOpenID 保存用户的openid
func (m *MemberLogic) SaveUserOpenID(data map[string]string) error {
	return m.db.SaveUserOpenID(data)
}

// GetWxUserOpID xx
func (m *MemberLogic) GetWxUserOpID(url string) (string, error) {
	fmt.Println(url)
	client, err := http.NewHTTPClient(http.WithRequestTimeout(5 * time.Second))
	if err != nil {
		return "", err
	}
	body, statusCode, err := client.Get(url)
	if err != nil {
		return "", err
	}
	if statusCode != 200 {
		return "", context.NewErrorf(statusCode, "读取系统信息失败,HttpStatus:%d, body:%s", statusCode, body)
	}

	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return "", fmt.Errorf("字符串转json发生错误，err：%v", err)
	}

	//wx返回全是200,只有通过errcode去判断
	if errcode, ok := data["errcode"]; ok && errcode != 0 {
		return "", context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("微信返回错误：%s", types.GetString(data["errmsg"])))
	}

	return types.GetString(data["openid"]), nil
}
