package logic

import (
	"encoding/json"
	"fmt"
	"gsms/express_service/modules/util/security"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/logger"
	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
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
	CheckUerInfo(userName, password string) error
	SaveWxStateCode(code, content string) error
	ExistsWxStateCode(code string) (bool, error)
	GetWxLoginInfoByStateCode(stateCode string) (string, error)
	SaveWxLoginInfo(state, content string) error
	GetWxUserOpID(url string) (string, error)
	ExistsOpenId(content string) error
	GetUserInfoByOpID(opID, ident string) (*model.LoginState, error)
	GetSendUserByName(userName, ident string) (senduser string, err error)
	SendValidCode(userName, sendUser string) error
	ValidVerifyCode(userName, validatecode string) (bool, error)

	ValidStateAndGetOpenID(state, code string, logger logger.ILogger) (string, error)
	SaveUserOpenID(content, state string, logger logger.ILogger) error
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

//CheckUerInfo 验证用户名密码,以及openid
func (m *MemberLogic) CheckUerInfo(userName, password string) error {
	rows, err := m.db.GetUserInfo(userName)
	if err != nil {
		return context.NewError(context.ERR_SERVICE_UNAVAILABLE, err)
	}
	if rows.IsEmpty() {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "用户名或密码错误")
	}
	data := rows.Get(0)
	if strings.ToLower(data.GetString("password")) != strings.ToLower(md5.Encrypt(password)) {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "用户名或密码错误")
	}
	status := data.GetInt("status")
	if status == enum.UserLock || status == enum.UserDisable {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "用户被锁定或被禁用,请联系管理员")
	}
	if data.GetString("wx_openid") != "" {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "账号已绑定微信")
	}
	return nil
}

//SaveWxStateCode xx
func (m *MemberLogic) SaveWxStateCode(code, content string) error {
	return m.cache.SaveWxStateCode(code, content)
}

//ExistsWxStateCode 是否存在wx state code 防伪造
func (m *MemberLogic) ExistsWxStateCode(code string) (bool, error) {
	return m.cache.ExistsWxStateCode(code)
}

//SaveWxLoginInfo 保存微信登录的openid信息
func (m *MemberLogic) SaveWxLoginInfo(state, content string) error {
	return m.cache.SaveWxLoginInfo(state, content)
}

// GetWxLoginInfoByStateCode 根据statecode去cache中取登录openid信息
func (m *MemberLogic) GetWxLoginInfoByStateCode(stateCode string) (string, error) {
	content, err := m.cache.GetWxLoginInfoByStateCode(stateCode)
	fmt.Printf("cache get data: %s", content)
	if err != nil {
		return "", err
	}
	if content == "1" || content == "" {
		return "", nil
	}
	tokenInfo := make(map[string]interface{})
	err = json.Unmarshal([]byte(content), &tokenInfo)
	if err != nil {
		return "", context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("解析返回结果失败 %s：(%s)", content, err))
	}

	openID, ok := tokenInfo["openid"]
	if !ok || openID == "" {
		return "", context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("微信返回错误：%s", content))
	}
	return openID.(string), nil
}

// GetWxUserOpID xx
func (m *MemberLogic) GetWxUserOpID(url string) (string, error) {
	resp, err := m.http.Get(url)
	if err != nil {
		return "", context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("请求:%v失败(%v)", url, err))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("读取远程数据失败 %s %v", url, err))
	}

	if resp.StatusCode != 200 {
		return "", context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("登录失败,HttpStatus:%d, body:%s", resp.StatusCode, string(body)))
	}

	token := make(map[string]interface{})
	err = json.Unmarshal(body, &token)
	if err != nil {
		return "", context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("解析返回结果失败 %s：%v(%s)", url, err, string(body)))
	}

	//wx返回全是200,只有通过errcode去判断
	if errcode, ok := token["errcode"]; ok && errcode != 0 {
		return "", context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("微信返回错误：%s", token["errmsg"].(string)))
	}

	return string(body), nil
}

//ExistsOpenId xx
func (m *MemberLogic) ExistsOpenId(content string) error {
	return m.db.ExistsOpenId(content)
}

//GetUserInfoByOpID xxx
func (m *MemberLogic) GetUserInfoByOpID(opID, ident string) (*model.LoginState, error) {
	data, err := m.db.QueryByOpenID(opID, ident)
	if err != nil {
		return nil, err
	}
	return (*model.LoginState)(data), nil
}

