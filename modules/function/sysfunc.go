package function

import (
	"github.com/micro-plat/hydra/component"

)

type ISystemFunc interface {
	Query(sysid int) (result []map[string]interface{}, err error)
	Enable(id int,status int) (err error)
	Delete(id int) (err error)
	Edit(input map[string]interface{}) (err error)
	Add(input map[string]interface{}) (err error)
}

type SystemFunc struct {
	c  component.IContainer
	db IDbSystemFunc
}

func NewSystemFunc(c component.IContainer) *SystemFunc {
	return &SystemFunc{
		c:  c,
		db: NewDbSystemFunc(c),
	}
}

//Query 获取用系统管理列表
func (u *SystemFunc) Query(sysid int) (data []map[string]interface{}, err error) {
	data, err = u.db.Query(sysid)
	if err != nil {
		return nil,  err
	}
	return
}

func(u *SystemFunc) Enable(id int,status int) (err error){
	err = u.db.Enable(id,status)
	if err != nil {
		return err
	}
	return nil
}
//删除系统
func (u *SystemFunc) Delete(id int) (err error){
	err = u.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *SystemFunc) Edit(input map[string]interface{}) (err error){
	err = u.db.Edit(input)
	if err != nil {
		return err
	}
	return nil
}

func (u *SystemFunc) Add(input map[string]interface{}) (err error){
	err = u.db.Add(input)
	if err != nil {
		return err
	}
	return nil
}