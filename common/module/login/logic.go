package login

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/common/module/const/enum"
	"github.com/micro-plat/sso/common/module/model"
	"github.com/micro-plat/sso/common/module/system"
)

//LoginLogic 用户登录相关
type LoginLogic struct {
	cache ICacheMember
	db    IDBMember
	sysDB system.IDbSystem
}

//NewLoginLogic 创建登录对象
func NewLoginLogic() *LoginLogic {
	return &LoginLogic{
		cache: NewCacheMember(),
		db:    NewDBMember(),
		sysDB: system.NewDbSystem(),
	}
}

//CheckWxValidCode 验证微信验证码是否正确
func (m *LoginLogic) CheckWxValidCode(userName, wxCode string) error {
	conf := model.GetConf()
	if !conf.RequireWxCode {
		return nil
	}
	if strings.EqualFold(wxCode, "") {
		errs.NewError(model.ERR_USER_EMPTY_VALIDATECODE, "验证码不能为空")
	}
	if err := m.cache.CheckLoginValidateCode(userName, wxCode); err != nil {
		return err
	}
	return nil
}

//CheckUserIsLocked 检查用户是否被锁定
func (m *LoginLogic) CheckUserIsLocked(userName string) error {
	conf := model.GetConf()
	failCount := conf.UserLoginFailCount
	count, err := m.cache.GetLoginFailCnt(userName)
	if err != nil {
		return err
	}

	//用户是否被锁定
	if count <= failCount {
		return nil
	}

	//解锁时间是否过期
	if exists := m.cache.ExistsUnLockTime(userName); exists {
		return errs.NewError(model.ERR_USER_LOCKED, "用户被锁定,请联系管理员")
	} else {
		if err := m.unLockUser(userName); err != nil {
			return err
		}
	}
	return nil
}

// ChangePwd 修改密码
func (m *LoginLogic) ChangePwd(userID int, expassword string, newpassword string) (err error) {
	data, err := m.db.QueryUserOldPwd(userID)
	if err != nil {
		return err
	}
	userInfo, err := m.db.QueryByID(userID)
	if err != nil {
		return err
	}

	if strings.EqualFold(strings.ToLower(md5.Encrypt(expassword)), strings.ToLower(data.Get(0).GetString("password"))) {
		m.cache.SetLoginSuccess(userInfo.GetString("user_name"))
		return m.db.ChangePwd(userID, newpassword)
	}

	conf := model.GetConf()
	count, err := m.cache.SetLoginFail(userInfo.GetString("user_name"))
	if err != nil {
		return err
	}
	if count <= conf.UserLoginFailCount {
		errorCode := m.generateWrongPwdErrorCode(count)
		err = errs.NewError(errorCode, "原密码错误，通过errorcode区分次数")
		return err
	}

	//更新用户状态
	if err := m.db.UpdateUserStatus(userID, enum.UserLock); err != nil {
		return err
	}
	//设置解锁过期时间
	if err := m.cache.SetUnLockTime(userInfo.GetString("user_name"), conf.UserLockTime); err != nil {
		return err
	}
	return errs.NewError(model.ERR_USER_LOCKED, "用户被锁定，请联系管理员")
}

//Login 登录系统
func (m *LoginLogic) Login(userName, password, ident string) (s *model.LoginState, err error) {
	var ls *model.MemberState
	if ls, err = m.db.Query(userName, password, ident); err != nil {
		return nil, err
	}

	if err = m.checkUserInfo(userName, password, ls); err != nil {
		return nil, err
	}

	return (*model.LoginState)(ls), err
}

//CheckSystemStatus 检查系统的状态
func (m *LoginLogic) CheckSystemStatus(ident string) error {
	if strings.EqualFold(ident, "") {
		return nil
	}
	data, err := m.sysDB.QuerySysInfoByIdent(ident)
	if err != nil {
		return err
	}
	if data.GetInt("enable") == enum.SystemDisable {
		return errs.NewError(model.ERR_SYS_LOCKED, "系统被禁用,不能登录")
	}
	return nil
}

