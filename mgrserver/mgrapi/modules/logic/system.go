package logic

import (
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/access/system"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

type ISystemLogic interface {
	Get(ident string) (s db.QueryRow, err error)
	GetAll(userId int64) (s db.QueryRows, err error)
	Query(name string, status string, pi int, ps int) (data db.QueryRows, count int, err error)
	Delete(id int) (err error)
	Add(input *model.AddSystemInput) (err error)
	ChangeStatus(sysId int, status int) (err error)
	Edit(input *model.SystemEditInput) (err error)
	Sort(sysID, sortrank, levelID, id, parentId int, isUp bool) (err error)
	GetUsers(systemName string) (user db.QueryRows, allUser db.QueryRows, err error)
	ChangeSecret(id int, secret string) error
}

type SystemLogic struct {
	db system.IDbSystem
}

func NewSystemLogic() *SystemLogic {
	return &SystemLogic{
		db: system.NewDbSystem(),
	}
}

//Get 从数据库中获取系统信息
func (u *SystemLogic) Get(ident string) (s db.QueryRow, err error) {
	if s, err = u.db.Get(ident); err != nil {
		return nil, err
	}
	return s, nil
}

func (u *SystemLogic) GetAll(userId int64) (s db.QueryRows, err error) {
	return u.db.GetAll(userId)
}

//Query 获取用系统管理列表
func (u *SystemLogic) Query(name string, status string, pi int, ps int) (data db.QueryRows, count int, err error) {
	data, count, err = u.db.Query(name, status, pi, ps)
	if err != nil {
		return nil, 0, err
	}
	return data, count, err
}

//Delete 删除系统
func (u *SystemLogic) Delete(id int) (err error) {
	if err = u.db.Delete(id); err != nil {
		return
	}
	return nil
}

//Add 添加系统
func (u *SystemLogic) Add(input *model.AddSystemInput) (err error) {
	//1验证系统名称,ident是否重复

	count, err := u.db.ExistsNameOrIdent(input.Name, input.Ident)
	if err != nil {
		return err
	}
	if count > 0 {
		return errs.NewError(model.ERR_SYS_NAMEORIDENTEXISTS, "系统名称和英文名称已存在")
	}

	if err = u.db.Add(input); err != nil {
		return
	}
	return nil
}

//ChangeStatus 修改系统状态
func (u *SystemLogic) ChangeStatus(sysID int, status int) (err error) {
	if err = u.db.ChangeStatus(sysID, status); err != nil {
		return
	}
	return nil
}

//Edit 编辑系统
func (u *SystemLogic) Edit(input *model.SystemEditInput) (err error) {
	if err = u.db.Edit(input); err != nil {
		return
	}
	return nil
}

// Sort 对菜单功能排序
func (u *SystemLogic) Sort(sysID, sortrank, levelID, id, parentId int, isUp bool) (err error) {
	if err = u.db.Sort(sysID, sortrank, levelID, id, parentId, isUp); err != nil {
		return
	}
	return nil
}

//GetUsers 获取系统下所有用户
func (u *SystemLogic) GetUsers(systemName string) (user db.QueryRows, allUser db.QueryRows, err error) {
	return u.db.GetUsers(systemName)
}

//ChangeSecret 修改secret
func (u *SystemLogic) ChangeSecret(id int, secret string) error {
	return u.db.ChangeSecret(id, secret)
}
