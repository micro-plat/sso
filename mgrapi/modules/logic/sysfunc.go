package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/mgrapi/modules/access/function"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

type ISystemFuncLogic interface {
	Get(sysid int) (result []map[string]interface{}, err error)
	ChangeStatus(id int, status int) (err error)
	Delete(id int) (err error)
	Edit(input *model.SystemFuncEditInput) (err error)
	Add(input *model.SystemFuncAddInput) (err error)
}

type SystemFuncLogic struct {
	c     component.IContainer
	cache function.ICacheSystemFunc
	db    function.IDbSystemFunc
}

func NewSystemFuncLogic(c component.IContainer) *SystemFuncLogic {
	return &SystemFuncLogic{
		c:     c,
		cache: function.NewCacheSystemFunc(c),
		db:    function.NewDbSystemFunc(c),
	}
}

//Query 获取用系统管理列表
func (u *SystemFuncLogic) Get(sysid int) (data []map[string]interface{}, err error) {
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
func (u *SystemFuncLogic) ChangeStatus(id int, status int) (err error) {
	if err = u.db.ChangeStatus(id, status); err != nil {
		return
	}
	return u.cache.Fresh()
}

//Delete 删除系统功能
func (u *SystemFuncLogic) Delete(id int) (err error) {
	if err = u.db.Delete(id); err != nil {
		return
	}
	return u.cache.Fresh()

}

//Edit 编辑功能
func (u *SystemFuncLogic) Edit(input *model.SystemFuncEditInput) (err error) {
	if err = u.db.Edit(input); err != nil {
		return
	}
	return u.cache.Fresh()
}

//Add 添加功能
func (u *SystemFuncLogic) Add(input *model.SystemFuncAddInput) (err error) {
	if err = u.db.Add(input); err != nil {
		return
	}
	return u.cache.Fresh()
}
