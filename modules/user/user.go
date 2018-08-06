package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type IUser interface {
	Query(input *QueryUserInput) (data db.QueryRows, count int, err error)
	ChangeStatus(userID int, status int) (err error)
	Delete(userID int) (err error)
	Get(userID int) (data db.QueryRow, err error)
	Save(input *UserEditInput) (err error)
	Edit(username string,tel string,email string)(err error)
	ChangePwd(user_id int,expassword string,newpassword string)(err error)
	Bind(email string,openID string) (err error)
	SetEmail(Guid string,email string) (err error)
	GetEmail(Guid string) (email string,err error)
}

type User struct {
	c     component.IContainer
	db    IDbUser
	cache ICacheUser
}

func NewUser(c component.IContainer) *User {
	return &User{
		c:     c,
		db:    NewDbUser(c),
		cache: NewCacheUser(c),
	}
}

func(u *User) SetEmail(Guid string,email string) (err error) {
	return u.cache.SetEmail(Guid,email)
}

func(u *User) GetEmail(Guid string) (email string,err error) {
	return u.cache.GetEmail(Guid)
}

//Query 获取用户信息列表
func (u *User) Query(input *QueryUserInput) (data db.QueryRows, count int, err error) {
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
func (u *User) ChangeStatus(userID int, status int) (err error) {
	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.ChangeStatus(userID, status)
}

//Delete 删除用户
func (u *User) Delete(userID int) (err error) {
	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.Delete(userID)
}

//Get 查询用户信息
func (u *User) Get(userID int) (data db.QueryRow, err error) {
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

//Save 保存用户信息
func (u *User) Save(input *UserEditInput) (err error) {
	if err := u.cache.Delete(); err != nil {
		return err
	}
	if input.IsAdd == 1 {
		return u.db.Add(input)
	}
	return u.db.Edit(input)
}

func (u *User) Edit(username string,tel string,email string)(err error){
	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.EditInfo(username,tel,email)
}

func (u *User) ChangePwd(user_id int,expassword string,newpassword string)(err error) {
	if err := u.cache.Delete(); err != nil {
		return err
	}
	return u.db.ChangePwd(user_id,expassword,newpassword)
}

func (u *User) Bind(email string,openID string) (err error) {
	return u.db.Bind(email,openID)
}
