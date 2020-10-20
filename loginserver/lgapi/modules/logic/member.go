package logic

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/net/http"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/lib4go/utility"
	commodel "github.com/micro-plat/sso/common/module/model"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/access/member"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/access/system"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/const/enum"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/const/sqls"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/model"
)

//IMemberLogic 用户登录
type IMemberLogic interface {
	CheckSystemStatus(ident string) error
	CreateLoginUserCode(userID int64) (code string, err error)
	//CheckUserIsLocked(userName string) error
	//SendWxValidCode(userName, openID, ident string) error
	//ChangePwd(userID int, expassword string, newpassword string) (err error)
	CheckHasRoles(userID int64, ident string) error
	GenerateCodeAndSysInfo(ident string, userID int64) (map[string]string, error)
	CheckUerInfo(userID int64, sign, timestamp string) error
	GenerateWxStateCode(userID int64) (string, error)
	ValidStateAndGetOpenID(stateCode, wxCode string) (map[string]string, error)
	UpdateUserOpenID(data map[string]string) error
	ValidUserInfo(userName string) (string, error)
	QueryUserInfoByID(uid int64) (db.QueryRow, error)
	ValidUserAndGetUserInfo(userName string) (db.QueryRow, error)
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

// // ChangePwd 修改密码
// func (m *MemberLogic) ChangePwd(userID int, expassword string, newpassword string) (err error) {
// 	return m.db.ChangePwd(userID, expassword, newpassword)
// }

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
		return context.NewError(commodel.ERR_SYS_LOCKED, "系统被禁用,不能登录")
	}
	return nil
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
		return context.NewError(commodel.ERR_USER_LOCKED, "用户被锁定,暂时无法登录")
	}
	if status == enum.UserDisable {
		return context.NewError(commodel.ERR_USER_FORBIDDEN, "用户被禁用，暂时无法登录")
	}

	return m.db.CheckUserHasAuth(ident, userID)
}

//CheckUerInfo 验证要绑定的用户信息
func (m *MemberLogic) CheckUerInfo(userID int64, sign, timestamp string) error {
	values := net.NewValues()
	values.Set("user_id", string(userID))
	values.Set("timestamp", timestamp)
	values = values.Sort()
	raw := values.Join("", "") + model.WxBindSecrect
	if !strings.EqualFold(sign, md5.Encrypt(raw)) {
		return context.NewError(commodel.ERR_BIND_INFOWRONG, "绑定信息错误,请重新去用户系统扫码")
	}
	sendTime, _ := strconv.ParseInt(timestamp, 10, 64)
	if time.Now().Unix()-sendTime > int64(commodel.GetConf(m.c).BindTimeOut) {
		return context.NewError(commodel.ERR_QRCODE_TIMEOUT, "二维码过期,请联系管理员重新生成")
	}

	data, err := m.db.QueryByID(userID)
	if err != nil {
		return err
	}
	status := data.GetInt("status")
	if status == enum.UserLock {
		return context.NewError(commodel.ERR_USER_FORBIDDEN, "用户被禁用")
	}
	if status == enum.UserDisable {
		return context.NewError(commodel.ERR_USER_LOCKED, "用户被锁定")
	}
	if data.GetString("wx_openid") != "" {
		return context.NewError(commodel.ERR_USER_EXISTSWX, "用户已绑定微信")
	}

	return nil
}

//GenerateWxStateCode 生成微信stateCode凭证
func (m *MemberLogic) GenerateWxStateCode(userID int64) (string, error) {
	stateCode := utility.GetGUID()
	if err := m.cache.SaveWxStateCode(stateCode, strconv.FormatInt(userID, 10)); err != nil {
		return "", err
	}
	return stateCode, nil
}

//ValidStateAndGetOpenID 验证state信息并获取openid
func (m *MemberLogic) ValidStateAndGetOpenID(stateCode, wxCode string) (map[string]string, error) {
	userID, err := m.cache.GetWxStateCodeUserId(stateCode)
	if err != nil {
		return nil, err
	}
	if userID == "" {
		return nil, context.NewError(model.ERR_BIND_TIMEOUT, "绑定超时")
	}

	config := commodel.GetConf(m.c)
	url := config.WxTokenURL + "?appid=" + config.WxAppID + "&secret=" + config.WxSecret + "&code=" + wxCode + "&grant_type=authorization_code"
	openID, err := m.GetWxUserOpID(url)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"openid": openID,
		"userid": userID,
	}, nil
}

//UpdateUserOpenID 保存用户的openid
func (m *MemberLogic) UpdateUserOpenID(data map[string]string) error {
	return m.db.UpdateUserOpenID(data)
}

//QueryUserInfoByID 通过user_id获取用户信息
func (m *MemberLogic) QueryUserInfoByID(uid int64) (db.QueryRow, error) {
	return m.db.QueryByID(uid)
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
		return "", err
	}
	if datas.IsEmpty() {
		return "", context.NewError(model.ERR_USER_NOTEXISTS, "用户不存在")
	}
	if datas.Get(0).GetString("wx_openid") == "" {
		return "", context.NewError(model.ERR_USER_NOTBINDWX, "用户还未绑定微信账户")
	}
	return datas.Get(0).GetString("wx_openid"), nil
}

//ValidUserAndGetUserInfo 验证用户是否存在并获取用户信息
func (m *MemberLogic) ValidUserAndGetUserInfo(userName string) (db.QueryRow, error) {
	db := m.c.GetRegularDB()

	count, q, a, err := db.Scalar(sqls.ValidUserNameExist, map[string]interface{}{
		"user_name": userName,
	})
	if err != nil {
		return nil, fmt.Errorf("验证用户是否存在失败(err:%v),sql:%s,参数:%v", err, q, a)
	}

	if types.GetInt(count) <= 0 {
		return nil, fmt.Errorf("用户不存在:%s", userName)
	}

	userInfo, q, a, err := db.Query(sqls.QueryUserByUserName, map[string]interface{}{
		"user_name": userName,
	})

	if err != nil {
		return nil, fmt.Errorf("根据用户名获取用户信息失败(err:%v),sql:%s,参数:%v", err, q, a)
	}

	if len(userInfo.Get(0).GetString("mobile")) == 0 {
		return nil, fmt.Errorf("用户没有绑定手机号")
	}

	return userInfo.Get(0), nil
}
