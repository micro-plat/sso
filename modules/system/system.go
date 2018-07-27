package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type ISystem interface {
	Get(ident string) (s db.QueryRow, err error)
	Query(name string, status string, pi int, ps int) (data db.QueryRows, count int, err error)
	Delete(id int) (err error)
	Add(input *AddSystemInput) (err error)
	ChangeStatus(sysId int, status int) (err error)
	Edit(input *SystemEditInput) (err error)
}

type System struct {
	c     component.IContainer
	cache ICacheSystem
	db    IDbSystem
}

func NewSystem(c component.IContainer) *System {
	return &System{
		c:     c,
		cache: NewCacheSystem(c),
		db:    NewDbSystem(c),
	}
}

//Get 从数据库中获取系统信息
func (u *System) Get(ident string) (s db.QueryRow, err error) {
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

//Query 获取用系统管理列表
func (u *System) Query(name string, status string, pi int, ps int) (data db.QueryRows, count int, err error) {
	//从缓存获取数据
	data,count, err = u.cache.QuerySysInfo(name, status, pi, ps)
	if data == nil || err != nil {
		data, count, err = u.db.Query(name, status, pi, ps)
		if err != nil {
			return nil, 0, err
		}
		//保存系统数据到缓存
		if err = u.cache.SaveSysInfo(name, status, pi, ps, data,count); err != nil {
			return nil, 0, err
		}
	}
	return data, count, nil
}

//Delete 删除系统
func (u *System) Delete(id int) (err error) {
	if err = u.db.Delete(id); err != nil {
		return
	}
	//更新缓存
	return u.cache.FreshSysInfo()
}

//Add 添加系统
func (u *System) Add(input *AddSystemInput) (err error) {
	if err = u.db.Add(input); err != nil {
		return
	}
	//更新缓存
	return u.cache.FreshSysInfo()

}

//ChangeStatus 修改系统状态
func (u *System) ChangeStatus(sysID int, status int) (err error) {
	if err = u.db.ChangeStatus(sysID, status); err != nil {
		return
	}
	//更新缓存
	return u.cache.FreshSysInfo()
}

//Edit 编辑系统
func (u *System) Edit(input *SystemEditInput) (err error) {
	if err = u.db.Edit(input); err != nil {
		return
	}
	//更新缓存
	return u.cache.FreshSysInfo()

}
