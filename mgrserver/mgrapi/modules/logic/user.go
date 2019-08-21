package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/access/user"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/const/enum"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

type IUserLogic interface {
	Query(input *model.QueryUserInput) (data db.QueryRows, count int, err error)
	ChangeStatus(userID int, status int, userName string) (err error)
	Delete(userID int) (err error)
	Get(userID int) (data db.QueryRow, err error)
	GetAll(sysID, pi, ps int) (data db.QueryRows, count int, err error)
	Save(input *model.UserInputNew) (err error)
	Add(input *model.UserInputNew) (err error)
	Edit(username string, tel string, email string) (err error)
	SetDefaultPwd(userID int) error
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

	// //从缓存中获取用户信息，不存在时从数据库中获取
	// data, count, err = u.cache.Query(input)
	// if data == nil || count == 0 || err != nil {
	// 	if data, count, err = u.db.Query(input); err != nil {
	// 		return nil, 0, err
	// 	}
	// 	if err = u.cache.Save(input, data, count); err != nil {
	// 		return nil, 0, err
	// 	}
	// }
	// return data, count, nil
}

//ChangeStatus 修改用户状态
func (u *UserLogic) ChangeStatus(userID int, status int, userName string) (err error) {
	if err := u.cache.Delete(); err != nil {
		return err
	}
	if status == enum.UserNormal {
		u.cache.DeleteLockUserInfo(userName)
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

//GetAll GetAll
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
	info, err := u.db.GetUserInfoByName(input.UserName)
	if err != nil {
		return err
	}
	if info != nil && info.GetInt64("user_id") != input.UserID {
		return context.NewError(context.ERR_BAD_REQUEST, "此用户名已被使用")
	}

	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.Edit(input)
}

//Add 新增用户
func (u *UserLogic) Add(input *model.UserInputNew) (err error) {
	info, err := u.db.GetUserInfoByName(input.UserName)
	if err != nil {
		return err
	}
	if info != nil {
		return context.NewError(context.ERR_BAD_REQUEST, "此用户名已被使用")
	}

	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.Add(input)
}

//Edit 修改用户信息
func (u *UserLogic) Edit(username string, tel string, email string) (err error) {
	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.EditInfo(username, tel, email)
}

//SetDefaultPwd  重置密码
func (u *UserLogic) SetDefaultPwd(userID int) error {
	return u.db.ResetPwd(userID)
}
