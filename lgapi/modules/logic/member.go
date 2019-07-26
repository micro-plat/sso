package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
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
	SaveWxLoginStateCode(code string) error
	ExistsWxLoginStateCode(code string) (bool, error)
	GetWxUserOpID(url string) (string, error)
	GetUserInfoByOpID(opID, ident string) (*model.LoginState, error)
}

//MemberLogic 用户登录管理
type MemberLogic struct {
	cache member.ICacheMember
	db    member.IDBMember
	http  *http.Client
}

//NewMemberLogic 创建登录对象
func NewMemberLogic(c component.IContainer) *MemberLogic {
	return &MemberLogic{
		cache: member.NewCacheMember(c),
		db:    member.NewDBMember(c),
		http:  &http.Client{Timeout: 5 * time.Second},
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

//SaveWxLoginStateCode xx
func (m *MemberLogic) SaveWxLoginStateCode(code string) error {
	return m.cache.SaveWxLoginStateCode(code)
}

//ExistsWxLoginStateCode 是否存在wx state code 防伪造
func (m *MemberLogic) ExistsWxLoginStateCode(code string) (bool, error) {
	return m.cache.ExistsWxLoginStateCode(code)
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

	return token["openid"].(string), nil
}

//GetUserInfoByOpID xxx
func (m *MemberLogic) GetUserInfoByOpID(opID, ident string) (*model.LoginState, error) {
	data, err := m.db.QueryByOpenID(opID, ident)
	if err != nil {
		return nil, err
	}
	return (*model.LoginState)(data), nil
}