//GetSendUserByName 获取要发送验证码的用户名
func (m *MemberLogic) GetSendUserByName(userName, ident string) (senduser string, err error) {
	data, err := m.db.QueryByName(userName, ident)
	if err != nil {
		return "", err
	}

	senduser = userName
	extParams := data.ExtParams
	extObj := map[string]string{}
	if strings.EqualFold(extParams, "") {
		return
	}
	if err = json.Unmarshal([]byte(extParams), &extObj); err != nil {
		return
	}
	if extSenduser, ok := extObj["senduser"]; ok && extSenduser != "" {
		senduser = extSenduser
	}
	return
}

//SendValidCode 发送验证码 (userName是用户登录的名字，senduser是发给公众号那边对应的)
func (m *MemberLogic) SendValidCode(userName, sendUser string) error {
	config := model.GetConf(m.c)

	systemno := "sso"
	key := config.SendCodeKey
	requrl := config.SendCodeReqUrl
	content := `你好，欢迎使用%s$$$登录验证码$$$%s$$$微信验证码有效时间5分钟，过期作废【安全提醒：请勿将微信验证码提供给他人】`
	systemName := "单点登录系统"
	title := "验证码下发"
	keyword := "验证码下发"

	randd := rand.New(rand.NewSource(time.Now().UnixNano()))
	validcode := fmt.Sprintf("%06v", randd.Int31n(1000000))

	if err := m.cache.CreateValiCode(userName, validcode); err != nil {
		return err
	}

	fmt.Println(validcode)

	content = fmt.Sprintf(content, systemName, validcode)
	vals := net.NewValues()
	params := map[string]string{
		"msg_type":  "6",
		"use_type":  "1",
		"send_user": sendUser,
		"title":     url.QueryEscape(title),
		"content":   url.QueryEscape(content),
		"key_words": url.QueryEscape(keyword),
		"system_no": systemno,
	}

	//md5(key+send_user+system_no+key)
	params["sign"] = security.Md5(key + sendUser + systemno + key)
	vals.SetSMap(params)
	queryParams := vals.Join("=", "&")
	requestURL := requrl + "?" + queryParams
	fmt.Println("sendValidcodeOut:", requestURL)

	resp, err := m.http.Get(requestURL)
	if err != nil {
		return context.NewError(context.ERR_NOT_EXTENDED, fmt.Sprintf("调用发送微信验证码的接口失败: %v+", err))
	}
	if resp.StatusCode != 200 {
		return context.NewError(context.ERR_NOT_EXTENDED, fmt.Sprintf("调用发送微信验证码的接口失败: 状态码:%d", resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return context.NewError(context.ERR_NOT_EXTENDED, fmt.Sprintf("调用发送微信验证码的接口失败: 返回内容为: %s", string(body)))
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		return context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("解析返回结果失败 %s", string(body)))
	}

	//wx返回全是200,只有通过errcode去判断
	if status, ok := result["Status"]; ok && status != "1000" {
		return context.NewError(context.ERR_BAD_REQUEST, result["Message"].(string))
	}
	return err
}

// ValidVerifyCode 验证通过公众号发的验证码
func (m *MemberLogic) ValidVerifyCode(userName, validatecode string) (bool, error) {
	return m.cache.VerifyValidCode(userName, validatecode)
}

//ValidStateAndGetOpenID 验证state并且获取openid
func (m *MemberLogic) ValidStateAndGetOpenID(state, code string, logger logger.ILogger) (string, error) {
	logger.Info("2:验证state code是否存在, 防止伪造")
	if flag, _ := m.ExistsWxStateCode(state); !flag {
		return "", context.NewError(context.ERR_REQUEST_TIMEOUT, fmt.Errorf("微信登录标识过期,请重新登录"))
	}

	logger.Info("3:调用wx接口,获取用户openid")
	config := model.GetConf(m.c)
	url := config.WxTokenUrl + "?appid=" + config.Appid + "&secret=" + config.Secret + "&code=" + code + "&grant_type=authorization_code"
	logger.Infof("获取用户openid的url: %s", url)

	content, err := m.GetWxUserOpID(url)
	if err != nil {
		logger.Errorf("调用wx api出错: %v+", err)
		return "", err
	}
	return content, err
}

//SaveUserOpenID 保存用户的openid
func (m *MemberLogic) SaveUserOpenID(content, state string, logger logger.ILogger) error {
	userName, err := m.cache.GetContentByStateCode(state)
	if err != nil || userName == "" {
		logger.Infof("获取缓存数据出错: %v+, %s", err, userName)
		return context.NewError(context.ERR_SERVER_ERROR, "出现错误，请稍后在绑定")
	}
	var token map[string]interface{}
	if err := json.Unmarshal([]byte(content), &token); err != nil {
		logger.Infof("转换数据出错: %v+, 内容: %s", err, content)
		return context.NewError(context.ERR_SERVER_ERROR, "出现错误，请稍后在绑定")
	}
	return m.db.SaveUserOpenID(userName, token["openid"].(string))
}
