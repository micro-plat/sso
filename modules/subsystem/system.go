package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type ISystem interface {
	Query(page int) (data db.QueryRows, count interface{}, err error)
	QueryWithField(input map[string]interface{}) (data db.QueryRows, err error)
	DeleteByID(id int) (err error)
	Add(input map[string]interface{}) (err error)
	UpdateEnable(sysId int,status int) (err error)
	UpdateEdit(input map[string]interface{}) (err error)
}

type System struct {
	c  component.IContainer
	db IDbSystem
}

func NewSystem(c component.IContainer) *System {
	return &System{
		c:  c,
		db: NewDbSystem(c),
	}
}

//Query 获取用系统管理列表
func (u *System) Query(page int) (data db.QueryRows, count interface{}, err error) {
	data, count, err = u.db.Query(page)
	if err != nil {
		return nil, nil, err
	}
	return data, count, nil
}

func (u *System) QueryWithField(input map[string]interface{}) (data db.QueryRows, err error) {
	data, err = u.db.QueryWithField(input)
	if err != nil {
		return nil,  err
	}
	return data,nil
}

func (u *System) DeleteByID(id int) (err error){
	err = u.db.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *System) Add(input map[string]interface{}) (err error) {
	err = u.db.Add(input)
	if err != nil {
		return err
	}
	return nil
}

func(u *System) UpdateEnable(sysId int, status int) (err error){
	err = u.db.UpdateEnable(sysId, status)
	if err != nil {
		return err
	}
	return nil
}

func (u *System) UpdateEdit(input map[string]interface{}) (err error){
	err = u.db.UpdateEdit(input)
	if err != nil {
		return err
	}
	return nil
}