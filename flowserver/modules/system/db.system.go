package system

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/flowserver/modules/const/sql"
)

type IDbSystem interface {
	Get(ident string) (s db.QueryRow, err error)
	GetAll(userId int64) (s db.QueryRows, err error)
	Query(name string, status string, pi int, ps int) (data db.QueryRows, count int, err error)
	Delete(id int) (err error)
	Add(input *AddSystemInput) (err error)
	ChangeStatus(sysID int, status int) (err error)
	Edit(input *SystemEditInput) (err error)
	GetUsers(systemName string) (user db.QueryRows, allUser db.QueryRows, err error)
}

type SystemEditInput struct {
	Enable        string `form:"enable" json:"enable" valid:"required"`
	Id            string `form:"id" json:"id" valid:"required"`
	Index_url     string `form:"index_url" json:"index_url" valid:"required"`
	Login_timeout string `form:"login_timeout" json:"login_timeout" valid:"required"`
	Logo          string `form:"logo" json:"logo" valid:"required"`
	Name          string `form:"name" json:"name" valid:"required"`
	Theme         string `form:"theme" json:"theme"`
	Layout        string `form:"layout" json:"layout"`
	Ident         string `form:"ident" json:"ident"`
	Wechat_status string `form:"wechat_status" json:"wechat_status" valid:"required"`
}

type AddSystemInput struct {
	Name          string `form:"name" json:"name" valid:"required"`
	Addr          string `form:"addr" json:"addr" valid:"required"`
	Time_out      string `form:"time_out" json:"time_out" valid:"required"`
	Logo          string `form:"logo" json:"logo" valid:"required"`
	Style         string `form:"style" json:"style" valid:"required"`
	Theme         string `form:"theme" json:"theme"`
	Ident         string `form:"ident" json:"ident" vaild:"required"`
	Wechat_status string `form:"wechat_status" json:"wechat_status" valid:"required"`
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
		"name":   name,
		"enable": status,
	})
	fmt.Println("data:", c, q, a)
	if err != nil {
		return nil, 0, fmt.Errorf("获取系统管理列表条数发生错误(err:%v),sql:(%s),输入参数:%v,", err, q, a)
	}
	data, q, a, err = db.Query(sql.QuerySubSystemPageList, map[string]interface{}{
		"name":   name,
		"enable": status,
		"pi":     pi,
		"ps":     ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取系统管理列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	fmt.Println("data:", data, pi, ps, q, a)
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

func (u *DbSystem) Add(input *AddSystemInput) (err error) {
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
		"login_url":     "http://" + strings.Split(strings.Split(input.Addr, "//")[1], "/")[0] + "/member/login",
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

func (u *DbSystem) Edit(input *SystemEditInput) (err error) {
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
	}
	_, q, a, err := db.Execute(sql.UpdateEdit, params)
	if err != nil {
		return fmt.Errorf("更新系统管理数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
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
