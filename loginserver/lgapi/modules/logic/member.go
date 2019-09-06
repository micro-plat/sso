package logic

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"math/rand"
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
	CheckSystemStatus(ident string) error
	CreateLoginUserCode(userID int64) (code string, err error)
	CheckUserIsLocked(userName string) error
	SendWxValidCode(userName, openID, ident string) error
	CheckWxValidCode(userName, wxCode string) error
	Login(userName, password, ident string) (*model.LoginState, error)
	ChangePwd(userID int, expassword string, newpassword string) (err error)
	CheckHasRoles(userID int64, ident string) error
	GenerateCodeAndSysInfo(ident string, userID int64) (map[string]string, error)
	CheckUerInfo(userID int64, sign, timestamp string) error
	GenerateWxStateCode(userID int64) (string, error)
	ValidStateAndGetOpenID(stateCode, wxCode string) (map[string]string,error)
	UpdateUserOpenID(data map[string]string) error
	ValidUserInfo(userName string) (string,error)
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

//CheckWxValidCode 验证微信验证码是否正确
func (m *MemberLogic) CheckWxValidCode(userName,wxCode string) error {
	if !model.GetConf(m.c).RequireWxCode {
		return nil
	}
	if strings.EqualFold(wxCode, "") {
		context.NewError(model.ERR_USER_EMPTY_VALIDATECODE, "验证码不能为空")
	}
	if err := m.cache.CheckLoginValidateCode(userName, wxCode);err != nil {
		return err
	}
	return nil
}

