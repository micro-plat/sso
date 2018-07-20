package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type ISystem interface {
	Get(ident string) (s db.QueryRow, err error)
	Query(page int,name string,status string) (data db.QueryRows, count interface{}, err error)
	Delete(id int) (err error)
	Add(input *AddSystemInput) (err error)
	ChangeStatus(sysId int,status int) (err error)
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

//Query 从数据库中获取系统信息
func (m *System) Get(ident string) (s db.QueryRow, err error) {
	//从缓存中获取用户信息，不存在时从数据库中获取
	s, err = m.cache.Query(ident)
	if s == nil || err != nil {
		if s, err = m.db.Get(ident); err != nil {
			return nil, err
		}
		//保存用户数据到缓存
		if err = m.cache.Save(s); err != nil {
			return nil, err
		}
	}
	return s, err
}

//Query 获取用系统管理列表
func (u *System) Query(page int,name string,status string) (data db.QueryRows, count interface{}, err error) {
	//从缓存获取数据
	// data, err = u.cache.QuerySysInfo(page,name,status)
	// if data == nil || err != nil {
		data, count, err = u.db.Query(page,name,status)
		if  err != nil {
			return nil, nil, err
		}
		//保存系统数据到缓存
	// 	if err = u.cache.SaveSysInfo(page,name,status,data); err != nil {
	// 		return nil, err,nil
	// 	}
	// }
	return data, count, nil
}

//Delete 删除系统
func (u *System) Delete(id int) (err error){
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
func(u *System) ChangeStatus(sysId int, status int) (err error){
	if err = u.db.ChangeStatus(sysId, status); err != nil {
		return
	}
	//更新缓存
	return u.cache.FreshSysInfo()
	 
}
//Edit 编辑系统
func (u *System) Edit(input *SystemEditInput) (err error){
	if err = u.db.Edit(input); err != nil {
		return
	}
	//更新缓存
	return u.cache.FreshSysInfo()
	 
}
