package system

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrapi/modules/const/sql"
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
	Up(sysID int, sortrank int, levelID int, id int) (err error)
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
	data, _, _, err := db.Query(sql.QuerySystemInfo, map[string]interface{}{
		"ident": ident,
	})
	return data.Get(0), err
}

func (l *DbSystem) GetAll(userId int64) (s db.QueryRows, err error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sql.QueryAllSystemInfo, map[string]interface{}{
		"user_id": userId,
	})
	return data, err

}

//Query 获取用系统列表
func (u *DbSystem) Query(name string, status string, pi int, ps int) (data db.QueryRows, count int, err error) {
	db := u.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QuerySubSystemTotalCount, map[string]interface{}{
		"name":   " and name like '%" + name + "%'",
		"enable": status,
	})

	if err != nil {
		return nil, 0, fmt.Errorf("获取系统管理列表条数发生错误(err:%v),sql:(%s),输入参数:%v,", err, q, a)
	}
	data, q, a, err = db.Query(sql.QuerySubSystemPageList, map[string]interface{}{
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
	_, q, a, err := db.Execute(sql.DeleteSubSystemById, map[string]interface{}{
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
	_, q, a, err := db.Execute(sql.AddSubSystem, params)
	if err != nil {
		return fmt.Errorf("添加系统管理数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (u *DbSystem) ChangeStatus(sysId int, status int) (err error) {
	db := u.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.UpdateEnable, map[string]interface{}{
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
	_, q, a, err := db.Execute(sql.UpdateEdit, params)
	if err != nil {
		return fmt.Errorf("更新系统管理数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

// Up 功能菜单上下调整
func (u *DbSystem) Up(sysID int, sortrank int, levelID int, id int) (err error) {
	db := u.c.GetRegularDB()

	data, q, a, err := db.Query(sql.QuerySsoSystemMenu, map[string]interface{}{
		"sys_id":   sysID,
		"level_id": levelID,
	})

	if err != nil {
		return fmt.Errorf("查询系统列表错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	var num int
	var i int
	for index, value := range data {
		if sortrank == types.GetInt(value["sortrank"]) && index >= 1 {
			num = data[index-1].GetInt("sortrank")
			i = data[index-1].GetInt("id")
		}
	}

	//以下没有用事务
	fmt.Printf("传入编号:%d, 前一个编号:%d, 序号:%d, 排序号:%d", id, num, i, sortrank)
	_, q, a, err = db.Execute(sql.UpSsoSystemMenu, map[string]interface{}{
		"sys_id":   sysID,
		"level_id": levelID,
		"id":       id,
		"num":      num,
	})

	if err != nil {
		return fmt.Errorf("更新系统管理排序发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	_, q, a, err = db.Execute(sql.UpSsoSystemMenuList, map[string]interface{}{
		"sys_id":   sysID,
		"level_id": levelID,
		"sortrank": sortrank,
		"i":        i,
	})

	if err != nil {
		return fmt.Errorf("更新系统管理排序发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return nil
}

//GetUsers 获取系统下所有用户
func (u *DbSystem) GetUsers(systemName string) (user db.QueryRows, allUser db.QueryRows, err error) {

	db := u.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetUsers, map[string]interface{}{
		"system_name": systemName,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("获取系统下所有用户发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	datas, q, a, err := db.Query(sql.GetAllUser, map[string]interface{}{})
	if err != nil {
		return nil, nil, fmt.Errorf("获取所有用户发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, datas, nil

}
