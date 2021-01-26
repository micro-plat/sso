package login

import (
	"fmt"
	"strings"
	"time"

	"github.com/lib4dev/vcs"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/access/member"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/access/system"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/const/enum"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/model"
	"github.com/micro-plat/sso/sso/errorcode"
)

//LoginLogic 用户登录相关
type LoginLogic struct {
	cache ICacheMember
	db    IDBMember
	memdb member.IDBMember
	sysDB system.IDbSystem
}

//NewLoginLogic 创建登录对象
func NewLoginLogic() *LoginLogic {
	return &LoginLogic{
		cache: NewCacheMember(),
		db:    NewDBMember(),
		memdb: member.NewDBMember(),
		sysDB: system.NewDbSystem(),
	}
}

//CheckValidCode 验证微信验证码是否正确
func (m *LoginLogic) CheckValidCode(userName, ident, validCode string) error {
	conf := model.GetLoginConf()
	if !conf.RequireValidCode {
		return nil
	}
	if strings.EqualFold(validCode, "") {
		return errs.NewError(errorcode.ERR_USER_EMPTY_VALIDATECODE, "验证码不能为空")
	}
	userInfo, err := m.memdb.GetUserInfo(userName)
	if err != nil {
		return errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}
	userAccount := ""
	switch conf.ValidCodeType {
	case enum.ValidCodeTypeSMS:
		userAccount = userInfo.GetString("mobile")
	case enum.ValidCodeTypeWechat:
		userAccount = userInfo.GetString("wx_openid")
	default:
		return errs.NewError(errorcode.ERR_VALID_CODE_TYPE_ERROR, fmt.Errorf("无效的ValidCodeType:%s", conf.ValidCodeType))
	}
	return vcs.VerifySmsCode(ident, userAccount, validCode, hydra.G.GetPlatName())
}

//CheckUserIsLocked 检查用户是否被锁定
func (m *LoginLogic) CheckUserIsLocked(userName string) error {
	conf := model.GetLoginConf()
	failCount := conf.UserLoginFailLimit
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
		return errs.NewError(errorcode.ERR_USER_LOCKED, "用户被锁定,请联系管理员")
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

	if strings.EqualFold(strings.ToLower(expassword), strings.ToLower(data.Get(0).GetString("password"))) {
		m.cache.SetLoginSuccess(userInfo.GetString("user_name"))
		return m.db.ChangePwd(userID, newpassword)
	}

	conf := model.GetLoginConf()
	count, err := m.cache.SetLoginFail(userInfo.GetString("user_name"))
	if err != nil {
		return err
	}
	if count <= conf.UserLoginFailLimit {
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
	return errs.NewError(errorcode.ERR_USER_LOCKED, "用户被锁定，请联系管理员")
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
		return errs.NewError(errorcode.ERR_SYS_LOCKED, "系统被禁用,不能登录")
	}
	return nil
}

//CheckUserInfo 检查用户
func (m *LoginLogic) checkUserInfo(userName, password string, state *model.MemberState) (err error) {
	if state.Status == enum.UserDisable {
		return errs.NewError(errorcode.ERR_USER_FORBIDDEN, "用户被禁用，请联系管理员")
	}
	if state.Status == enum.UserLock {
		return errs.NewError(errorcode.ERR_USER_LOCKED, "用户被锁定，请联系管理员")
	}

	if strings.ToLower(state.Password) == strings.ToLower(password) {
		m.cache.SetLoginSuccess(userName)
		return nil
	}

	count, err := m.cache.SetLoginFail(userName)
	if err != nil {
		return err
	}
	conf := model.GetLoginConf()
	if count <= conf.UserLoginFailLimit {
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

	return errs.NewError(errorcode.ERR_USER_LOCKED, "用户被锁定，请联系管理员")
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
		return errorcode.ERR_USER_PWDWRONG_5
	case 2:
		return errorcode.ERR_USER_PWDWRONG_4
	case 3:
		return errorcode.ERR_USER_PWDWRONG_3
	case 4:
		return errorcode.ERR_USER_PWDWRONG_2
	case 5:
		return errorcode.ERR_USER_PWDWRONG_1
	}
	return errorcode.ERR_USER_LOCKED
}

//Login 用户名密码登录
func (m *LoginLogic) SLogin(req model.LoginReq) (*model.LoginState, error) {

	ident := req.Ident
	if err := m.CheckSystemStatus(ident); err != nil {
		err = errs.NewError(errorcode.ERR_SYS_LOCKED, fmt.Errorf("判断系统是否被禁用:%v", err))
		return nil, err
	}

	userName := req.UserName
	if err := m.CheckUserIsLocked(userName); err != nil {
		err = errs.NewError(errorcode.ERR_USER_LOCKED, fmt.Errorf("判断用户是否被锁定, 锁定时间过期后要解锁:%v", err))
		return nil, err
	}

	if err := m.CheckValidCode(userName, req.Ident, req.ValidCode); err != nil {
		err = errs.NewError(errorcode.ERR_VALIDATECODE_WRONG, fmt.Errorf("判断用户输入的验证码:%v", err))
		return nil, err
	}

	member, err := m.Login(userName, req.Password, ident)
	if err != nil {
		return nil, err
	}

	m.UpdateUserLoginTime(member.UserID)

	return member, nil
}
