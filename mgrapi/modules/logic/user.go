package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/mgrapi/modules/access/user"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

type IUserLogic interface {
	Query(input *model.QueryUserInput) (data db.QueryRows, count int, err error)
	ChangeStatus(userID int, status int) (err error)
	Delete(userID int) (err error)
	Get(userID int) (data db.QueryRow, err error)
	GetAll(sysID, pi, ps int) (data db.QueryRows, count int, err error)
	Save(input *model.UserInputNew) (err error)
	Add(input *model.UserInputNew) (err error)
	Edit(username string, tel string, email string) (err error)
	ChangePwd(user_id int, expassword string, newpassword string) (err error) //对外api要调用
	ResetPwd(user_id int64) (err error)
	Bind(email string, openID string) (err error)
	SetEmail(Guid string, email string) (err error)
	GetEmail(Guid string) (email string, err error)
	IsSendEmail(input *model.UserInputNew) (b bool, err error)
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

func (u *UserLogic) SetEmail(Guid string, email string) (err error) {
	return u.cache.SetEmail(Guid, email)
}

func (u *UserLogic) GetEmail(Guid string) (email string, err error) {
	return u.cache.GetEmail(Guid)
}

func (u *UserLogic) IsSendEmail(input *model.UserInputNew) (b bool, err error) {

	return u.db.IsSendEmail(input)
}

//Query 获取用户信息列表
func (u *UserLogic) Query(input *model.QueryUserInput) (data db.QueryRows, count int, err error) {
	//从缓存中获取用户信息，不存在时从数据库中获取
	data, count, err = u.cache.Query(input)
	if data == nil || count == 0 || err != nil {
		if data, count, err = u.db.Query(input); err != nil {
			return nil, 0, err
		}
		if err = u.cache.Save(input, data, count); err != nil {
			return nil, 0, err
		}
	}
	return data, count, nil
}

//ChangeStatus 修改用户状态
func (u *UserLogic) ChangeStatus(userID int, status int) (err error) {
	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.ChangeStatus(userID, status)
}

//Delete 删除用户
func (u *UserLogic) Delete(userID int) (err error) {
	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.Delete(userID)
}

//Get 查询用户信息
func (u *UserLogic) Get(userID int) (data db.QueryRow, err error) {
	data, err = u.cache.QueryUser(userID)
	if data == nil || err != nil {
		if data, err = u.db.Get(userID); err != nil {
			return nil, err
		}
		if err = u.cache.SaveUser(userID, data); err != nil {
			return nil, err
		}
	}
	return data, nil
}

func (u *UserLogic) GetAll(sysID, pi, ps int) (data db.QueryRows, count int, err error) {
	data, count, err = u.cache.QueryUserBySys(sysID, pi, ps)
	if data == nil || err != nil {
		if data, count, err = u.db.GetAll(sysID, pi, ps); err != nil {
			return nil, 0, err
		}
		if err = u.cache.SaveUserBySys(sysID, pi, ps, data, count); err != nil {
			return nil, 0, err
		}
	}
	return data, count, nil
}

//Save 保存要编辑的用户信息
func (u *UserLogic) Save(input *model.UserInputNew) (err error) {
	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.Edit(input)
}

func (u *UserLogic) Add(input *model.UserInputNew) (err error) {
	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.Add(input)
}

func (u *UserLogic) Edit(username string, tel string, email string) (err error) {
	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.EditInfo(username, tel, email)
}

func (u *UserLogic) ChangePwd(user_id int, expassword string, newpassword string) (err error) {
	return u.db.ChangePwd(user_id, expassword, newpassword)
}

func (u *UserLogic) ResetPwd(user_id int64) (err error) {
	return u.db.ResetPwd(user_id)
}

func (u *UserLogic) Bind(email string, openID string) (err error) {
	return u.db.Bind(email, openID)
}