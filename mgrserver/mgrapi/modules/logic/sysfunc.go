package logic

import (
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/access/function"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

type ISystemFuncLogic interface {
	Get(sysid int) (result []map[string]interface{}, err error)
	ChangeStatus(id int, status int) (err error)
	Delete(id int) (err error)
	Edit(input *model.SystemFuncEditInput) (err error)
	Add(input *model.SystemFuncAddInput) (err error)
}

type SystemFuncLogic struct {
	cache function.ICacheSystemFunc
	db    function.IDbSystemFunc
}

func NewSystemFuncLogic() *SystemFuncLogic {
	return &SystemFuncLogic{
		cache: function.NewCacheSystemFunc(),
		db:    function.NewDbSystemFunc(),
	}
}

//Get 获取用系统管理列表
func (u *SystemFuncLogic) Get(sysid int) (data []map[string]interface{}, err error) {
	data, err = u.db.Get(sysid)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//ChangeStatus 修改功能状态
func (u *SystemFuncLogic) ChangeStatus(id int, status int) (err error) {
	if err = u.db.ChangeStatus(id, status); err != nil {
		return
	}
	return nil
}

//Delete 删除系统功能
func (u *SystemFuncLogic) Delete(id int) (err error) {
	if err = u.db.Delete(id); err != nil {
		return
	}
	return nil
}

//Edit 编辑功能
func (u *SystemFuncLogic) Edit(input *model.SystemFuncEditInput) (err error) {
	if err = u.db.Edit(input); err != nil {
		return
	}
	return nil
}

//Add 添加功能
func (u *SystemFuncLogic) Add(input *model.SystemFuncAddInput) (err error) {
	if err = u.db.Add(input); err != nil {
		return
	}
	return nil
}