//CheckUserIsLocked 检查用户是否被锁定
func (m *MemberLogic) CheckUserIsLocked(userName string) error {
	failCount := model.GetConf(m.c).UserLoginFailCount
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
func (m *MemberLogic) Login(userName, password, ident string) (s *model.LoginState, err error) {
	var ls *model.MemberState
	if ls, err = m.db.Query(userName, password, ident); err != nil {
		return nil, err
	}

	if err = m.checkUserInfo(userName, password, ls); err != nil {
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
	if strings.EqualFold(ident, "") {
		return nil
	}
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
func (m *MemberLogic) checkUserInfo(userName, password string, state *model.MemberState) (err error) {
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

	conf := model.GetConf(m.c)
	err = context.NewError(model.ERR_USER_PWDWRONG, "用户名或密码错误")
	if count < conf.UserLoginFailCount {
		return err
	}

	//更新用户状态
	if err := m.db.UpdateUserStatus(state.UserID, enum.UserLock); err != nil {
		return err
	}
	//设置解锁过期时间
	if err := m.cache.SetUnLockTime(userName, conf.UserLockTime); err != nil {
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
	values := net.NewValues()
	values.Set("user_id", string(userID))
	values.Set("timestamp", timestamp)
	values = values.Sort()
	raw := values.Join("", "") + model.WxBindSecrect
	if !strings.EqualFold(sign, md5.Encrypt(raw)) {
		return context.NewError(model.ERR_BIND_INFOWRONG, "绑定信息错误,请重新去用户系统扫码")
	}
	sendTime, _ := strconv.ParseInt(timestamp, 10, 64)
	if time.Now().Unix()- sendTime > int64(model.GetConf(m.c).BindTimeOut) {
		return context.NewError(model.ERR_QRCODE_TIMEOUT, "二维码过期,请联系管理员重新生成")
	}

	data, err := m.db.QueryByID(userID)
	if err != nil {
		return err
	}
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
	if err := m.cache.SaveWxStateCode(stateCode, strconv.FormatInt(userID,10)); err != nil {
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

//UpdateUserOpenID 保存用户的openid
func (m *MemberLogic) UpdateUserOpenID(data map[string]string) error {
	return m.db.UpdateUserOpenID(data)
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

//ValidUserInfo 验证用户信息openid
func (m *MemberLogic) ValidUserInfo(userName string) (string, error) {
	datas, err := m.db.GetUserInfo(userName)
	if err != nil {
		return "",err
	}
	if datas.IsEmpty() {
		return "", context.NewError(model.ERR_USER_NOTEXISTS, "用户不存在")
	}
	if datas.Get(0).GetString("wx_openid") == "" {
		return "", context.NewError(model.ERR_USER_NOTBINDWX, "用户还未绑定微信账户")
	}
	return datas.Get(0).GetString("wx_openid"), nil
}

//SendWxValidCode 发送微信验证码
func (m *MemberLogic) SendWxValidCode(userName, openID, ident string) error {
	//1: 发送微信验证码
	token, err := m.GetFreshToken()
	if err != nil {
		return err
	}
	randd := rand.New(rand.NewSource(time.Now().UnixNano()))
	validcode := fmt.Sprintf("%06v", randd.Int31n(1000000))
	m.sendCode(openID, token, validcode, ident)

	//2:保存到redis中
	if err := m.cache.SetLoginValidateCode(validcode, userName); err != nil {
		return err
	}
	return nil
}

//GetFreshToken 动态获取token
func (m *MemberLogic) GetFreshToken() (string,error) {
	cfg := model.GetConf(m.c)
	url := fmt.Sprintf("%s/%s/wechat/token/get", cfg.RefreshWxTokenHost, cfg.WxAppID)
	client, err := http.NewHTTPClient(http.WithRequestTimeout(5 * time.Second))
	if err != nil {
		return "", err
	}
	body, statusCode, err := client.Get(url)
	if err != nil {
		return "", err
	}
	if statusCode != 200 {
		return "", context.NewErrorf(statusCode, "获取token信息失败,HttpStatus:%d, body:%s", statusCode, body)
	}

	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return "", fmt.Errorf("字符串转json发生错误，err：%v", err)
	}

	if errcode, ok := data["errcode"]; ok && types.GetInt(errcode) != 0 {
		return "", context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("获取token失败: %s", types.GetString(data["errmsg"])))
	}

	return types.GetString(data["access_token"]), nil
}

//SendCode 调用微信接口发验证码
func (m *MemberLogic) sendCode(openID, accessToken, validCode, ident string) error {
	cfg := model.GetConf(m.c)
	data := m.constructSendData(openID,validCode,ident,cfg.LoginValidCodeTemplateID)
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s?access_token=%s", cfg.WxSendTemplateMsgURL, accessToken)
	client, err := http.NewHTTPClient(http.WithRequestTimeout(5 * time.Second))	
	content, statusCode, err := client.Post(url, string(dataJSON))

	if err != nil {
		return err
	}
	if statusCode != 200 {
		return context.NewErrorf(statusCode, "发送验证码信息失败,HttpStatus:%d, body:%s", statusCode, content)
	}

	sendResult := make(map[string]interface{})
	err = json.Unmarshal([]byte(content), &sendResult)
	if err != nil {
		return fmt.Errorf("字符串转json发生错误，err：%v", err)
	}

	if errcode, ok := sendResult["errcode"]; ok && errcode != 0 {
		return context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("发送验证码信息失败: %s", types.GetString(sendResult["errmsg"])))
	}

	return nil
}

//构造发送验证码的实体数据
func (m *MemberLogic) constructSendData(openID, validCode, ident, templateID string) model.TemplateMsg {
	sendTitle := "用户系统"
	if !strings.EqualFold(ident, "") {
		system, _ := m.sysDB.QuerySysInfoByIdent(ident)
		if system != nil {
			sendTitle = system.GetString("name")
		}
	}

	data := model.TemplateMsg{
		Touser: openID,
		TemplateID: templateID,
		Data: &model.TemplateData{
			First: model.KeyWordData{
				Value: fmt.Sprintf("你好,登录到[%s]需要进行验证码验证,请无泄露", sendTitle),
				Color: "#173177",
			},
			Keyword1: model.KeyWordData{
				Value: validCode,
				Color: "#173177",
			},
			Keyword2: model.KeyWordData{
				Value: "5分钟",
				Color: "#173177",
			},
			Keyword3: model.KeyWordData{
				Value: time.Now().Format("2006-01-02 15:04:05"),
				Color: "#173177",
			},
			Remark: model.KeyWordData{
				Value: "若非本人操作,可能你的账号存在安全风险,请及时修改密码",
				Color: "#173177",
			},
		},
	}
	return data
}
