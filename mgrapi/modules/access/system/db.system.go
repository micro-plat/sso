package system

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrapi/modules/const/sqls"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

type IDbSystem interface {
	Get(ident string) (s db.QueryRow, err error)
	GetAll(userId int64) (s db.QueryRows, err error)
	Query(name string, status string, pi int, ps int) (data db.QueryRows, count int, err error)
	Delete(id int) (err error)
	Add(input *model.AddSystemInput) (err error)
	ChangeStatus(sysID int, status int) (err error)
	Edit(input *model.SystemEditInput) (err error)
	Sort(sysID, sortRank, levelID, id, parentId int, isUp bool) (err error)
	GetUsers(systemName string) (user db.QueryRows, allUser db.QueryRows, err error)
}

type DbSystem struct {
	c component.IContainer
}

func NewDbSystem(c component.IContainer) *DbSystem {
	return &DbSystem{
		c: c,
	}
}

//Get 从数据库中获取系统信息
func (l *DbSystem) Get(ident string) (s db.QueryRow, err error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sqls.QuerySystemInfo, map[string]interface{}{
		"ident": ident,
	})
	return data.Get(0), err
}

func (l *DbSystem) GetAll(userId int64) (s db.QueryRows, err error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sqls.QueryAllSystemInfo, map[string]interface{}{
		"user_id": userId,
	})
	return data, err

}

//Query 获取用系统列表
func (u *DbSystem) Query(name string, status string, pi int, ps int) (data db.QueryRows, count int, err error) {
	db := u.c.GetRegularDB()
	c, q, a, err := db.Scalar(sqls.QuerySubSystemTotalCount, map[string]interface{}{
		"name":   " and name like '%" + name + "%'",
		"enable": status,
	})

	if err != nil {
		return nil, 0, fmt.Errorf("获取系统管理列表条数发生错误(err:%v),sql:(%s),输入参数:%v,", err, q, a)
	}
	data, q, a, err = db.Query(sqls.QuerySubSystemPageList, map[string]interface{}{
		"name":   " and t.name like '%" + name + "%'",
		"enable": status,
		"start":  (pi - 1) * ps,
		"ps":     ps,
	})

	if err != nil {
		return nil, 0, fmt.Errorf("获取系统管理列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c), nil
}

func (u *DbSystem) Delete(id int) (err error) {
	db := u.c.GetRegularDB()
	_, q, a, err := db.Execute(sqls.DeleteSubSystemById, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return fmt.Errorf("删除系统管理列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (u *DbSystem) Add(input *model.AddSystemInput) (err error) {
	if input.Wechat_status == "" {
		input.Wechat_status = "0"
	}

	db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"name":          input.Name,
		"addr":          input.Addr,
		"time_out":      input.Time_out,
		"logo":          input.Logo,
		"style":         input.Style,
		"theme":         input.Theme,
		"ident":         input.Ident,
		"wechat_status": input.Wechat_status,
		"login_url":     "http://member/login",
		"secret":        input.Secret,
	}
	_, q, a, err := db.Execute(sqls.AddSubSystem, params)
	if err != nil {
		return fmt.Errorf("添加系统管理数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (u *DbSystem) ChangeStatus(sysId int, status int) (err error) {
	db := u.c.GetRegularDB()
	_, q, a, err := db.Execute(sqls.UpdateEnable, map[string]interface{}{
		"id":     sysId,
		"enable": status,
	})
	if err != nil {
		return fmt.Errorf("更新系统管理状态发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (u *DbSystem) Edit(input *model.SystemEditInput) (err error) {
	db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"enable":        input.Enable,
		"id":            input.Id,
		"index_url":     input.Index_url,
		"login_timeout": input.Login_timeout,
		"logo":          input.Logo,
		"name":          input.Name,
		"layout":        input.Layout,
		"theme":         input.Theme,
		"ident":         input.Ident,
		"wechat_status": input.Wechat_status,
		"secret":        input.Secret,
	}
	_, q, a, err := db.Execute(sqls.UpdateEdit, params)
	if err != nil {
		return fmt.Errorf("更新系统管理数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

// Sort 功能菜单上下调整
func (u *DbSystem) Sort(sysID, sortRank, levelID, id, parentId int, isUp bool) (err error) {
	params := map[string]interface{}{
		"sys_id":   sysID,
		"level_id": levelID,
		"parent":   parentId,
	}

	params["sortrank"] = fmt.Sprintf(" and t.sortrank > %d ", sortRank)
	params["orderby"] = " order by t.sortrank asc "
	if isUp {
		params["sortrank"] = fmt.Sprintf(" and t.sortrank < %d", sortRank)
		params["orderby"] = " order by t.sortrank desc "
	}

	db := u.c.GetRegularDB()
	data, q, a, err := db.Query(sqls.QuerySsoSystemMenu, params)

	if err != nil {
		return fmt.Errorf("查询系统列表错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	if len(data) <= 0 {
		return fmt.Errorf("没有可以调换的菜单")
	}
	changeRow := data[0]

	db2 := u.c.GetRegularDB()
	trans, err := db2.Begin()
	if err != nil {
		return fmt.Errorf("调换的菜单时创建事务失败: %s", err.Error())
	}

	_, q, a, err = trans.Execute(sqls.UpSsoSystemMenu,
		map[string]interface{}{
			"sys_id":   sysID,
			"level_id": levelID,
			"id":       id,
			"sortrank": changeRow.GetString("sortrank"),
		})

	if err != nil {
		trans.Rollback()
		return fmt.Errorf("更新系统管理排序发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	_, q, a, err = trans.Execute(sqls.UpSsoSystemMenu,
		map[string]interface{}{
			"sys_id":   changeRow.GetString("sys_id"),
			"level_id": changeRow.GetString("level_id"),
			"id":       changeRow.GetString("id"),
			"sortrank": sortRank,
		})

	if err != nil {
		trans.Rollback()
		return fmt.Errorf("更新系统管理排序发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	trans.Commit()
	return nil
}

//GetUsers 获取系统下所有用户
func (u *DbSystem) GetUsers(systemName string) (user db.QueryRows, allUser db.QueryRows, err error) {

	db := u.c.GetRegularDB()
	data, q, a, err := db.Query(sqls.GetUsers, map[string]interface{}{
		"system_name": systemName,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("获取系统下所有用户发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	datas, q, a, err := db.Query(sqls.GetAllUser, map[string]interface{}{})
	if err != nil {
		return nil, nil, fmt.Errorf("获取所有用户发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, datas, nil

}
