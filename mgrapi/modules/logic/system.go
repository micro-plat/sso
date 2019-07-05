package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/mgrapi/modules/access/system"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

type ISystemLogic interface {
	Get(ident string) (s db.QueryRow, err error)
	GetAll(userId int64) (s db.QueryRows, err error)
	Query(name string, status string, pi int, ps int) (data db.QueryRows, count int, err error)
	Delete(id int) (err error)
	Add(input *model.AddSystemInput) (err error)
	ChangeStatus(sysId int, status int) (err error)
	Edit(input *model.SystemEditInput) (err error)
	GetUsers(systemName string) (user db.QueryRows, allUser db.QueryRows, err error)
}

type SystemLogic struct {
	c     component.IContainer
	cache system.ICacheSystem
	db    system.IDbSystem
}

func NewSystemLogic(c component.IContainer) *SystemLogic {
	return &SystemLogic{
		c:     c,
		cache: system.NewCacheSystem(c),
		db:    system.NewDbSystem(c),
	}
}

//Get 从数据库中获取系统信息
func (u *SystemLogic) Get(ident string) (s db.QueryRow, err error) {
	//从缓存中获取用户信息，不存在时从数据库中获取
	s, err = u.cache.Query(ident)
	if s == nil || err != nil {
		if s, err = u.db.Get(ident); err != nil {
			return nil, err
		}
		//保存用户数据到缓存
		if err = u.cache.Save(s); err != nil {
			return nil, err
		}
	}
	return s, err
}

func (u *SystemLogic) GetAll(userId int64) (s db.QueryRows, err error) {
	return u.db.GetAll(userId)
}

//Query 获取用系统管理列表
func (u *SystemLogic) Query(name string, status string, pi int, ps int) (data db.QueryRows, count int, err error) {
	//从缓存获取数据
	data, count, err = u.cache.QuerySysInfo(name, status, pi, ps)
	if data == nil || err != nil || count == 0 {
		data, count, err = u.db.Query(name, status, pi, ps)
		if err != nil {
			return nil, 0, err
		}
		//保存系统数据到缓存
		if err = u.cache.SaveSysInfo(name, status, pi, ps, data, count); err != nil {
			return nil, 0, err
		}
	}
	return data, count, nil
}

//Delete 删除系统
func (u *SystemLogic) Delete(id int) (err error) {
	if err = u.db.Delete(id); err != nil {
		return
	}
	//更新缓存
	return u.cache.FreshSysInfo()
}

//Add 添加系统
func (u *SystemLogic) Add(input *model.AddSystemInput) (err error) {
	if err = u.db.Add(input); err != nil {
		return
	}
	//更新缓存
	return u.cache.FreshSysInfo()

}

//ChangeStatus 修改系统状态
func (u *SystemLogic) ChangeStatus(sysID int, status int) (err error) {
	if err = u.db.ChangeStatus(sysID, status); err != nil {
		return
	}
	//更新缓存
	return u.cache.FreshSysInfo()
}

//Edit 编辑系统
func (u *SystemLogic) Edit(input *model.SystemEditInput) (err error) {
	if err = u.db.Edit(input); err != nil {
		return
	}
	//更新缓存
	return u.cache.FreshSysInfo()

}

//GetUsers 获取系统下所有用户
func (u *SystemLogic) GetUsers(systemName string) (user db.QueryRows, allUser db.QueryRows, err error) {
	return u.db.GetUsers(systemName)
}
