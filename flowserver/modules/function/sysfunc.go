package function

import (
	"github.com/micro-plat/hydra/component"
)

type ISystemFunc interface {
	Get(sysid int) (result []map[string]interface{}, err error)
	ChangeStatus(id int, status int) (err error)
	Delete(id int) (err error)
	Edit(input *SystemFuncEditInput) (err error)
	Add(input *SystemFuncAddInput) (err error)
}

type SystemFunc struct {
	c     component.IContainer
	cache ICacheSystemFunc
	db    IDbSystemFunc
}

func NewSystemFunc(c component.IContainer) *SystemFunc {
	return &SystemFunc{
		c:     c,
		cache: NewCacheSystemFunc(c),
		db:    NewDbSystemFunc(c),
	}
}

//Query 获取用系统管理列表
func (u *SystemFunc) Get(sysid int) (data []map[string]interface{}, err error) {
	//从缓存中获取功能信息，不存在时从数据库中获取
	data, err = u.cache.Query(sysid)
	if data == nil || err != nil {
		data, err = u.db.Get(sysid)
		if err != nil {
			return nil, err
		}
		//保存用户数据到缓存
		if err = u.cache.Save(sysid, data); err != nil {
			return nil, err
		}
	}
	return data, err
}

//ChangeStatus 修改功能状态
func (u *SystemFunc) ChangeStatus(id int, status int) (err error) {
	if err = u.db.ChangeStatus(id, status); err != nil {
		return
	}
	return u.cache.Fresh()
}

//Delete 删除系统功能
func (u *SystemFunc) Delete(id int) (err error) {
	if err = u.db.Delete(id); err != nil {
		return
	}
	return u.cache.Fresh()

}

//Edit 编辑功能
func (u *SystemFunc) Edit(input *SystemFuncEditInput) (err error) {
	if err = u.db.Edit(input); err != nil {
		return
	}
	return u.cache.Fresh()
}

//Add 添加功能
func (u *SystemFunc) Add(input *SystemFuncAddInput) (err error) {
	if err = u.db.Add(input); err != nil {
		return
	}
	return u.cache.Fresh()
}