//CheckUserInfo 检查用户
func (m *LoginLogic) checkUserInfo(userName, password string, state *model.MemberState) (err error) {
	if state.Status == enum.UserDisable {
		return errs.NewError(model.ERR_USER_FORBIDDEN, "用户被禁用，请联系管理员")
	}
	if state.Status == enum.UserLock {
		return errs.NewError(model.ERR_USER_LOCKED, "用户被锁定，请联系管理员")
	}

	if strings.ToLower(state.Password) == strings.ToLower(password) {
		m.cache.SetLoginSuccess(userName)
		return nil
	}

	count, err := m.cache.SetLoginFail(userName)
	if err != nil {
		return err
	}
	conf := model.GetConf()
	if count <= conf.UserLoginFailCount {
		errorCode := m.generateWrongPwdErrorCode(count)
		err = errs.NewError(errorCode, "用户名或密码错误，通过errorcode区分次数")
		return err
	}

	//更新用户状态
	if err := m.db.UpdateUserStatus(int(state.UserID), enum.UserLock); err != nil {
		return err
	}
	//设置解锁过期时间
	if err := m.cache.SetUnLockTime(userName, conf.UserLockTime); err != nil {
		return err
	}

	return errs.NewError(model.ERR_USER_LOCKED, "用户被锁定，请联系管理员")
}

//unLockUser 解锁用户
func (m *LoginLogic) unLockUser(userName string) error {
	m.cache.SetLoginSuccess(userName)
	return m.db.UnLock(userName)
}

//UpdateUserLoginTime 记录用户成功登录时间
func (m *LoginLogic) UpdateUserLoginTime(userID int64) error {
	return m.db.UpdateUserLoginTime(userID)
}

//SendWxValidCode 发送微信验证码
func (m *LoginLogic) SendWxValidCode(userName, openID, ident string) error {
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
func (m *LoginLogic) GetFreshToken() (string, error) {
	conf := model.GetConf()
	url := fmt.Sprintf("%s/%s/wechat/token/get", conf.RefreshWxTokenHost, conf.WxAppID)
	client := components.Def.HTTP().GetRegularClient()
	body, statusCode, err := client.Get(url)
	if err != nil {
		return "", err
	}
	if statusCode != 200 {
		return "", errs.NewErrorf(statusCode, "获取token信息失败,HttpStatus:%d, body:%s", statusCode, body)
	}

	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return "", fmt.Errorf("字符串转json发生错误，err：%v", err)
	}

	if errcode, ok := data["errcode"]; ok && types.GetInt(errcode) != 0 {
		return "", errs.NewError(http.StatusNotExtended, fmt.Errorf("获取token失败: %s", types.GetString(data["errmsg"])))
	}

	return types.GetString(data["access_token"]), nil
}

//SendCode 调用微信接口发验证码
func (m *LoginLogic) sendCode(openID, accessToken, validCode, ident string) error {
	conf := model.GetConf()
	data := m.constructSendData(openID, validCode, ident, conf.LoginValidCodeTemplateID)
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s?access_token=%s", conf.WxSendTemplateMsgURL, accessToken)
	client := components.Def.HTTP().GetRegularClient()
	content, statusCode, err := client.Post(url, string(dataJSON))

	if err != nil {
		return err
	}
	if statusCode != 200 {
		return errs.NewErrorf(statusCode, "发送验证码信息失败,HttpStatus:%d, body:%s", statusCode, content)
	}

	sendResult := make(map[string]interface{})
	err = json.Unmarshal([]byte(content), &sendResult)
	if err != nil {
		return fmt.Errorf("字符串转json发生错误，err：%v", err)
	}

	if errcode, ok := sendResult["errcode"]; ok && errcode != 0 {
		return errs.NewError(http.StatusNotExtended, fmt.Errorf("发送验证码信息失败: %s", types.GetString(sendResult["errmsg"])))
	}

	return nil
}

//构造发送验证码的实体数据
func (m *LoginLogic) constructSendData(openID, validCode, ident, templateID string) model.TemplateMsg {
	sendTitle := "用户系统"
	if !strings.EqualFold(ident, "") {
		system, _ := m.sysDB.QuerySysInfoByIdent(ident)
		if system != nil {
			sendTitle = system.GetString("name")
		}
	}

	data := model.TemplateMsg{
		Touser:     openID,
		TemplateID: templateID,
		Data: &model.TemplateData{
			First: model.KeyWordData{
				Value: fmt.Sprintf("你好,登录到[%s]需要进行验证码验证,请无泄露", sendTitle),
				Color: "#173177",
			},
			Keyword1: model.KeyWordData{
				Value: validCode,
				Color: "#FF0000",
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

//generateWrongPwdErrorCode 生成密码错误次数的errorCode
//currentCount 当前错误次数,totalCount 总次数(6)
func (m *LoginLogic) generateWrongPwdErrorCode(currentCount int) int {
	switch currentCount {
	case 1:
		return model.ERR_USER_PWDWRONG_5
	case 2:
		return model.ERR_USER_PWDWRONG_4
	case 3:
		return model.ERR_USER_PWDWRONG_3
	case 4:
		return model.ERR_USER_PWDWRONG_2
	case 5:
		return model.ERR_USER_PWDWRONG_1
	}
	return model.ERR_USER_LOCKED
}
