package logic

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/jmz331/gpinyin"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/access/user"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/const/enum"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model"
)

type IUserLogic interface {
	Query(input *model.QueryUserInput) (data db.QueryRows, count int, err error)
	ChangeStatus(userID int, status int, mobile string, belongID, belongType int) (err error)
	Delete(userID, belongID, belongType int) (err error)
	Get(userID int) (data db.QueryRow, err error)
	GetAll(sysID, pi, ps int) (data db.QueryRows, count int, err error)
	Save(input *model.UserInputNew) (err error)
	Add(input *model.UserInputNew) (err error)
	SetDefaultPwd(userID, belongID, belongType int) error
	GenerateQrcodeInfo(userID int) (map[string]interface{}, error)

	GenerateUserNameByFullName(fullName string) string
}

type UserLogic struct {
	c     component.IContainer
	db    user.IDbUser
	cache user.ICacheUser
}

func NewUserLogic(c component.IContainer) *UserLogic {
	return &UserLogic{
		c:     c,
		db:    user.NewDbUser(c),
		cache: user.NewCacheUser(c),
	}
}

//Query 获取用户信息列表
func (u *UserLogic) Query(input *model.QueryUserInput) (data db.QueryRows, count int, err error) {
	if data, count, err = u.db.Query(input); err != nil {
		return nil, 0, err
	}
	return data, count, nil
}

//ChangeStatus 修改用户状态
func (u *UserLogic) ChangeStatus(userID int, status int, mobile string, belongID, belongType int) (err error) {
	if status == enum.UserNormal {
		u.cache.DeleteLockUserInfo(mobile)
	}
	return u.db.ChangeStatus(userID, status, belongID, belongType)
}

//Delete 删除用户
func (u *UserLogic) Delete(userID, belongID, belongType int) (err error) {
	return u.db.Delete(userID, belongID, belongType)
}

//Get 查询用户信息
func (u *UserLogic) Get(userID int) (data db.QueryRow, err error) {
	if data, err = u.db.Get(userID); err != nil {
		return nil, err
	}

	return data, nil
}

//GetAll GetAll
func (u *UserLogic) GetAll(sysID, pi, ps int) (data db.QueryRows, count int, err error) {
	if data, count, err = u.db.GetAll(sysID, pi, ps); err != nil {
		return nil, 0, err
	}
	return data, count, nil
}

//Save 保存要编辑的用户信息
func (u *UserLogic) Save(input *model.UserInputNew) (err error) {
	info2, err := u.db.GetUserInfoByMobile(input.Mobile)
	if err != nil {
		return
	}
	if info2 != nil && info2.GetInt64("user_id") != input.UserID {
		return context.NewError(model.ERR_USER_MOBILEEXISTS, "此姓名已被使用")
	}
	return u.db.Edit(input)
}

//Add 新增用户
func (u *UserLogic) Add(input *model.UserInputNew) (err error) {
	info2, err := u.db.GetUserInfoByMobile(input.Mobile)
	if err != nil {
		return nil
	}
	if info2 != nil {
		return context.NewError(model.ERR_USER_MOBILEEXISTS, "此电话号码已被使用")
	}

	return u.db.Add(input)
}

//SetDefaultPwd  重置密码
func (u *UserLogic) SetDefaultPwd(userID, belongID, belongType int) error {
	return u.db.ResetPwd(userID, belongID, belongType)
}

//GenerateQrcodeInfo 生成用户绑定二维码的信息
func (u *UserLogic) GenerateQrcodeInfo(userID int) (map[string]interface{}, error) {
	//1 验证用户s
	data, err := u.db.Get(userID)
	if err != nil {
		return nil, err
	}
	status := data.GetInt("status")
	if status == enum.UserLock {
		return nil, context.NewError(model.ERR_USER_LOCKED, "用户被锁定")
	}
	if status == enum.UserDisable {
		return nil, context.NewError(model.ERR_USER_FORBIDDEN, "用户被禁用")
	}

	//生成二维码数据
	timestamp := types.GetString(time.Now().Unix())
	values := net.NewValues()
	values.Set("user_id", string(userID))
	values.Set("timestamp", timestamp)

	values = values.Sort()
	raw := values.Join("", "") + model.WxBindSecrect

	return map[string]interface{}{
		"user_id":   userID,
		"timestamp": timestamp,
		"sign":      md5.Encrypt(raw),
	}, nil
}

//GenerateUserNameByFullName 根据名字生成登录名
func (u *UserLogic) GenerateUserNameByFullName(fullName string) string {
	numberSufix := ""
	arrName := strings.Split(gpinyin.ConvertToPinyinString(fullName, "-", gpinyin.PINYIN_WITHOUT_TONE), "-")
	matched, _ := regexp.MatchString(`.*\d{1}$`, fullName)
	if matched {
		numberSufix = arrName[len(arrName)-1]
		arrName = arrName[:len(arrName)-1]
	}

	result := fmt.Sprintf("%s%s", string(arrName[0]), string(arrName[1]))
	if len(arrName) > 2 {
		result = string(arrName[0])
		for i := 1; i < len(arrName); i++ {
			result = fmt.Sprintf("%s%s", result, string(string(arrName[i])[0]))
		}
	}
	return fmt.Sprintf("%s%s", result, numberSufix)
}
