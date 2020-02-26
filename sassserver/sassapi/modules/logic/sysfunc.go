package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/access/function"
)

type ISystemFuncLogic interface {
	Get(sysid int) (result []map[string]interface{}, err error)
	ChangeStatus(id int, status int) (err error)
	Delete(id int) (err error)
}

type SystemFuncLogic struct {
	c  component.IContainer
	db function.IDbSystemFunc
}

func NewSystemFuncLogic(c component.IContainer) *SystemFuncLogic {
	return &SystemFuncLogic{
		c:  c,
		db: function.NewDbSystemFunc(c),
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
